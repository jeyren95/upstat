// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `
{
  "openapi": "3.0.0",
  "info": {
    "title": "Upstat API",
    "description": "This is an auto-generated API Docs for Upstat API.",
    "version": "1.0"
  },
  "servers": [
    {
      "url": "/",
      "description": "Default Server URL"
    }
  ],
  "paths": {
    "/api/auth/refresh-token": {
      "post": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "description": " Renew access and refresh tokens."
      }
    },
    "/api/auth/signin": {
      "post": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UserSignInOut"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "description": " Auth user and return access and refresh token.",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UserSignIn"
              }
            }
          },
          "required": true
        }
      }
    },
    "/api/auth/signout": {
      "post": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "description": " De-authorize user and delete refresh token from Redis."
      }
    },
    "/api/auth/signup": {
      "post": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UserSignInOut"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "description": " Create a new user.",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UserSignUp"
              }
            }
          },
          "required": true
        }
      }
    },
    "/api/monitors/create": {
      "post": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/AddMonitorIn"
              }
            }
          },
          "required": true
        }
      }
    },
    "/api/monitors/delete/{id}": {
      "delete": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          }
        ]
      }
    },
    "/api/monitors/heartbeat/{id}": {
      "get": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/HeartbeatsOut"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          },
          {
            "name": "startTime",
            "in": "query",
            "description": "Start Time",
            "required": true,
            "schema": {
              "type": "string",
              "format": "date-time"
            }
          }
        ]
      }
    },
    "/api/monitors/info/{id}": {
      "get": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MonitorInfoOut"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          }
        ]
      }
    },
    "/api/monitors/list": {
      "get": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MonitorsListOut"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        }
      }
    },
    "/api/monitors/pause/{id}": {
      "put": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          }
        ]
      }
    },
    "/api/monitors/resume/{id}": {
      "put": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          }
        ]
      }
    },
    "/api/monitors/summary/{id}": {
      "get": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/MonitorSummary"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          }
        ]
      }
    },
    "/api/monitors/update/{id}": {
      "put": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "description": "Monitor ID",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Monitor ID"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/AddMonitorIn"
              }
            }
          },
          "required": true
        }
      }
    },
    "/api/users/setup": {
      "get": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/NeedSetup"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        }
      }
    },
    "/api/users/update-password": {
      "post": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UpdatePasswordIn"
              }
            }
          },
          "required": true
        }
      }
    },
    "/api/users/{username}": {
      "patch": {
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/SuccessResponse"
                }
              }
            }
          },
          "400": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/ErrorResponse"
                }
              }
            }
          }
        },
        "parameters": [
          {
            "name": "username",
            "in": "path",
            "description": "Username",
            "required": true,
            "schema": {
              "type": "string",
              "format": "string",
              "description": "Username"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UpdateAccountIn"
              }
            }
          },
          "required": true
        }
      }
    }
  },
  "components": {
    "schemas": {
      "AddMonitorIn": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string"
          },
          "url": {
            "type": "string"
          },
          "type": {
            "type": "string"
          },
          "frequency": {
            "type": "integer"
          },
          "method": {
            "type": "string"
          }
        }
      },
      "ErrorResponse": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      },
      "Heartbeat": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer"
          },
          "monitor_id": {
            "type": "integer"
          },
          "timestamp": {
            "type": "string",
            "format": "date-time"
          },
          "status_code": {
            "type": "string"
          },
          "status": {
            "type": "string"
          },
          "latency": {
            "type": "integer"
          },
          "message": {
            "type": "string"
          }
        }
      },
      "HeartbeatsOut": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          },
          "heartbeat": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Heartbeat"
            }
          }
        }
      },
      "Monitor": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer"
          },
          "name": {
            "type": "string"
          },
          "url": {
            "type": "string"
          },
          "type": {
            "type": "string"
          },
          "method": {
            "type": "string"
          },
          "frequency": {
            "type": "integer"
          },
          "status": {
            "type": "string"
          }
        }
      },
      "MonitorInfoOut": {
        "type": "object",
        "properties": {
          "monitor": {
            "type": "object",
            "$ref": "#/components/schemas/Monitor"
          },
          "message": {
            "type": "string"
          }
        }
      },
      "MonitorItem": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string"
          },
          "name": {
            "type": "string"
          },
          "url": {
            "type": "string"
          },
          "frequency": {
            "type": "integer"
          },
          "status": {
            "type": "string"
          },
          "heartbeat": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "integer"
                },
                "monitor_id": {
                  "type": "integer"
                },
                "timestamp": {
                  "type": "string",
                  "format": "date-time"
                },
                "status_code": {
                  "type": "string"
                },
                "status": {
                  "type": "string"
                },
                "latency": {
                  "type": "integer"
                },
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "MonitorSummary": {
        "type": "object",
        "properties": {
          "averageLatency": {
            "type": "number"
          },
          "dayUptime": {
            "type": "number"
          },
          "monthUptime": {
            "type": "number"
          }
        }
      },
      "MonitorsListOut": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          },
          "monitors": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "url": {
                  "type": "string"
                },
                "frequency": {
                  "type": "integer"
                },
                "status": {
                  "type": "string"
                },
                "heartbeat": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "id": {
                        "type": "integer"
                      },
                      "monitor_id": {
                        "type": "integer"
                      },
                      "timestamp": {
                        "type": "string",
                        "format": "date-time"
                      },
                      "status_code": {
                        "type": "string"
                      },
                      "status": {
                        "type": "string"
                      },
                      "latency": {
                        "type": "integer"
                      },
                      "message": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      },
      "NeedSetup": {
        "type": "object",
        "properties": {
          "needSetup": {
            "type": "boolean"
          }
        }
      },
      "SuccessResponse": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      },
      "UpdateAccountIn": {
        "type": "object",
        "properties": {
          "firstname": {
            "type": "string"
          },
          "lastname": {
            "type": "string"
          }
        }
      },
      "UpdatePasswordIn": {
        "type": "object",
        "properties": {
          "currentPassword": {
            "type": "string"
          },
          "newPassword": {
            "type": "string"
          }
        }
      },
      "User": {
        "type": "object",
        "properties": {
          "username": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "firstname": {
            "type": "string"
          },
          "lastname": {
            "type": "string"
          }
        }
      },
      "UserSignIn": {
        "type": "object",
        "properties": {
          "username": {
            "type": "string"
          },
          "password": {
            "type": "string"
          }
        }
      },
      "UserSignInOut": {
        "type": "object",
        "properties": {
          "user": {
            "type": "object",
            "$ref": "#/components/schemas/User"
          },
          "message": {
            "type": "string"
          }
        }
      },
      "UserSignUp": {
        "type": "object",
        "properties": {
          "username": {
            "type": "string"
          },
          "email": {
            "type": "string"
          },
          "password": {
            "type": "string"
          }
        }
      },
      "serializers.AddMonitorIn": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string"
          },
          "url": {
            "type": "string"
          },
          "type": {
            "type": "string"
          },
          "frequency": {
            "type": "integer"
          },
          "method": {
            "type": "string"
          }
        }
      },
      "serializers.ErrorResponse": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      },
      "serializers.SuccessResponse": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      },
      "serializers.UserSignInOut": {
        "type": "object",
        "properties": {
          "user": {
            "type": "object",
            "$ref": "#/components/schemas/User"
          },
          "message": {
            "type": "string"
          }
        }
      }
    }
  }
}
`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Upstat API",
	Description:      "This is an auto-generated API Docs for Upstat API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
