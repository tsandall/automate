package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	constants_v2 "github.com/chef/automate/components/authz-service/constants/v2"
	"github.com/chef/automate/components/authz-service/engine"
	storage_errors "github.com/chef/automate/components/authz-service/storage"
	"github.com/chef/automate/components/authz-service/storage/postgres"
	"github.com/chef/automate/components/authz-service/storage/postgres/datamigration"
	"github.com/chef/automate/components/authz-service/storage/postgres/migration"
	v2 "github.com/chef/automate/components/authz-service/storage/v2"
	"github.com/chef/automate/lib/grpc/auth_context"
	"github.com/chef/automate/lib/logger"
	"github.com/chef/automate/lib/projectassignment"
)

// These must match what SQL function query_rule_table_associations returns.
const (
	pgApplied = "applied"
	pgStaged  = "staged"
)

type pg struct {
	db          *sql.DB
	engine      engine.Engine
	logger      logger.Logger
	dataMigConf datamigration.Config
	conninfo    string
}

var singletonInstance *pg
var once sync.Once

// GetInstance returns the signleton instance. Will be nil if not yet initialized.
func GetInstance() *pg {
	return singletonInstance
}

// New instantiates the singleton postgres storage backend.
// Will only initialize once. Will simply return nil if already initialized.
func Initialize(ctx context.Context, e engine.Engine, l logger.Logger, migConf migration.Config,
	dataMigConf datamigration.Config) error {

	var err error
	once.Do(func() {
		l.Infof("applying database migrations from %s", migConf.Path)
		var db *sql.DB
		db, err = postgres.New(ctx, migConf)
		singletonInstance = &pg{db: db, engine: e, logger: l, dataMigConf: dataMigConf, conninfo: migConf.PGURL.String()}
	})
	return err
}

