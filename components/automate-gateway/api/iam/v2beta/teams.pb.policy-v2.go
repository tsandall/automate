// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/teams.proto

package v2beta

import (
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"
)

func init() {
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/ListTeams", "iam:teams", "iam:teams:list", "GET", "/iam/v2beta/teams", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/GetTeam", "iam:teams:{id}", "iam:teams:get", "GET", "/iam/v2beta/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/CreateTeam", "iam:teams", "iam:teams:create", "POST", "/iam/v2beta/teams", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/UpdateTeam", "iam:teams:{id}", "iam:teams:update", "PUT", "/iam/v2beta/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/DeleteTeam", "iam:teams:{id}", "iam:teams:delete", "DELETE", "/iam/v2beta/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteTeamReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/GetTeamMembership", "iam:teams:{id}:users", "iam:teamUsers:list", "GET", "/iam/v2beta/teams/{id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamMembershipReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/AddTeamMembers", "iam:teams:{id}", "iam:teamUsers:create", "POST", "/iam/v2beta/teams/{id}/users:add", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AddTeamMembersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/RemoveTeamMembers", "iam:teams:{id}", "iam:teamUsers:delete", "POST", "/iam/v2beta/teams/{id}/users:remove", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RemoveTeamMembersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/GetTeamsForMember", "iam:users:{id}:teams", "iam:userTeams:get", "GET", "/iam/v2beta/users/{id}/teams", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamsForMemberReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2beta.Teams/ApplyV2DataMigrations", "iam:teams", "iam:teams:update", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
