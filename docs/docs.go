// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthcheck": {
            "get": {
                "description": "Health checking for the service",
                "produces": [
                    "text/plain"
                ],
                "operationId": "HealthCheckHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get Users",
                "produces": [
                    "application/json"
                ],
                "operationId": "UsersHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Users"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Response Code",
                    "type": "string"
                },
                "error": {
                    "description": "Response Error",
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "UserId",
                    "type": "integer"
                },
                "name": {
                    "description": "UserName",
                    "type": "string"
                }
            }
        },
        "entities.Users": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Users",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.User"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GM Startd API",
	Description:      "GM Startd API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
