package api

func init() {
	Swagger.Add("applications", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/applications/applications.proto",
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
    "/beta/applications/delete_disconnected_services": {
      "post": {
        "operationId": "DeleteDisconnectedServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesRes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/applicationsDisconnectedServicesReq"
            }
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/disconnected_services": {
      "get": {
        "operationId": "GetDisconnectedServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesRes"
            }
          }
        },
        "parameters": [
          {
            "name": "threshold_seconds",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/service-groups": {
      "get": {
        "operationId": "GetServiceGroups",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServiceGroups"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "pagination.page",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/service-groups/{service_group_id}": {
      "get": {
        "operationId": "GetServicesBySG",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesBySGRes"
            }
          }
        },
        "parameters": [
          {
            "name": "service_group_id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "pagination.page",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          },
          {
            "name": "health",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/service_groups_health_counts": {
      "get": {
        "operationId": "GetServiceGroupsHealthCounts",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsHealthCounts"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/services": {
      "get": {
        "operationId": "GetServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesRes"
            }
          }
        },
        "parameters": [
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          },
          {
            "name": "pagination.page",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "pagination.size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "sorting.field",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sorting.order",
            "in": "query",
            "required": false,
            "type": "string",
            "enum": [
              "ASC",
              "DESC"
            ],
            "default": "ASC"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/services-distinct-values": {
      "get": {
        "operationId": "GetServicesDistinctValues",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesDistinctValuesRes"
            }
          }
        },
        "parameters": [
          {
            "name": "field_name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "query_fragment",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "filter",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi"
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/stats": {
      "get": {
        "operationId": "GetServicesStats",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsServicesStatsRes"
            }
          }
        },
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/applications/version": {
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
          "ApplicationsService"
        ]
      }
    },
    "/beta/retention/service_groups/delete_disconnected_services/config": {
      "get": {
        "operationId": "GetDeleteDisconnectedServicesConfig",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsPeriodicJobConfig"
            }
          }
        },
        "tags": [
          "ApplicationsService"
        ]
      },
      "post": {
        "operationId": "UpdateDeleteDisconnectedServicesConfig",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsUpdateDeleteDisconnectedServicesConfigRes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/applicationsPeriodicJobConfig"
            }
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    },
    "/beta/retention/service_groups/disconnected_services/config": {
      "get": {
        "operationId": "GetDisconnectedServicesConfig",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsPeriodicMandatoryJobConfig"
            }
          }
        },
        "tags": [
          "ApplicationsService"
        ]
      },
      "post": {
        "operationId": "UpdateDisconnectedServicesConfig",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/applicationsUpdateDisconnectedServicesConfigRes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/applicationsPeriodicMandatoryJobConfig"
            }
          }
        ],
        "tags": [
          "ApplicationsService"
        ]
      }
    }
  },
  "definitions": {
    "applicationsDisconnectedServicesReq": {
      "type": "object",
      "properties": {
        "threshold_seconds": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "applicationsHealthCounts": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32"
        },
        "ok": {
          "type": "integer",
          "format": "int32"
        },
        "warning": {
          "type": "integer",
          "format": "int32"
        },
        "critical": {
          "type": "integer",
          "format": "int32"
        },
        "unknown": {
          "type": "integer",
          "format": "int32"
        },
        "disconnected": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "applicationsHealthStatus": {
      "type": "string",
      "enum": [
        "OK",
        "WARNING",
        "CRITICAL",
        "UNKNOWN",
        "NONE"
      ],
      "default": "OK",
      "description": "- NONE: The representation of NO health check status\nTODO @afiune how much effort would be to change\nthe OK enum to be NONE",
      "title": "The HealthStatus enum matches the habitat implementation for health-check status:\n=\u003e https://www.habitat.sh/docs/reference/#health-check"
    },
    "applicationsPeriodicJobConfig": {
      "type": "object",
      "properties": {
        "running": {
          "type": "boolean",
          "format": "boolean"
        },
        "threshold": {
          "type": "string",
          "title": "To match the ingest API at /retention/nodes/missing-nodes/config, we use a\nstring format that is a subset of elasticsearch's date math. See the\nsimpledatemath package under lib/ for more details"
        }
      }
    },
    "applicationsPeriodicMandatoryJobConfig": {
      "type": "object",
      "properties": {
        "threshold": {
          "type": "string",
          "title": "To match the ingest API at /retention/nodes/missing-nodes/config, we use a\nstring format that is a subset of elasticsearch's date math. See the\nsimpledatemath package under lib/ for more details"
        }
      },
      "title": "it's like a PeriodicJobConfig but the user isn't allowed to change whether\nor not the job runs"
    },
    "applicationsService": {
      "type": "object",
      "properties": {
        "supervisor_id": {
          "type": "string"
        },
        "release": {
          "type": "string"
        },
        "group": {
          "type": "string"
        },
        "health_check": {
          "$ref": "#/definitions/applicationsHealthStatus"
        },
        "status": {
          "$ref": "#/definitions/applicationsServiceStatus"
        },
        "application": {
          "type": "string"
        },
        "environment": {
          "type": "string"
        },
        "fqdn": {
          "type": "string"
        },
        "channel": {
          "type": "string"
        },
        "update_strategy": {
          "type": "string"
        },
        "site": {
          "type": "string"
        },
        "previous_health_check": {
          "$ref": "#/definitions/applicationsHealthStatus"
        },
        "current_health_since": {
          "type": "string"
        },
        "health_updated_at": {
          "type": "string",
          "format": "date-time"
        },
        "disconnected": {
          "type": "boolean",
          "format": "boolean"
        },
        "last_event_occurred_at": {
          "type": "string",
          "format": "date-time"
        },
        "last_event_since": {
          "type": "string"
        }
      }
    },
    "applicationsServiceGroup": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "release": {
          "type": "string",
          "title": "Combination of the version and release in a single string like:\nExample: 0.1.0/8743278934278923"
        },
        "status": {
          "$ref": "#/definitions/applicationsHealthStatus"
        },
        "health_percentage": {
          "type": "integer",
          "format": "int32",
          "title": "The health_percentage can be a number between 0-100"
        },
        "services_health_counts": {
          "$ref": "#/definitions/applicationsHealthCounts"
        },
        "id": {
          "type": "string"
        },
        "application": {
          "type": "string"
        },
        "environment": {
          "type": "string"
        },
        "package": {
          "type": "string",
          "title": "Combination of the origin and package name in a single string like:\nExample: core/redis"
        },
        "disconnected_count": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "A service group message is the representation of one single service group that\nis internally generated by aggregating all the services"
    },
    "applicationsServiceGroups": {
      "type": "object",
      "properties": {
        "service_groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/applicationsServiceGroup"
          }
        }
      }
    },
    "applicationsServiceStatus": {
      "type": "string",
      "enum": [
        "RUNNING",
        "INITIALIZING",
        "DEPLOYING",
        "DOWN"
      ],
      "default": "RUNNING",
      "title": "The ServiceStatus enum describes the status of the service\n@afiune have we defined these states somewhere?"
    },
    "applicationsServicesBySGRes": {
      "type": "object",
      "properties": {
        "group": {
          "type": "string"
        },
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/applicationsService"
          }
        },
        "services_health_counts": {
          "$ref": "#/definitions/applicationsHealthCounts"
        }
      }
    },
    "applicationsServicesDistinctValuesRes": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "applicationsServicesRes": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/applicationsService"
          }
        }
      }
    },
    "applicationsServicesStatsRes": {
      "type": "object",
      "properties": {
        "total_service_groups": {
          "type": "integer",
          "format": "int32"
        },
        "total_services": {
          "type": "integer",
          "format": "int32"
        },
        "total_supervisors": {
          "type": "integer",
          "format": "int32"
        },
        "total_deployments": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "applicationsUpdateDeleteDisconnectedServicesConfigRes": {
      "type": "object"
    },
    "applicationsUpdateDisconnectedServicesConfigRes": {
      "type": "object"
    },
    "queryPagination": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "size": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "querySortOrder": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "querySorting": {
      "type": "object",
      "properties": {
        "field": {
          "type": "string"
        },
        "order": {
          "$ref": "#/definitions/querySortOrder"
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
