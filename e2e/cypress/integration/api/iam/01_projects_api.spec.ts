import { describeIfIAMV2p1, adminApiToken } from '../../constants';

describeIfIAMV2p1('projects API: applying project', () => {
  const cypressPrefix = 'test-projects-api';
  const avengersProject = {
    id: `${cypressPrefix}-project1-${Cypress.moment().format('MMDDYYhhmm')}`,
    name: 'Test Avengers Project'
  };

  interface Project {
    id: string;
    name: string;
    status: string;
  }

  const xmenProject = {
    id: `${cypressPrefix}-project2-${Cypress.moment().format('MMDDYYhhmm')}`,
    name: 'Test X-men Project'
  };

  const avengersRule = {
    id: 'avengers-rule-1',
    name: 'first rule of avengers project',
    type: 'NODE',
    project_id: avengersProject.id,
    conditions: [
      {
        attribute: 'CHEF_ORGANIZATION',
        operator: 'EQUALS',
        values: ['avengers']
      }
    ]
  };

  const xmenRule = {
    id: 'xmen-rule-1',
    name: 'first rule of xmen project',
    type: 'NODE',
    project_id: xmenProject.id,
    conditions: [
      {
        attribute: 'CHEF_ORGANIZATION',
        operator: 'EQUALS',
        values: ['xmen']
      }
    ]
  };

  // testing values
  const noRulesStr = 'NO_RULES';
  const editsPendingStr = 'EDITS_PENDING';
  const rulesAppliedStr = 'RULES_APPLIED';

  before(() => {
    // Cypress recommends state cleanup in the before block to ensure
    // it gets run every time:
    // tslint:disable-next-line:max-line-length
    // https://docs.cypress.io/guides/references/best-practices.html#Using-after-or-afterEach-hooks
    cy.cleanupV2IAMObjectsByIDPrefixes(cypressPrefix, ['projects']);

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'POST',
      url: 'api/v0/ingest/events/chef/node-multiple-deletes',
      body: {
        node_ids: [
          'f6a5c33f-bef5-433b-815e-a8f6e69e6b1b',
          '82760210-4686-497e-b039-efca78dee64b',
          '9c139ad0-89a5-44bc-942c-d7f248b155ba',
          '6453a764-2415-4934-8cee-2a008834a74a'
        ]
      },
      failOnStatusCode: false
    });

    for (const project of [avengersProject, xmenProject]) {
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'POST',
        url: '/apis/iam/v2beta/projects',
        body: project
      });
    }

    let totalNodes = 0;
    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: '(unassigned)'
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      totalNodes = response.body.length;
    });

    cy.fixture('converge/avengers1.json').then(node1 => {
      cy.fixture('converge/avengers2.json').then(node2 => {
        cy.fixture('converge/xmen1.json').then(node3 => {
          cy.fixture('converge/xmen2.json').then(node4 => {
            for (const node of [node1, node2, node3, node4]) {
              cy.request({
                headers: { 'api-token': adminApiToken },
                method: 'POST',
                url: '/data-collector/v0',
                body: node
              });
            }
          });
        });
      });
    });
    // There's no waiting involved, it's sending request after request.
    const maxRetries = 200;
    waitForNodes(totalNodes, maxRetries);

    // confirm nodes are unassigned
    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: '(unassigned)'
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(totalNodes + 4);
    });

    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: avengersProject.id
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(0);
    });

    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: xmenProject.id
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(0);
    });
  });

  after(() => {
    cy.cleanupV2IAMObjectsByIDPrefixes(cypressPrefix, ['projects']);
  });

  it('new rules get applied to nodes', () => {
    // initially no rules
    for (const project of [avengersProject, xmenProject]) {
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'GET',
        url: `/apis/iam/v2beta/projects/${project.id}/rules`
      }).then((response) => {
        expect(response.body.rules).to.have.length(0);
      });

      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'GET',
        url: `/apis/iam/v2beta/projects/${project.id}`
      }).then((response) => {
        expect(response.body.project.status).to.equal(noRulesStr);
      });
    }

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id || id === xmenProject.id)
        .forEach(({ status }) => expect(status).to.equal(noRulesStr));
    });

    for (const rule of [avengersRule, xmenRule]) {
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'POST',
        url: `/apis/iam/v2beta/projects/${rule.project_id}/rules`,
        body: rule
      });
    }

    // confirm rules are staged
    for (const project of [avengersProject, xmenProject]) {
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'GET',
        url: `/apis/iam/v2beta/projects/${project.id}/rules`
      }).then((response) => {
        expect(response.body.rules).to.have.length(1);
        for (const rule of response.body.rules) {
          expect(rule).to.have.property('status', 'STAGED');
        }
      });

      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'GET',
        url: `/apis/iam/v2beta/projects/${project.id}`
      }).then((response) => {
        expect(response.body.project.status).to.equal(editsPendingStr);
      });
    }

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id || id === xmenProject.id)
        .forEach(({ status }) => expect(status).to.equal(editsPendingStr));
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'POST',
      url: '/apis/iam/v2beta/apply-rules'
    });
    waitUntilApplyRulesNotRunning(100);

    // confirm rules are applied
    for (const project of [avengersProject, xmenProject]) {
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'GET',
        url: `/apis/iam/v2beta/projects/${project.id}/rules`
      }).then((response) => {
        expect(response.body.rules).to.have.length(1);
        for (const rule of response.body.rules) {
          expect(rule).to.have.property('status', 'APPLIED');
        }
      });

      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'GET',
        url: `/apis/iam/v2beta/projects/${project.id}`
      }).then((response) => {
        expect(response.body.project.status).to.equal(rulesAppliedStr);
      });
    }

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id || id === xmenProject.id)
        .forEach(({ status }) => expect(status).to.equal(rulesAppliedStr));
    });

    // confirm nodes are assigned to projects correctly
    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: avengersProject.id
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(2);
    });

    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: xmenProject.id
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(2);
    });
  });

  it('rules with updated conditions get applied to nodes', () => {

    // change avengers rule to include both organizations
    const updatedAvengersRule = avengersRule;

    updatedAvengersRule.conditions = [
      {
        attribute: 'CHEF_ORGANIZATION',
        operator: 'MEMBER_OF',
        values: ['avengers', 'xmen']
      }
    ];

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: `/apis/iam/v2beta/projects/${avengersProject.id}`
    }).then((response) => {
      expect(response.body.project.status).to.equal(rulesAppliedStr);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id)
        .forEach(({ status }) => expect(status).to.equal(rulesAppliedStr));
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'PUT',
      url: `/apis/iam/v2beta/projects/${avengersRule.project_id}/rules/${avengersRule.id}`,
      body: updatedAvengersRule
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: `/apis/iam/v2beta/projects/${avengersProject.id}`
    }).then((response) => {
      expect(response.body.project.status).to.equal(editsPendingStr);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id)
        .forEach(({ status }) => expect(status).to.equal(editsPendingStr));
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'POST',
      url: '/apis/iam/v2beta/apply-rules'
    });
    waitUntilApplyRulesNotRunning(100);

    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: avengersProject.id
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(4);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: `/apis/iam/v2beta/projects/${avengersProject.id}`
    }).then((response) => {
      expect(response.body.project.status).to.equal(rulesAppliedStr);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id)
        .forEach(({ status }) => expect(status).to.equal(rulesAppliedStr));
    });
  });

  it('deleted rules get applied to nodes', () => {

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: `/apis/iam/v2beta/projects/${avengersProject.id}`
    }).then((response) => {
      expect(response.body.project.status).to.equal(rulesAppliedStr);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id)
        .forEach(({ status }) => expect(status).to.equal(rulesAppliedStr));
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'DELETE',
      url: `/apis/iam/v2beta/projects/${avengersRule.project_id}/rules/${avengersRule.id}`
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: `/apis/iam/v2beta/projects/${avengersProject.id}`
    }).then((response) => {
      expect(response.body.project.status).to.equal(editsPendingStr);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id)
        .forEach(({ status }) => expect(status).to.equal(editsPendingStr));
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'POST',
      url: '/apis/iam/v2beta/apply-rules'
    });
    waitUntilApplyRulesNotRunning(100);

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: `/apis/iam/v2beta/projects/${avengersProject.id}`
    }).then((response) => {
      expect(response.body.project.status).to.equal(noRulesStr);
    });

    cy.request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/apis/iam/v2beta/projects'
    }).then((response) => {
      const projects: Project[] = response.body.projects;
      projects.filter(({ id, status }) => id === avengersProject.id)
        .forEach(({ status }) => expect(status).to.equal(noRulesStr));
    });

    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: avengersProject.id
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(0);
    });

    cy.request({
      headers: {
        'api-token': adminApiToken,
        projects: '(unassigned)'
      },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    }).then((response) => {
      expect(response.body).to.have.length(2);
    });
  });

  // must correspond to enum type in automate-gateway/.../common/rules.proto
  const attributeValuesMap = new Map([
    ['CHEF_SERVER', 'chef-server-dev.test'],
    ['ENVIRONMENT', '_default'],
    ['CHEF_ROLE', 'web'],
    ['CHEF_TAG', 'test_tag'],
    ['CHEF_POLICY_GROUP', 'test_group'],
    ['CHEF_POLICY_NAME', 'test_name']
    // leave out CHEF_ORGANIZATION since it's tested above
  ]);

  for (const [attribute, value] of attributeValuesMap) {
    const targetNodeId = 'f6a5c33f-bef5-433b-815e-a8f6e69e6b1b';

    it(`ingest rule with attribute ${attribute} gets applied to node`, () => {
      const rule = {
        id: 'rule-2',
        name: 'rule 2',
        type: 'NODE',
        project_id: avengersProject.id,
        conditions: [
          {
            attribute: attribute,
            operator: 'EQUALS',
            values: [value]
          }
        ]
      };
      // create rule with attribute value
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'POST',
        url: `/apis/iam/v2beta/projects/${rule.project_id}/rules`,
        body: rule
      });

      // apply rules
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'POST',
        url: '/apis/iam/v2beta/apply-rules'
      });
      waitUntilApplyRulesNotRunning(100);

      // see filter works
      cy.request({
        headers: {
          'api-token': adminApiToken,
          projects: avengersProject.id
        },
        method: 'GET',
        url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
      }).then((response) => {
        expect(response.body).to.have.length(1);
        expect(response.body[0].id).to.equal(targetNodeId);
      });

      // delete rule
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'DELETE',
        url: `/apis/iam/v2beta/projects/${avengersRule.project_id}/rules/${rule.id}`
      }).then((response) => {
        expect(response.status).to.equal(200);
      });

      // apply rules to reset node to unassigned
      cy.request({
        headers: { 'api-token': adminApiToken },
        method: 'POST',
        url: '/apis/iam/v2beta/apply-rules'
      }).then((response) => {
        expect(response.status).to.equal(200);
      });
      waitUntilApplyRulesNotRunning(100);
    });
  }
});

function waitForNodes(totalNodes: number, maxRetries: number) {
  cy
    .request({
      headers: { 'api-token': adminApiToken },
      method: 'GET',
      url: '/api/v0/cfgmgmt/nodes?pagination.size=10'
    })
    .then((resp: Cypress.ObjectLike) => {
      // to avoid getting stuck in an infinite loop
      if (maxRetries === 0) {
        return;
      }
      if (resp.body.length === totalNodes + 4) {
        return;
      }

      waitForNodes(totalNodes, maxRetries - 1);
    });
}

function waitUntilApplyRulesNotRunning(attempts: number): void {
  if (attempts === -1) {
    throw new Error('apply-rules never finished');
  }
  cy.request({
    headers: { 'api-token': adminApiToken },
    url: '/apis/iam/v2beta/apply-rules'
  }).then((response) => {
    if (response.body.state === 'not_running') {
      return;
    } else {
      cy.log(`${attempts} attempts remaining: waiting for apply-rules to be not_running`);
      cy.wait(1000);
      waitUntilApplyRulesNotRunning(--attempts);
    }
  });
}
