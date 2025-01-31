package api

func init() {
	Swagger.Add("iam_v2beta_teams", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/iam/v2beta/teams.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/iam/v2beta/teams": {
      "get": {
        "operationId": "ListTeams",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaListTeamsResp"
            }
          }
        },
        "tags": [
          "Teams"
        ]
      },
      "post": {
        "operationId": "CreateTeam",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaCreateTeamResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v2betaCreateTeamReq"
            }
          }
        ],
        "tags": [
          "Teams"
        ]
      }
    },
    "/iam/v2beta/teams/{id}": {
      "get": {
        "operationId": "GetTeam",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaGetTeamResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Teams"
        ]
      },
      "delete": {
        "operationId": "DeleteTeam",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaDeleteTeamResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Teams"
        ]
      },
      "put": {
        "operationId": "UpdateTeam",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaUpdateTeamResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v2betaUpdateTeamReq"
            }
          }
        ],
        "tags": [
          "Teams"
        ]
      }
    },
    "/iam/v2beta/teams/{id}/users": {
      "get": {
        "operationId": "GetTeamMembership",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaGetTeamMembershipResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Teams"
        ]
      }
    },
    "/iam/v2beta/teams/{id}/users:add": {
      "post": {
        "operationId": "AddTeamMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaAddTeamMembersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v2betaAddTeamMembersReq"
            }
          }
        ],
        "tags": [
          "Teams"
        ]
      }
    },
    "/iam/v2beta/teams/{id}/users:remove": {
      "post": {
        "operationId": "RemoveTeamMembers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaRemoveTeamMembersResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v2betaRemoveTeamMembersReq"
            }
          }
        ],
        "tags": [
          "Teams"
        ]
      }
    },
    "/iam/v2beta/users/{id}/teams": {
      "get": {
        "operationId": "GetTeamsForMember",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaGetTeamsForMemberResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Teams"
        ]
      }
    }
  },
  "definitions": {
    "v2betaAddTeamMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "user_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaAddTeamMembersResp": {
      "type": "object",
      "properties": {
        "user_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaApplyV2DataMigrationsResp": {
      "type": "object"
    },
    "v2betaCreateTeamReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaCreateTeamResp": {
      "type": "object",
      "properties": {
        "team": {
          "$ref": "#/definitions/v2betaTeam"
        }
      }
    },
    "v2betaDeleteTeamResp": {
      "type": "object",
      "properties": {
        "team": {
          "$ref": "#/definitions/v2betaTeam"
        }
      }
    },
    "v2betaGetTeamMembershipResp": {
      "type": "object",
      "properties": {
        "user_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaGetTeamResp": {
      "type": "object",
      "properties": {
        "team": {
          "$ref": "#/definitions/v2betaTeam"
        }
      }
    },
    "v2betaGetTeamsForMemberResp": {
      "type": "object",
      "properties": {
        "teams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v2betaTeam"
          }
        }
      }
    },
    "v2betaListTeamsResp": {
      "type": "object",
      "properties": {
        "teams": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v2betaTeam"
          }
        }
      }
    },
    "v2betaRemoveTeamMembersReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "user_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaRemoveTeamMembersResp": {
      "type": "object",
      "properties": {
        "user_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaTeam": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaUpdateTeamReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v2betaUpdateTeamResp": {
      "type": "object",
      "properties": {
        "team": {
          "$ref": "#/definitions/v2betaTeam"
        }
      }
    }
  }
}
`)
}
