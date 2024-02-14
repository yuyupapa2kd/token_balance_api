// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/coin/{address}": {
            "get": {
                "description": "Get Coin Balance of given Address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coin"
                ],
                "summary": "Get Coin Balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/resource.ResJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/resource.ResBalance"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/resource.ResJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/resource.ResErr"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/token/{symbol}/{address}": {
            "get": {
                "description": "Get Token Balance of given Address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Allowed Token Symbol : SOP, LOUI, ksETH, ksUSDT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/resource.ResJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/resource.ResBalance"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/resource.ResJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/resource.ResErr"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "resource.ResBalance": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                }
            }
        },
        "resource.ResErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "resource.ResJSON": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