type Querier interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * POLICIES  * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// CreatePolicy stores a new policy and its statements in postgres and returns the final policy.
func (p *pg) CreatePolicy(ctx context.Context, pol *v2.Policy, checkProjects bool) (*v2.Policy, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Note(sr): we're using BeginTx with the context that'll be cancelled in a
	// `defer` when the function ends. This should rollback transactions that
	// haven't been committed -- what would happen when any of the following
	// `err != nil` cases return early.
	// However, I haven't played with this extensively, so there's a bit of a
	// chance that this understanding is just plain wrong.

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	if checkProjects {
		err = p.ensureNoProjectsMissingWithQuerier(ctx, tx, pol.Projects)
		if err != nil {
			return nil, p.processError(err)
		}

		err = projectassignment.EnsureProjectAssignmentAuthorized(ctx,
			p.engine,
			auth_context.FromContext(auth_context.FromIncomingMetadata(ctx)).Subjects,
			pol.Projects)
		if err != nil {
			return nil, p.processError(err)
		}
	}

	err = p.insertPolicyWithQuerier(ctx, pol, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.associatePolicyWithProjects(ctx, pol.ID, pol.Projects, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.insertPolicyStatementsWithQuerier(ctx, pol.ID, pol.Statements, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	// Currently, we don't change anything from what is passed in.
	return pol, nil
}

func (p *pg) PurgeSubjectFromPolicies(ctx context.Context, sub string) ([]string, error) {
	var polIDs []string
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}
	// Note(sr) 2018-11-26: We're keeping the iam_members reference. Should we
	// remove it? "Just" removing the iam_members entry and relying to CASCADE to
	// remove the membership rows from iam_policy_members doesn't do the trick here
	// -- not if we care about the affected policy IDs. (We at the moment don't
	// prescribe this, but it feels like the better choice.)

	row := tx.QueryRowContext(ctx, `
		WITH pol_db_ids AS (
			DELETE FROM iam_policy_members
			WHERE member_id=(SELECT db_id FROM iam_members WHERE name=$1)
			RETURNING policy_id
		)
		SELECT array_agg(id)
		FROM iam_policies
		WHERE db_id IN (SELECT * FROM pol_db_ids)`, sub)
	err = row.Scan(pq.Array(&polIDs))
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, p.processError(err)
	}

	return polIDs, nil
}

func (p *pg) ListPolicies(ctx context.Context) ([]*v2.Policy, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var pols []*v2.Policy
	rows, err := p.db.QueryContext(ctx,
		"SELECT query_policies from query_policies($1)", pq.Array(projectsFilter))
	if err != nil {
		return nil, p.processError(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	for rows.Next() {
		var pol v2.Policy
		if err := rows.Scan(&pol); err != nil {
			return nil, p.processError(err)
		}
		pols = append(pols, &pol)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error retrieving result rows")
	}
	return pols, nil
}

func (p *pg) GetPolicy(ctx context.Context, id string) (*v2.Policy, error) {
	pol, err := p.queryPolicy(ctx, id, p.db, false)
	if err != nil {
		return nil, p.processError(err)
	}

	return pol, nil
}

func (p *pg) DeletePolicy(ctx context.Context, id string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return p.processError(err)
	}

	// Project filtering handled in here
	_, err = p.queryPolicy(ctx, id, tx, false)
	if err != nil {
		return p.processError(err)
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM iam_policies WHERE id=$1;`,
		id,
	)
	if err != nil {
		return p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}

	return nil
}

func (p *pg) UpdatePolicy(ctx context.Context, pol *v2.Policy, checkProjects bool) (*v2.Policy, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	// Project filtering handled in here. We'll return a 404 right away if we can't find
	// the policy via ID as filtered by projects. Also locks relevant rows if in v2.1 mode
	// so we can check project assignment permissions without them being changed under us.
	oldPolicy, err := p.queryPolicy(ctx, pol.ID, tx, checkProjects)
	if err != nil {
		return nil, p.processError(err)
	}

	if checkProjects {
		projectDiff := projectassignment.CalculateProjectDiff(oldPolicy.Projects, pol.Projects)

		if len(projectDiff) != 0 {
			err = p.ensureNoProjectsMissingWithQuerier(ctx, tx, projectDiff)
			if err != nil {
				return nil, p.processError(err)
			}

			err = projectassignment.EnsureProjectAssignmentAuthorized(ctx,
				p.engine,
				auth_context.FromContext(auth_context.FromIncomingMetadata(ctx)).Subjects,
				projectDiff)
			if err != nil {
				return nil, p.processError(err)
			}
		}
	}

	// Since we are forcing users to update the entire policy, we should delete
	// all existing statements for simplicity for now. Let's not delete the actual
	// policy row to preserve that record / id.
	//
	// This will cascade delete all related statements.
	if _, err := tx.ExecContext(ctx,
		"DELETE FROM iam_statements WHERE policy_id=policy_db_id($1)",
		pol.ID,
	); err != nil {
		if err := p.processError(err); err != storage_errors.ErrNotFound {
			return nil, err
		}
	}

	res, err := tx.ExecContext(ctx,
		"UPDATE iam_policies SET (name, type) = ($2, $3) WHERE id = $1 RETURNING id",
		pol.ID, pol.Name, pol.Type.String(),
	)
	if err != nil {
		return nil, p.processError(err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return nil, p.processError(err)
	}
	if affected == 0 {
		return nil, storage_errors.ErrNotFound
	}

	// Update policy's projects
	if err := p.associatePolicyWithProjects(ctx, pol.ID, pol.Projects, tx); err != nil {
		return nil, p.processError(err)
	}

	// Also replace any existing policy members and update with new members.
	if err := p.replacePolicyMembersWithQuerier(ctx, pol.ID, pol.Members, tx); err != nil {
		return nil, p.processError(err)
	}

	if err := p.insertPolicyStatementsWithQuerier(ctx, pol.ID, pol.Statements, tx); err != nil {
		return nil, p.processError(err)
	}

	if err := p.notifyPolicyChange(ctx, tx); err != nil {
		return nil, p.processError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	// Currently, we don't change anything from what is passed in.
	return pol, nil
}

func (p *pg) ApplyV2DataMigrations(_ context.Context) error {
	return p.dataMigConf.Migrate()
}

func (p *pg) GetPolicyChangeID(ctx context.Context) (string, error) {
	var policyChangeID string

	row := p.db.QueryRowContext(ctx, "SELECT policy_change_id FROM policy_change_tracker LIMIT 1;")

	if err := row.Scan(&policyChangeID); err != nil {
		return "", p.processError(err)
	}

	return policyChangeID, nil
}

func (p *pg) GetPolicyChangeNotifier(ctx context.Context) (v2.PolicyChangeNotifier, error) {
	return newPolicyChangeNotifier(ctx, p.conninfo)
}

// insertPolicyWithQuerier inserts a new custom policy. It does not return the
// new policy since there are no DEFAULTS in the iam_policy table.
func (p *pg) insertPolicyWithQuerier(ctx context.Context, inputPol *v2.Policy, q Querier) error {
	_, err := q.ExecContext(ctx,
		`SELECT insert_iam_policy($1, $2, $3);`,
		inputPol.ID, inputPol.Name, inputPol.Type.String(),
	)
	if err != nil {
		return err
	}

	err = p.replacePolicyMembersWithQuerier(ctx, inputPol.ID, inputPol.Members, q)
	return errors.Wrap(err, "replace policy members")
}

// insertPolicyStatements WithQuerier inserts a new statement and associates it with an existing policy.
// Does not return the statements since they will be exactly the same as passed in since
// statements have no defaults in the database.
func (p *pg) insertPolicyStatementsWithQuerier(ctx context.Context,
	policyID string, inputStatements []v2.Statement,
	q Querier) error {
	for _, s := range inputStatements {
		_, err := q.ExecContext(ctx,
			`SELECT insert_iam_statement_into_policy($1, $2, $3, $4, $5, $6);`,
			policyID, s.Effect.String(), pq.Array(s.Actions),
			pq.Array(s.Resources), s.Role, pq.Array(s.Projects),
		)
		if err != nil {
			return p.processError(err)
		}
	}

	return nil
}

// insertPolicyProjectsWithQuerier creates new associations between a policy and its projects.
func (p *pg) associatePolicyWithProjects(ctx context.Context,
	policyID string, inProjects []string,
	q Querier) error {

	// TODO this might be simplified as we modify how projects are assigned
	// Drop any existing associations.
	_, err := q.ExecContext(ctx,
		"DELETE FROM iam_policy_projects WHERE policy_id=policy_db_id($1)", policyID)
	if err != nil {
		return err
	}
	for _, project := range inProjects {
		_, err := q.ExecContext(ctx,
			`INSERT INTO iam_policy_projects (policy_id, project_id) VALUES (policy_db_id($1), project_db_id($2))`,
			&policyID, &project)
		if err != nil {
			return p.processError(err)
		}
	}

	return nil
}

func (p *pg) notifyPolicyChange(ctx context.Context, q Querier) error {
	// We keep track of an id with each change. This lets us be smart about only updating
	// the OPA rules when it might change.
	_, err := q.ExecContext(ctx, "SELECT notify_policy_change()")
	return err
}

// queryPolicy returns a policy based on id or an error. Can optionally lock for updates.
func (p *pg) queryPolicy(ctx context.Context, id string, q Querier, selectForUpdate bool) (*v2.Policy, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var pol v2.Policy
	query := "SELECT query_policy($1, $2)"
	if selectForUpdate {
		query = "SELECT query_policy($1, $2) FOR UPDATE"
	}
	if err := q.QueryRowContext(ctx, query, id, pq.Array(projectsFilter)).
		Scan(&pol); err != nil {
		return nil, err
	}

	return &pol, nil
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * *  MEMBERS  * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

func (p *pg) ListPolicyMembers(ctx context.Context, id string) ([]v2.Member, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	// Project filtering handled in here. We'll return a 404 here right away if
	// we can't find the policy via ID as filtered by projects.
	_, err = p.queryPolicy(ctx, id, tx, false)
	if err != nil {
		return nil, p.processError(err)
	}

	members, err := p.getPolicyMembersWithQuerier(ctx, id, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	return members, nil
}

func (p *pg) AddPolicyMembers(ctx context.Context, id string, members []v2.Member) ([]v2.Member, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	// Project filtering handled in here. We'll return a 404 right away if we can't find
	// the policy via ID as filtered by projects.
	_, err = p.queryPolicy(ctx, id, tx, false)
	if err != nil {
		return nil, p.processError(err)
	}

	for _, member := range members {
		err := p.insertOrReusePolicyMemberWithQuerier(ctx, id, member, tx)
		if err != nil {
			return nil, p.processError(err)
		}
	}

	members, err = p.getPolicyMembersWithQuerier(ctx, id, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}
	return members, nil
}

func (p *pg) ReplacePolicyMembers(ctx context.Context, policyID string, members []v2.Member) ([]v2.Member, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	// Project filtering handled in here. We'll return a 404 right away if we can't find
	// the policy via ID as filtered by projects.
	_, err = p.queryPolicy(ctx, policyID, tx, false)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.replacePolicyMembersWithQuerier(ctx, policyID, members, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	// fetch fresh data so returned data will reflect that any pre-existing members re-use existing IDs
	members, err = p.getPolicyMembersWithQuerier(ctx, policyID, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	return members, err
}

// RemovePolicyMembers takes in a policy ID and a
// list of members to remove and return the list of remaining users.
func (p *pg) RemovePolicyMembers(ctx context.Context,
	policyID string, members []v2.Member) ([]v2.Member, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	// Project filtering handled in here. We'll return a 404 right away if we can't find
	// the policy via ID as filtered by projects.
	_, err = p.queryPolicy(ctx, policyID, tx, false)
	if err != nil {
		return nil, p.processError(err)
	}

	// Note: we're not using member_db_id() here, since we want to gracefully
	// ignore "not found" errors.
	for _, member := range members {
		_, err := tx.ExecContext(ctx,
			`DELETE FROM iam_policy_members WHERE policy_id=policy_db_id($1) AND
				member_id=(SELECT db_id from iam_members WHERE name=$2)`,
			policyID, member.Name)
		if err != nil {
			err = p.processError(err)
			switch err {
			case storage_errors.ErrNotFound: // continue
			default:
				return nil, err
			}
		}
	}

	// fetch fresh data so returned data will reflect that any pre-existing members re-use existing IDs
	members, err = p.getPolicyMembersWithQuerier(ctx, policyID, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	return members, nil
}

func (p *pg) replacePolicyMembersWithQuerier(ctx context.Context, policyID string, members []v2.Member,
	q Querier) error {
	// Cascading drop any existing members.
	_, err := q.ExecContext(ctx,
		`DELETE FROM iam_policy_members WHERE policy_id=policy_db_id($1);`, policyID)
	if err != nil {
		return err
	}

	// Insert new members.
	for _, member := range members {
		err = p.insertOrReusePolicyMemberWithQuerier(ctx, policyID, member, q)
		if err != nil {
			return err
		}
	}
	return nil
}

// insertOrReusePolicyMemberWithQuerier takes in a member (including a new ID) and a policyID.
// If the member already exists in iam_members, it will ignore the new ID and use
// the existing one. Otherwise, it'll just use the existing ID. In either case,
// it inserts the new or existing member into iam_policy_members association table.
func (p *pg) insertOrReusePolicyMemberWithQuerier(ctx context.Context, policyID string, member v2.Member,
	q Querier) error {
	// First, we insert the member but on conflict do nothing. Then, we insert the member
	// into the policy. This is safe to do non-transactionally right now, since we don't support
	// updating either iam_members id or name columns which is the entire table. Also, we are currently
	// not deleting any of the rows, but reusing them per name string.

	_, err := q.ExecContext(ctx,
		"INSERT INTO iam_members (name) VALUES ($1) ON CONFLICT DO NOTHING",
		member.Name)
	if err != nil {
		return errors.Wrapf(err, "failed to upsert member %s", member.Name)
	}

	// For now, let's just ignore conflicts if someone is trying to add a user that is already a member.
	_, err = q.ExecContext(ctx,
		`INSERT INTO iam_policy_members (policy_id, member_id)
			VALUES (policy_db_id($1), member_db_id($2)) ON CONFLICT DO NOTHING`, policyID, member.Name)
	return errors.Wrapf(err, "failed to upsert member link: member=%s, policy_id=%s", member.Name, policyID)
}

func (p *pg) getPolicyMembersWithQuerier(ctx context.Context, id string, q Querier) ([]v2.Member, error) {
	rows, err := q.QueryContext(ctx,
		`SELECT m.name FROM iam_policy_members AS pm
			JOIN iam_members AS m ON pm.member_id=m.db_id
			WHERE pm.policy_id=policy_db_id($1) ORDER BY m.name ASC`, id)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	members := []v2.Member{}
	for rows.Next() {
		var member v2.Member
		if err := rows.Scan(&member.Name); err != nil {
			return nil, p.processError(err)
		}
		members = append(members, member)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error retrieving result rows")
	}
	return members, nil
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * *   ROLES   * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

func (p *pg) CreateRole(ctx context.Context, role *v2.Role, checkProjects bool) (*v2.Role, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	if checkProjects {
		err = p.ensureNoProjectsMissingWithQuerier(ctx, tx, role.Projects)
		if err != nil {
			return nil, p.processError(err)
		}

		err = projectassignment.EnsureProjectAssignmentAuthorized(ctx,
			p.engine,
			auth_context.FromContext(auth_context.FromIncomingMetadata(ctx)).Subjects,
			role.Projects)
		if err != nil {
			return nil, p.processError(err)
		}
	}

	err = p.insertRoleWithQuerier(ctx, role, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	return role, nil
}

func (p *pg) ListRoles(ctx context.Context) ([]*v2.Role, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var roles []*v2.Role
	rows, err := p.db.QueryContext(ctx, "SELECT query_roles($1)", pq.Array(projectsFilter))
	if err != nil {
		return nil, p.processError(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	for rows.Next() {
		var role v2.Role
		if err := rows.Scan(&role); err != nil {
			return nil, p.processError(err)
		}
		roles = append(roles, &role)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error retrieving result rows")
	}
	return roles, nil
}

func (p *pg) GetRole(ctx context.Context, id string) (*v2.Role, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	doesIntersect, err := checkIfRoleIntersectsProjectsFilter(ctx, tx, id, projectsFilter)
	if err != nil {
		return nil, p.processError(err)
	}
	if !doesIntersect {
		return nil, storage_errors.ErrNotFound
	}

	var role v2.Role
	row := tx.QueryRowContext(ctx, `SELECT query_role($1);`, id)
	err = row.Scan(&role)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	return &role, nil
}

func (p *pg) DeleteRole(ctx context.Context, id string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return err
	}

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return p.processError(err)
	}

	doesIntersect, err := checkIfRoleIntersectsProjectsFilter(ctx, tx, id, projectsFilter)
	if err != nil {
		return p.processError(err)
	}
	if !doesIntersect {
		return storage_errors.ErrNotFound
	}

	res, err := tx.ExecContext(ctx, "DELETE FROM iam_roles WHERE id=$1", id)
	if err != nil {
		return p.processError(err)
	}

	err = p.singleRowResultOrNotFoundErr(res)
	if err != nil {
		return err
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}

	return nil
}

func (p *pg) UpdateRole(ctx context.Context, role *v2.Role, checkProjects bool) (*v2.Role, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return nil, p.processError(err)
	}

	doesIntersect, err := checkIfRoleIntersectsProjectsFilter(ctx, tx, role.ID, projectsFilter)
	if err != nil {
		return nil, p.processError(err)
	}
	if !doesIntersect {
		return nil, storage_errors.ErrNotFound
	}

	if checkProjects {
		var oldRole v2.Role
		// get the old role and lock the role for updates (still readable)
		// until the update completes or the transaction fails so that
		// the project diff doesn't change under us while we perform permission checks.
		row := tx.QueryRowContext(ctx, `SELECT query_role($1) FOR UPDATE;`, role.ID)
		err = row.Scan(&oldRole)
		if err != nil {
			return nil, p.processError(err)
		}

		projectDiff := projectassignment.CalculateProjectDiff(oldRole.Projects, role.Projects)

		if len(projectDiff) != 0 {
			err = p.ensureNoProjectsMissingWithQuerier(ctx, tx, projectDiff)
			if err != nil {
				return nil, p.processError(err)
			}

			err = projectassignment.EnsureProjectAssignmentAuthorized(ctx,
				p.engine,
				auth_context.FromContext(auth_context.FromIncomingMetadata(ctx)).Subjects,
				projectDiff)
			if err != nil {
				return nil, p.processError(err)
			}
		}
	}

	row := tx.QueryRowContext(ctx,
		`UPDATE iam_roles SET (name, actions) =
			($2, $3) WHERE id = $1 RETURNING db_id`,
		role.ID, role.Name, pq.Array(role.Actions),
	)
	// TODO: check not found case
	var dbID string
	if err := row.Scan(&dbID); err != nil {
		return nil, p.processError(err)
	}

	_, err = tx.ExecContext(ctx,
		"DELETE FROM iam_role_projects WHERE role_id=$1", dbID)
	if err != nil {
		return nil, p.processError(err)
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO iam_role_projects (role_id, project_id)
		SELECT $1, db_id FROM iam_projects WHERE id=ANY($2)`,
		dbID, pq.Array(role.Projects))
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.notifyPolicyChange(ctx, tx)
	if err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	return role, nil
}

func checkIfRoleIntersectsProjectsFilter(ctx context.Context, q Querier,
	id string, projectsFilter []string) (bool, error) {

	// If no filter was specified, do not filter.
	if len(projectsFilter) == 0 {
		return true, nil
	}

	// Return true or false if there is intersection between iam_role_projects and projectsFilter,
	// assuming '{(unassigned)}' in the case that iam_role_projects is empty. If a role of id
	// doesn't exist, this will return a proper SQL "no rows" error when passed to processError.
	row := q.QueryRowContext(ctx, "SELECT projects_match(role_projects($1), $2)", id, pq.Array(projectsFilter))

	var result bool
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (p *pg) insertRoleWithQuerier(ctx context.Context, role *v2.Role, q Querier) error {
	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return p.processError(err)
	}

	row := q.QueryRowContext(ctx, `INSERT INTO iam_roles (id, name, type, actions) VALUES ($1, $2, $3, $4)
		RETURNING db_id`,
		role.ID, role.Name, role.Type.String(), pq.Array(role.Actions))
	var dbID string
	if err := row.Scan(&dbID); err != nil {
		return p.processError(err)
	}

	_, err = q.ExecContext(ctx,
		`INSERT INTO iam_role_projects (role_id, project_id)
		SELECT $1, project_db_id(p) FROM unnest($2::TEXT[]) as p`,
		dbID, pq.Array(role.Projects))
	if err != nil {
		return p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}

	return nil
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * *    Rules    * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

func (p *pg) CreateRule(ctx context.Context, rule *v2.Rule) (*v2.Rule, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, p.processError(err)
	}

	assocMap, err := p.getMapOfRuleAssociations(ctx, tx, rule.ID, rule.ProjectID)
	if err != nil {
		return nil, p.processError(err)
	}

	// If any associations return, then the rule already exists in current, staged, or both tables
	if len(assocMap) > 0 {
		return nil, storage_errors.ErrConflict
	}

	row := tx.QueryRowContext(ctx,
		`INSERT INTO iam_staged_project_rules (id, project_id, name, type, deleted)
		VALUES ($1, project_db_id($2), $3, $4, false)
		RETURNING db_id`,
		rule.ID, rule.ProjectID, rule.Name, rule.Type.String())
	var ruleDbID string
	if err := row.Scan(&ruleDbID); err != nil {
		return nil, p.processError(err)
	}

	for _, condition := range rule.Conditions {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO iam_staged_rule_conditions (rule_db_id, value, attribute, operator) VALUES ($1, $2, $3, $4);`,
			ruleDbID, pq.Array(condition.Value), condition.Attribute.String(), condition.Operator.String(),
		)
		if err != nil {
			return nil, p.processError(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	rule.Status = pgStaged
	return rule, nil
}

func (p *pg) UpdateRule(ctx context.Context, rule *v2.Rule) (*v2.Rule, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, p.processError(err)
	}

	row := tx.QueryRowContext(ctx,
		`SELECT update_rule($1, $2, $3, $4, $5)`,
		rule.ID, rule.ProjectID, rule.Name, rule.Type.String(), pq.Array(projectsFilter))
	var ruleDbID int
	if err := row.Scan(&ruleDbID); err != nil {
		return nil, p.processError(err)
	}

	// Delete the existing conditions. Don't need to worry about "not found" case since a rule must have conditions.
	_, err = tx.ExecContext(ctx, `DELETE FROM iam_staged_rule_conditions WHERE rule_db_id=$1;`, ruleDbID)
	if err != nil {
		return nil, p.processError(err)
	}

	for _, condition := range rule.Conditions {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO iam_staged_rule_conditions (rule_db_id, value, attribute, operator) VALUES ($1, $2, $3, $4);`,
			ruleDbID, pq.Array(condition.Value), condition.Attribute.String(), condition.Operator.String(),
		)
		if err != nil {
			return nil, p.processError(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	rule.Status = pgStaged
	return rule, nil
}

func (p *pg) DeleteRule(ctx context.Context, projectID string, ruleID string) error {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return p.processError(err)
	}

	assocMap, err := p.getMapOfRuleAssociations(ctx, tx, ruleID, projectID)
	if err != nil {
		return p.processError(err)
	}

	ruleStaged := assocMap[pgStaged]
	ruleApplied := assocMap[pgApplied]

	if !ruleStaged && !ruleApplied {
		return storage_errors.ErrNotFound
	}

	if ruleApplied && ruleStaged {
		res, err := tx.ExecContext(ctx,
			`UPDATE iam_staged_project_rules
				SET deleted=true
				WHERE id=$1 AND project_id=project_db_id($2) AND projects_match_for_rule($2, $3)`,
			ruleID, projectID, pq.Array(projectsFilter),
		)
		if err != nil {
			return p.processError(err)
		}
		err = p.singleRowResultOrNotFoundErr(res)
		if err != nil {
			return err
		}
	} else if ruleApplied {
		res, err := tx.ExecContext(ctx,
			`SELECT db_id FROM iam_project_rules
				WHERE id=$1 AND project_id=project_db_id($2) AND projects_match_for_rule($2, $3)`,
			ruleID, projectID, pq.Array(projectsFilter),
		)
		if err != nil {
			return p.processError(err)
		}
		err = p.singleRowResultOrNotFoundErr(res)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx,
			`INSERT INTO iam_staged_project_rules (id, project_id, name, type, deleted)
				SELECT a.id, a.project_id, a.name, a.type, 'true'
				FROM iam_project_rules AS a
				WHERE a.id=$1 AND projects_match_for_rule(a.project_id, $2)`,
			ruleID, pq.Array(projectsFilter),
		)
		if err != nil {
			return p.processError(err)
		}
		// Code is built around expectation that rules always have at least one condition,
		// that means even in the case of impending deletion!
		// Value will never be seen, so a dummy value is OK here.
		_, err = tx.ExecContext(ctx,
			`INSERT INTO iam_staged_rule_conditions (rule_db_id, value, attribute, operator)
			 (SELECT db_id, '{dummy}', 'chef-server', 'equals'  FROM iam_staged_project_rules WHERE id=$1)`,
			ruleID,
		)
		if err != nil {
			return p.processError(err)
		}
	} else if ruleStaged {
		res, err := tx.ExecContext(ctx,
			`DELETE FROM iam_staged_project_rules
				WHERE id=$1 AND project_id=project_db_id($2) AND projects_match_for_rule($2, $3)`,
			ruleID, projectID, pq.Array(projectsFilter),
		)
		if err != nil {
			return p.processError(err)
		}
		err = p.singleRowResultOrNotFoundErr(res)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}

	return nil
}

func (p *pg) GetStagedOrAppliedRule(ctx context.Context, projectID string, ruleID string) (*v2.Rule, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var rule v2.Rule
	row := p.db.QueryRowContext(ctx, "SELECT query_staged_or_applied_rule($1, $2, $3)",
		ruleID, projectID, pq.Array(projectsFilter),
	)
	err = row.Scan(&rule)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, storage_errors.ErrNotFound
		}
		return nil, p.processError(err)
	}

	return &rule, nil
}

