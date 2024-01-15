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
        "/agent/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new agency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agency"
                ],
                "summary": "Create a new agency",
                "parameters": [
                    {
                        "description": "Agency",
                        "name": "agency",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Agency.CreateAgency"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    }
                }
            }
        },
        "/agent/list": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get agent list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agency"
                ],
                "summary": "Get agent list",
                "parameters": [
                    {
                        "description": "Agency",
                        "name": "agency",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/utils.RequestObj"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    }
                }
            }
        },
        "/agent/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agency"
                ],
                "summary": "Get agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agency"
                ],
                "summary": "Update agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Agency",
                        "name": "agency",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Agency.UpdateAgency"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Agency"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseObj"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Sign in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign in user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/user.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Agency.CreateAgency": {
            "type": "object",
            "required": [
                "address",
                "city",
                "description",
                "email",
                "name",
                "phone",
                "website"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "Agency.UpdateAgency": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "user.ErrorResponse": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "user.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.SignUpInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "passwordConfirm"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "passwordConfirm": {
                    "type": "string",
                    "minLength": 8
                },
                "photo": {
                    "type": "string"
                }
            }
        },
        "user.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "provider": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "utils.DefaultParam": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "utils.FilterObj": {
            "type": "object",
            "properties": {
                "field_name": {
                    "type": "string"
                },
                "field_type": {
                    "type": "string"
                },
                "operation": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                },
                "values": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "utils.Pagination": {
            "type": "object",
            "properties": {
                "current_page_no": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "total_elements": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "utils.RequestObj": {
            "type": "object",
            "properties": {
                "default_param": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/utils.DefaultParam"
                    }
                },
                "filter": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/utils.FilterObj"
                    }
                },
                "page_no": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseObj": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "pagination": {
                    "$ref": "#/definitions/utils.Pagination"
                },
                "response_code": {
                    "type": "integer"
                },
                "response_msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://tranquil-river-84673-f93f313aee4e.herokuapp.com",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Ramen API",
	Description:      "This is a sample API with Fiber and Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}