// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/media": {
            "get": {
                "description": "get medium for the given page and limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Show All Media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Medium"
                        }
                    }
                }
            },
            "post": {
                "description": "create the media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Create New Media",
                "parameters": [
                    {
                        "description": "Create Media Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateMediaDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Media"
                        }
                    }
                }
            }
        },
        "/media/{id}": {
            "get": {
                "description": "get the media for the given id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Show The Media By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Media"
                        }
                    }
                }
            },
            "put": {
                "description": "update the media by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Update The Media By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Media Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.UpdateMediaDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Media"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the media by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Delete The Media By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "203": {
                        "description": "Non-Authoritative Information"
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateMediaDto": {
            "type": "object",
            "required": [
                "homepage",
                "name",
                "thumbnail"
            ],
            "properties": {
                "homepage": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                }
            }
        },
        "dtos.UpdateMediaDto": {
            "type": "object",
            "properties": {
                "homepage": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                }
            }
        },
        "models.Media": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "homepage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Medium": {
            "type": "object",
            "properties": {
                "medium": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Media"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/types.MongoPaginate"
                }
            }
        },
        "types.MongoPaginate": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer",
                    "minimum": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "News API",
	Description:      "This is an server to manage news from mongoDB.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
