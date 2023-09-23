// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Maintainer Chen Ke",
            "url": "https://danxi.fduhole.com/about",
            "email": "dev@fduhole.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/divisions": {
            "get": {
                "description": "list all divisions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Division"
                ],
                "summary": "list all divisions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.DivisionResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            },
            "post": {
                "description": "create a division, only admin can create",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Division"
                ],
                "summary": "create a division",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.DivisionCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.DivisionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            }
        },
        "/divisions/{id}": {
            "get": {
                "description": "get a division",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Division"
                ],
                "summary": "get a division",
                "parameters": [
                    {
                        "type": "string",
                        "description": "division id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.DivisionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            },
            "put": {
                "description": "modify a division, only admin can modify",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Division"
                ],
                "summary": "modify a division",
                "parameters": [
                    {
                        "type": "string",
                        "description": "division id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.DivisionModifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.DivisionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a division, only admin can delete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Division"
                ],
                "summary": "delete a division",
                "parameters": [
                    {
                        "type": "string",
                        "description": "division id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.DivisionDeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            }
        },
        "/group/{id}": {
            "get": {
                "description": "get a course group, old version or v1 version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CourseGroup"
                ],
                "summary": "/group/{group_id}",
                "deprecated": true,
                "parameters": [
                    {
                        "type": "string",
                        "description": "course group id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CourseGroupV1Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "LoginRequest",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            }
        },
        "/register": {
            "put": {
                "description": "reset password with email, password and optional verification code if enabled",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "reset password",
                "parameters": [
                    {
                        "description": "ResetPasswordRequest",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            },
            "post": {
                "description": "register with email, password and optional verification code if enabled",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "register",
                "parameters": [
                    {
                        "description": "RegisterRequest",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schema.HttpBaseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.CourseGroupV1Response": {
            "type": "object",
            "properties": {
                "campus_name": {
                    "description": "开课校区",
                    "type": "string"
                },
                "code": {
                    "description": "课程组编号",
                    "type": "string"
                },
                "courses": {
                    "description": "课程组下的课程，slices 必须非空",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.CourseV1Response"
                    }
                },
                "department": {
                    "description": "开课学院",
                    "type": "string"
                },
                "id": {
                    "description": "课程组 ID",
                    "type": "integer"
                },
                "name": {
                    "description": "课程组名称",
                    "type": "string"
                }
            }
        },
        "schema.CourseV1Response": {
            "type": "object",
            "properties": {
                "campus_name": {
                    "description": "开课校区",
                    "type": "string"
                },
                "code": {
                    "description": "课程编号",
                    "type": "string"
                },
                "code_id": {
                    "description": "选课序号。用于区分同一课程编号的不同平行班",
                    "type": "string"
                },
                "credit": {
                    "description": "学分",
                    "type": "integer"
                },
                "department": {
                    "description": "开课学院",
                    "type": "string"
                },
                "id": {
                    "description": "课程 ID",
                    "type": "integer"
                },
                "max_student": {
                    "description": "最大选课人数",
                    "type": "integer"
                },
                "name": {
                    "description": "课程名称",
                    "type": "string"
                },
                "semester": {
                    "description": "学期",
                    "type": "integer"
                },
                "teachers": {
                    "description": "老师：多个老师用逗号分隔",
                    "type": "string"
                },
                "week_hour": {
                    "description": "周学时",
                    "type": "integer"
                },
                "year": {
                    "description": "学年",
                    "type": "integer"
                }
            }
        },
        "schema.DivisionCreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "分区详情：前端暂时不用",
                    "type": "string"
                },
                "name": {
                    "description": "分区名称: 树洞、评教等等",
                    "type": "string"
                }
            }
        },
        "schema.DivisionDeleteRequest": {
            "type": "object",
            "properties": {
                "to": {
                    "description": "ID of the target division that all the deleted division's holes will be moved to",
                    "type": "integer",
                    "default": 1
                }
            }
        },
        "schema.DivisionModifyRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "分区详情：前端暂时不用",
                    "type": "string"
                },
                "name": {
                    "description": "分区名称: 树洞、评教等等",
                    "type": "string"
                },
                "pinned": {
                    "description": "TODO: 置顶的树洞 id",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "schema.DivisionResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "分区详情：前端暂时不用",
                    "type": "string"
                },
                "division_id": {
                    "description": "旧版 id",
                    "type": "integer"
                },
                "id": {
                    "description": "新版 id",
                    "type": "integer"
                },
                "name": {
                    "description": "分区名称: 树洞、评教等等",
                    "type": "string"
                },
                "pinned": {
                    "description": "TODO: 置顶的树洞",
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "schema.HttpBaseError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schema.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "validation_detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.ValidateFieldError"
                    }
                }
            }
        },
        "schema.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "email in email blacklist\nTODO: add email blacklist",
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                }
            }
        },
        "schema.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "email in email blacklist\nTODO: add email blacklist",
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "verification": {
                    "type": "string"
                }
            }
        },
        "schema.ResetPasswordRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "email in email blacklist\nTODO: add email blacklist",
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "verification": {
                    "type": "string"
                }
            }
        },
        "schema.TokenResponse": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "refresh": {
                    "type": "string"
                }
            }
        },
        "schema.ValidateFieldError": {
            "type": "object",
            "properties": {
                "field": {
                    "description": "Field is the field name that failed validation\nuse registered tag name if registered",
                    "type": "string"
                },
                "message": {
                    "description": "Message is the error message",
                    "type": "string"
                },
                "param": {
                    "description": "Param is the parameter for the validation",
                    "type": "string"
                },
                "tag": {
                    "description": "Tag is the validation tag that failed.\nuse alias if defined\n\ne.g. \"required\", \"min\", \"max\", etc.",
                    "type": "string"
                },
                "value": {
                    "description": "Value is the actual value that failed validation"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "3.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Open Tree Hole Auth",
	Description:      "Next Generation of Auth microservice integrated with kong for registration and issuing tokens",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
