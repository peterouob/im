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
        "/index": {
            "get": {
                "tags": [
                    "首頁"
                ],
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
        "/user/createUser": {
            "get": {
                "tags": [
                    "使用者"
                ],
                "summary": "新增用戶",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "密碼",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "確認密碼",
                        "name": "repassword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "get": {
                "tags": [
                    "使用者"
                ],
                "summary": "刪除用戶",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/findUserByNameAndPwd": {
            "post": {
                "tags": [
                    "使用者"
                ],
                "summary": "登入使用者",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "tags": [
                    "使用者"
                ],
                "summary": "所有用戶",
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "tags": [
                    "使用者"
                ],
                "summary": "修改使用者",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Phone",
                        "name": "phone",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
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