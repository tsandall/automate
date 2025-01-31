#!/bin/bash

#shellcheck disable=SC2034
test_name="iam v2.1 only"
# Note: the inspec profile takes care of deleting any migrated v1 legacy
# policies
test_deploy_inspec_profiles=(a2-iam-v2p1-only-integration)
# Note: we can't run diagnostics AND inspec, so skip diagnostics
test_skip_diagnostics=true

do_setup() {
    do_setup_default

    hab_pkg_install core/curl
    hab_pkg_install -b core/jo
}

hab_curl() {
    hab pkg exec core/curl curl "$@"
}

# Note: the inspec tests assert that a viewer or operator user is able
# to retrieve the license status -- so we just add one.
do_deploy() {
    do_deploy_default
    log_info "applying dev license"
    chef-automate license apply "$A2_LICENSE"
}

do_test_deploy() {
    log_info "run chef-automate iam upgrade-to-v2 --beta2.1 --skip-policy-migration"
    chef-automate iam upgrade-to-v2 --beta2.1 --skip-policy-migration || return 1

    # ensure service startup works with IAM v2.1:
    # - kill authz-service to force startup,
    # - wait for service status to be healthy again
    log_info "restarting authz-service, waiting 5s for it to come up again"
    pkill -f authz-service && sleep 5 && chef-automate status -w || return 1

    verify_legacy_policies_not_migrated

    do_test_deploy_default
}

verify_legacy_policies_not_migrated() {
    local token legacy_policies code
    token=$(chef-automate iam token create ADMIN_TEST --admin)
    legacy_policies=(secrets-access-legacy
            events-access-legacy
            infrastructure-automation-access-legacy
            compliance-access-legacy
            nodes-access-legacy
            compliance-profile-access-legacy
            ingest-access-legacy
            node-managers-access-legacy
            telemetry-access-legacy)

    for id in "${legacy_policies[@]}"
    do
        echo "checking legacy policy $id..."
	# only capture the response code
        code="$(hab_curl -s -o /dev/null -k -w '%{http_code}' -H "api-token: $token" \
          "https://localhost/apis/iam/v2beta/policies/$id")"
	[[ $code -eq 404 ]] || return 1
    done
}