func (p *pg) FetchAppliedRulesByProjectIDs(ctx context.Context) (map[string][]*v2.Rule, error) {
	rules, err := p.listRulesUsingFunction(ctx, "SELECT query_rules($1)", false)
	if err != nil {
		return nil, err
	}

	projectRules := make(map[string][]*v2.Rule, len(rules))
	for _, rule := range rules {
		projectRules[rule.ProjectID] = append(projectRules[rule.ProjectID], rule)
	}

	return projectRules, nil
}

func (p *pg) ListRules(ctx context.Context) ([]*v2.Rule, error) {
	return p.listRulesUsingFunction(ctx, "SELECT query_rules($1)", true)
}

func (p *pg) ListStagedAndAppliedRules(ctx context.Context) ([]*v2.Rule, error) {
	return p.listRulesUsingFunction(ctx, "SELECT query_staged_and_applied_rules($1)", true)
}

func (p *pg) listRulesUsingFunction(ctx context.Context, query string, filterProjects bool) ([]*v2.Rule, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	projectsFilter := []string{}
	if filterProjects {
		var err error
		projectsFilter, err = projectsListFromContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	var rules []*v2.Rule
	rows, err := p.db.QueryContext(ctx, query, pq.Array(projectsFilter))
	if err != nil {
		return nil, p.processError(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	for rows.Next() {
		var rule v2.Rule
		if err := rows.Scan(&rule); err != nil {
			return nil, p.processError(err)
		}
		rules = append(rules, &rule)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error retrieving result rows")
	}
	return rules, nil
}

func (p *pg) ListRulesForProject(ctx context.Context, projectID string) ([]*v2.Rule, v2.ProjectRulesStatus, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, v2.RulesStatusError, err
	}

	// in our other APIs we use a a postgres query to do filtering
	// however in this case, we can't automatically assume NoRows means NotFound
	// because we want to differentiate between a project that is not in the project filter
	// and a project that has no rules
	if len(projectsFilter) > 0 {
		var projectInFilter bool
		for _, id := range projectsFilter {
			if id == projectID {
				projectInFilter = true
				break
			}
		}
		if !projectInFilter {
			return nil, v2.RulesStatusError, storage_errors.ErrNotFound
		}
	}

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, v2.RulesStatusError, p.processError(err)
	}

	var rules []*v2.Rule
	rows, err := tx.QueryContext(ctx, "SELECT query_rules_for_project($1, $2)",
		projectID, pq.Array(projectsFilter))
	if err != nil {
		return nil, v2.RulesStatusError, p.processError(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	anyStagedRules := false
	for rows.Next() {
		var rule v2.Rule
		if err := rows.Scan(&rule); err != nil {
			return nil, v2.RulesStatusError, p.processError(err)
		}
		if rule.Status == pgStaged {
			anyStagedRules = true
		}
		rules = append(rules, &rule)
	}
	if err := rows.Err(); err != nil {
		return nil, v2.RulesStatusError, errors.Wrap(err, "error retrieving result rows")
	}

	err = tx.Commit()
	if err != nil {
		return nil, v2.RulesStatusError, storage_errors.NewTxCommitError(err)
	}

	rulesStatus := v2.Applied
	if len(rules) == 0 {
		rulesStatus = v2.NoRules
	}
	if anyStagedRules {
		rulesStatus = v2.EditsPending
	}

	return rules, rulesStatus, nil
}

// ApplyStagedRules begins a db transaction, locks the rule tables, moves all staged rule updates
// and deletes into the applied rule table, and returns the database transaction. The transaction is returned
// so that other non-database concerns can be completed before freeing the lock to avoid race conditions.
func (p *pg) ApplyStagedRules(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return p.processError(err)
	}

	_, err = tx.ExecContext(ctx,
		`LOCK TABLE iam_project_rules;
			LOCK TABLE iam_rule_conditions;
			LOCK TABLE iam_staged_project_rules;
			LOCK TABLE iam_staged_rule_conditions; `,
	)
	if err != nil {
		return p.processError(err)
	}

	// Upsert all staged rules into applied rules marked for update, returning the id and db_id
	// of all rules affected so we can update their conditions below.
	rows, err := tx.QueryContext(ctx,
		`INSERT INTO iam_project_rules (id, project_id, name, type)
				SELECT s.id, s.project_id, s.name, s.type
					FROM iam_staged_project_rules AS s
					WHERE deleted=false
				ON CONFLICT (id) DO UPDATE
				SET name=excluded.name, type=excluded.type
				RETURNING id, db_id;`)
	if err != nil {
		return p.processError(err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	// For every staged rule updated, we need to update conditions.
	ids := make(map[string]string)
	for rows.Next() {
		var id string
		var dbID string
		err = rows.Scan(&id, &dbID)
		if err != nil {
			return p.processError(err)
		}
		ids[id] = dbID
	}
	if err := rows.Err(); err != nil {
		return errors.Wrap(err, "error retrieving result rows")
	}

	for id, dbID := range ids {
		_, err = tx.ExecContext(ctx,
			`DELETE FROM iam_rule_conditions WHERE rule_db_id=$1;`, dbID)
		if err != nil {
			return p.processError(err)
		}

		_, err = tx.ExecContext(ctx,
			`INSERT INTO iam_rule_conditions (rule_db_id, value, attribute, operator)
					SELECT $2, cond.value, cond.attribute, cond.operator
						FROM iam_staged_project_rules AS r
					LEFT OUTER JOIN iam_staged_rule_conditions AS cond
						ON rule_db_id=r.db_id
						WHERE r.id=$1;`,
			id, dbID,
		)
		if err != nil {
			return p.processError(err)
		}
	}

	_, err = tx.ExecContext(ctx,
		`DELETE FROM iam_project_rules
				WHERE id IN (SELECT id FROM iam_staged_project_rules WHERE deleted)`)
	if err != nil {
		return p.processError(err)
	}

	_, err = tx.ExecContext(ctx, `DELETE FROM iam_staged_project_rules;`)
	if err != nil {
		return p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return storage_errors.NewTxCommitError(err)
	}

	return nil
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * *   PROJECTS  * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

func (p *pg) CreateProject(ctx context.Context, project *v2.Project) (*v2.Project, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, p.processError(err)
	}

	if project.Type == v2.Custom {
		row := tx.QueryRowContext(ctx, "SELECT count(*) FROM iam_projects WHERE type='custom'")
		var numProjects int64
		if err := row.Scan(&numProjects); err != nil {
			return nil, p.processError(err)
		}

		if numProjects >= constants_v2.MaxProjects {
			return nil, storage_errors.ErrMaxProjectsExceeded
		}
	}

	if err := p.insertProjectWithQuerier(ctx, project, tx); err != nil {
		return nil, p.processError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, storage_errors.NewTxCommitError(err)
	}

	// Currently, we don't change anything from what is passed in.
	return project, nil
}

func (p *pg) UpdateProject(ctx context.Context, project *v2.Project) (*v2.Project, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// Update project if ID found AND intersection between projects and projectsFilter,
	// unless the projectsFilter is empty (v2.0 case).
	res, err := p.db.ExecContext(ctx,
		`UPDATE iam_projects SET name=$2
		WHERE id=$1 AND (array_length($3::TEXT[], 1) IS NULL OR id=ANY($3));`,
		project.ID, project.Name, pq.Array(projectsFilter))
	if err != nil {
		return nil, p.processError(err)
	}

	err = p.singleRowResultOrNotFoundErr(res)
	if err != nil {
		return nil, err
	}

	// Currently, we don't change anything from what is passed in.
	return project, nil
}

func (p *pg) GetProject(ctx context.Context, id string) (*v2.Project, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var project v2.Project
	// Retrieve project if ID found AND intersection between projects and projectsFilter,
	// unless the projectsFilter is empty (v2.0 case).
	row := p.db.QueryRowContext(ctx, `SELECT query_project($1, $2)`, id, pq.Array(projectsFilter))
	if err := row.Scan(&project); err != nil {
		return nil, p.processError(err)
	}
	return &project, nil
}

func (p *pg) DeleteProject(ctx context.Context, id string) error {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return err
	}

	// Delete project if ID found AND intersection between projects and projectsFilter,
	// unless the projectsFilter is empty (v2.0 case).
	res, err := p.db.ExecContext(ctx,
		`DELETE FROM iam_projects WHERE id=$1 AND (array_length($2::TEXT[], 1) IS NULL OR id=ANY($2));`,
		id, pq.Array(projectsFilter),
	)
	if err != nil {
		return p.processError(err)
	}

	err = p.singleRowResultOrNotFoundErr(res)
	if err != nil {
		return err
	}

	return nil
}

// EnsureNoProjectsMissing returns projectassignment.ProjectsMissingError if projects are missing,
// otherwise it returns nil.
func (p *pg) EnsureNoProjectsMissing(ctx context.Context, projectIDs []string) error {
	return p.ensureNoProjectsMissingWithQuerier(ctx, p.db, projectIDs)
}

func (p *pg) ensureNoProjectsMissingWithQuerier(ctx context.Context, q Querier, projectIDs []string) error {
	// Return any input ID that does not exist in the projects table.
	rows, err := p.db.QueryContext(ctx,
		`SELECT id FROM unnest($1::text[]) AS input(id)
			WHERE NOT EXISTS (SELECT * FROM iam_projects p WHERE input.id = p.id);`, pq.Array(projectIDs))
	if err != nil {
		return p.processError(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	projectsNotFound := make([]string, 0)
	for rows.Next() {
		var projectIDNotFound string
		if err := rows.Scan(&projectIDNotFound); err != nil {
			return p.processError(err)
		}
		projectsNotFound = append(projectsNotFound, projectIDNotFound)
	}

	if len(projectsNotFound) != 0 {
		return projectassignment.NewProjectsMissingErroror(projectsNotFound)
	}

	return nil
}

func (p *pg) ListProjects(ctx context.Context) ([]*v2.Project, error) {
	projectsFilter, err := projectsListFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// List all projects that have intersection between projects and projectsFilter,
	// unless the projectsFilter is empty (v2.0 case).
	rows, err := p.db.QueryContext(ctx, "SELECT query_projects($1)", pq.Array(projectsFilter))
	if err != nil {
		return nil, p.processError(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			p.logger.Warnf("failed to close db rows: %s", err.Error())
		}
	}()

	var projects []*v2.Project
	for rows.Next() {
		var project v2.Project
		if err := rows.Scan(&project); err != nil {
			return nil, p.processError(err)
		}
		projects = append(projects, &project)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error retrieving result rows")
	}
	return projects, nil
}

func (p *pg) insertProjectWithQuerier(ctx context.Context, project *v2.Project, q Querier) error {
	_, err := q.ExecContext(ctx, `INSERT INTO iam_projects (id, name, type) VALUES ($1, $2, $3);`,
		project.ID, project.Name, project.Type.String())
	return err
}

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * SUPPORT * * * * * * * * * * * * * * * * * * * */
/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

func (p *pg) Reset(ctx context.Context) error {
	if _, err := p.db.ExecContext(ctx,
		`TRUNCATE TABLE iam_policies, iam_members, iam_roles, iam_projects CASCADE;`); err != nil {
		return errors.Wrap(err, "truncate database")
	}

	if err := p.dataMigConf.Reset(); err != nil {
		return errors.Wrap(err, "reset v2 data migrations")
	}

	return nil
}

func (p *pg) Close() error {
	err := errors.Wrap(p.db.Close(), "close database connection")
	// reset the singleton
	once = *new(sync.Once)
	singletonInstance = nil
	return err
}

func (p *pg) Pristine(ctx context.Context) error {
	return p.recordMigrationStatus(ctx, enumPristine)
}

func (p *pg) singleRowResultOrNotFoundErr(result sql.Result) error {
	count, err := result.RowsAffected()
	if err != nil {
		return p.processError(err)
	}
	if count == 0 {
		return storage_errors.ErrNotFound
	}
	if count > 1 {
		return storage_errors.ErrDatabase
	}
	return nil
}

func (p *pg) getMapOfRuleAssociations(ctx context.Context, q Querier, id string, projectID string) (map[string]bool, error) {
	assocRow := q.QueryRowContext(ctx, "SELECT query_rule_table_associations($1, $2)", id, projectID)
	var associations []string
	if err := assocRow.Scan(pq.Array(&associations)); err != nil {
		return nil, err
	}

	set := make(map[string]bool, len(associations))
	for _, s := range associations {
		set[s] = true
	}
	return set, nil
}

func (p *pg) recordMigrationStatusAndNotifyPG(ctx context.Context, ms string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil /* use driver default */)
	if err != nil {
		return p.processError(err)
	}
	if err := p.recordMigrationStatusWithQuerier(ctx, ms, tx); err != nil {
		return p.processError(err)
	}
	if err := p.notifyPolicyChange(ctx, tx); err != nil {
		return p.processError(err)
	}
	if err := tx.Commit(); err != nil {
		return storage_errors.NewTxCommitError(err)
	}
	return nil
}

func (p *pg) Success(ctx context.Context) error {
	return p.recordMigrationStatusAndNotifyPG(ctx, enumSuccessful)
}

func (p *pg) SuccessBeta1(ctx context.Context) error {
	return p.recordMigrationStatusAndNotifyPG(ctx, enumSuccessfulBeta1)
}

func (p *pg) InProgress(ctx context.Context) error {
	return p.recordMigrationStatus(ctx, enumInProgress)
}

func (p *pg) Failure(ctx context.Context) error {
	return p.recordMigrationStatus(ctx, enumFailed)
}

func (p *pg) MigrationStatus(ctx context.Context) (v2.MigrationStatus, error) {
	var status string
	row := p.db.QueryRowContext(ctx, `SELECT state FROM migration_status`)
	err := row.Scan(&status)
	if err != nil {
		return 0, err // shouldn't happen, migration initializes state
	}
	switch status {
	case enumPristine:
		return v2.Pristine, nil
	case enumSuccessful:
		return v2.Successful, nil
	case enumSuccessfulBeta1:
		return v2.SuccessfulBeta1, nil
	case enumInProgress:
		return v2.InProgress, nil
	case enumFailed:
		return v2.Failed, nil
	}
	return 0, fmt.Errorf("unexpected migration status: %q", status)
}

const (
	enumPristine        = "init"
	enumInProgress      = "in-progress"
	enumSuccessful      = "successful"
	enumSuccessfulBeta1 = "successful-beta1"
	enumFailed          = "failed"
)

func (p *pg) recordMigrationStatus(ctx context.Context, ms string) error {
	return p.recordMigrationStatusWithQuerier(ctx, ms, p.db)
}

func (p *pg) recordMigrationStatusWithQuerier(ctx context.Context, ms string, q Querier) error {
	_, err := q.ExecContext(ctx, `UPDATE migration_status SET state=$1`, ms)
	return err
}

func (p *pg) processError(err error) error {
	p.logger.Debugf("err: %v", err)
	err = postgres.ProcessError(err)
	if err == storage_errors.ErrDatabase {
		p.logger.Warnf("unknown error type from database: %v", err)
	}
	return err
}

// projectsListFromContext returns the project list from the context.
// In the case that the project list was ["*"], we return an empty list,
// since we do not wish to filter on projects.
func projectsListFromContext(ctx context.Context) ([]string, error) {
	projectsFilter, err := auth_context.ProjectsFromIncomingContext(ctx)
	if err != nil {
		return nil, err
	}
	if auth_context.AllProjectsRequested(projectsFilter) {
		projectsFilter = []string{}
	}
	return projectsFilter, nil
}
