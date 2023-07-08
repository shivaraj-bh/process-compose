// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/hostname": {
            "get": {
                "description": "Get process compose hostname",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hostname"
                ],
                "summary": "Get Hostname",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/live": {
            "get": {
                "description": "Check if server is responding",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Liveness"
                ],
                "summary": "Liveness Check",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/process/info/{name}": {
            "get": {
                "description": "Retrieves the given process and its config",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get process config",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Process Config",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/process/logs/{name}/{endOffset}/{limit}": {
            "get": {
                "description": "Retrieves the process logs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get process logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset from the end of the log",
                        "name": "endOffset",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit of lines to get (0 will get all the lines till the end)",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Process Logs",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/process/ports/{name}": {
            "get": {
                "description": "Retrieves process open ports",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get process ports",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Process Ports",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/process/restart/{name}": {
            "post": {
                "description": "Restarts the process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Restart a process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Restarted Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process/scale/{name}/{scale}": {
            "patch": {
                "description": "Scale a process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Scale a process to a given replicas count",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "New amount of process replicas",
                        "name": "scale",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Scaled Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process/scale/{name}/{scale}": {
            "patch": {
                "description": "Scale a process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Scale a process to a given replicas count",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "New amount of process replicas",
                        "name": "scale",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Scaled Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process/start/{name}": {
            "post": {
                "description": "Starts the process if the state is not 'running' or 'pending'",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Start a process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Started Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process/stop/{name}": {
            "patch": {
                "description": "Sends kill signal to the process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Stop a process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Stopped Process Name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process/{name}": {
            "get": {
                "description": "Retrieves the given process and its status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get process state",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Process Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Process State",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/processes": {
            "get": {
                "description": "Retrieves all the configured processes and their status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Process"
                ],
                "summary": "Get all processes",
                "responses": {
                    "200": {
                        "description": "Processes Status",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
