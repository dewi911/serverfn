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
        "/task": {
            "get": {
                "description": "Get details of all tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get all tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Task"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new task with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Create task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.TaskCreateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.TaskResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "get": {
                "description": "Get details of a task by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get a task by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Task"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a task's status by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.TaskUpdateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a task by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Headers": {
            "type": "object",
            "properties": {
                "HTTPStatusCode": {
                    "type": "integer"
                },
                "authentication": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "headers": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "responseHeaders": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                },
                "responseLength": {
                    "type": "integer"
                }
            }
        },
        "domain.Task": {
            "type": "object",
            "properties": {
                "headers": {
                    "$ref": "#/definitions/domain.Headers"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "method": {
                    "type": "string",
                    "example": "GET"
                },
                "task_status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/domain.TaskStatus"
                        }
                    ],
                    "example": "new"
                },
                "url": {
                    "type": "string",
                    "example": "http://google.com"
                }
            }
        },
        "domain.TaskCreateInput": {
            "type": "object",
            "properties": {
                "headers": {
                    "$ref": "#/definitions/domain.Headers"
                },
                "method": {
                    "type": "string",
                    "example": "GET"
                },
                "task_status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/domain.TaskStatus"
                        }
                    ],
                    "example": "new"
                },
                "url": {
                    "type": "string",
                    "example": "http://google.com"
                }
            }
        },
        "domain.TaskResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "domain.TaskStatus": {
            "type": "string",
            "enum": [
                "new",
                "in_process",
                "done",
                "error"
            ],
            "x-enum-varnames": [
                "TaskStatusNew",
                "TaskStatusInProcess",
                "TaskStatusDone",
                "TaskStatusError"
            ]
        },
        "domain.TaskUpdateInput": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "$ref": "#/definitions/domain.TaskStatus"
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
	Title:            "Server API",
	Description:      "This is a server API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
