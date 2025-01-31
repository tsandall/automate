package api

func init() {
	Swagger.Add("authz_authz", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/authz/authz.proto",
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
    "/auth/introspect": {
      "get": {
        "operationId": "IntrospectAll",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseIntrospectResp"
            }
          }
        },
        "tags": [
          "Authorization"
        ]
      },
      "post": {
        "operationId": "Introspect",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseIntrospectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestIntrospectReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    },
    "/auth/introspect_some": {
      "post": {
        "operationId": "IntrospectSome",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseIntrospectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestIntrospectSomeReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    },
    "/auth/policies": {
      "get": {
        "operationId": "ListPolicies",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseListPoliciesResp"
            }
          }
        },
        "tags": [
          "Authorization"
        ]
      },
      "post": {
        "operationId": "CreatePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseCreatePolicyResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestCreatePolicyReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    },
    "/auth/policies/version": {
      "get": {
        "operationId": "GetVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/versionVersionInfo"
            }
          }
        },
        "tags": [
          "Authorization"
        ]
      }
    },
    "/auth/policies/{id}": {
      "delete": {
        "operationId": "DeletePolicy",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseDeletePolicyResp"
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
          "Authorization"
        ]
      }
    }
  },
  "definitions": {
    "requestCreatePolicyReq": {
      "type": "object",
      "properties": {
        "action": {
          "type": "string"
        },
        "subjects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "resource": {
          "type": "string"
        }
      }
    },
    "requestIntrospectReq": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string"
        },
        "parameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "requestIntrospectSomeReq": {
      "type": "object",
      "properties": {
        "paths": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "responseCreatePolicyResp": {
      "type": "object",
      "properties": {
        "action": {
          "type": "string"
        },
        "subjects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string"
        },
        "resource": {
          "type": "string"
        },
        "effect": {
          "type": "string"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        }
      },
      "description": "We aren't using a Policy message here since we want to\nreturn a flat object via our external HTTP API."
    },
    "responseDeletePolicyResp": {
      "type": "object",
      "properties": {
        "action": {
          "type": "string"
        },
        "subjects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string"
        },
        "resource": {
          "type": "string"
        },
        "effect": {
          "type": "string"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "responseIntrospectResp": {
      "type": "object",
      "properties": {
        "endpoints": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/responseMethodsAllowed"
          }
        }
      }
    },
    "responseListPoliciesResp": {
      "type": "object",
      "properties": {
        "policies": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/responsePolicy"
          }
        }
      }
    },
    "responseMethodsAllowed": {
      "type": "object",
      "properties": {
        "get": {
          "type": "boolean",
          "format": "boolean"
        },
        "put": {
          "type": "boolean",
          "format": "boolean"
        },
        "post": {
          "type": "boolean",
          "format": "boolean"
        },
        "delete": {
          "type": "boolean",
          "format": "boolean"
        },
        "patch": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "responsePolicy": {
      "type": "object",
      "properties": {
        "action": {
          "type": "string"
        },
        "subjects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string"
        },
        "resource": {
          "type": "string"
        },
        "effect": {
          "type": "string"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "versionVersionInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "sha": {
          "type": "string"
        },
        "built": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
