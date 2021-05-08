// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
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
        "/auth/customer": {
            "post": {
                "description": "Register new user in server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user with customer role",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login to service using given credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login to service",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/auth/owner": {
            "post": {
                "description": "Register new user in server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user with owner role",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TokenResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get User by id in token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User by id in token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update User by id in token. It can update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user by id  in token",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Register new user in server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user with owner role",
                "deprecated": true,
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/users/restaurants/": {
            "get": {
                "description": "Get all Restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "Get all Restaurant",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RestaurantResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "Create Restaurant",
                "parameters": [
                    {
                        "description": "Details of restaurant",
                        "name": "restaurant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateRestaurantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RestaurantResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/users/restaurants/:id/menus/": {
            "get": {
                "description": "GetMenu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants",
                    "Menus"
                ],
                "summary": "GetMenu Get menu using restaurantId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MenuResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "AddFoodToMenu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants",
                    "Menus"
                ],
                "summary": "AddFoodToMenu Adds Food to Menu using RestaurantId and Food",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Details of food",
                        "name": "restaurant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.AddFoodToMenuRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MenuResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/users/restaurants/:id/menus/:foodId/": {
            "delete": {
                "description": "RemoveFoodFromMenu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants",
                    "Menus"
                ],
                "summary": "RemoveFoodFromMenu Removes food from Menu using restaurantId and foodId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Id of food",
                        "name": "foodId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MenuResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/users/restaurants/{id}": {
            "get": {
                "description": "Get Restaurant by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "Get Restaurant by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RestaurantResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "Update Restaurant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Details of restaurant",
                        "name": "restaurant",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateRestaurantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RestaurantResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove restaurant from user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "Remove restaurant from user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserRestaurantResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/users/restaurants/{id}/address": {
            "put": {
                "description": "Replace address Restaurant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restaurants"
                ],
                "summary": "Replace address Restaurant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of restaurant",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Details of address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.AddAddressToRestaurantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.RestaurantAddressResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.AddAddressToRestaurantRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/requests.Address"
                }
            }
        },
        "requests.AddFoodToMenuRequest": {
            "type": "object",
            "properties": {
                "food": {
                    "type": "object",
                    "properties": {
                        "description": {
                            "type": "string"
                        },
                        "food_type": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        },
                        "price": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "requests.Address": {
            "type": "object",
            "properties": {
                "address_line_1": {
                    "type": "string"
                },
                "address_line_2": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "post_code": {
                    "type": "string"
                }
            }
        },
        "requests.CreateRestaurantRequest": {
            "type": "object",
            "properties": {
                "restaurant": {
                    "type": "object",
                    "required": [
                        "name"
                    ],
                    "properties": {
                        "address": {
                            "$ref": "#/definitions/requests.Address"
                        },
                        "description": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "requests.UpdateRestaurantRequest": {
            "type": "object",
            "properties": {
                "restaurant": {
                    "type": "object",
                    "required": [
                        "name"
                    ],
                    "properties": {
                        "description": {
                            "type": "string"
                        },
                        "name": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "requests.UserLoginRequest": {
            "type": "object",
            "properties": {
                "user": {
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
                }
            }
        },
        "requests.UserRegisterRequest": {
            "type": "object",
            "properties": {
                "user": {
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
                        },
                        "username": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "responses.FoodResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "food_id": {
                    "type": "string"
                },
                "food_type": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "responses.MenuResponse": {
            "type": "object",
            "properties": {
                "foods": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.FoodResponse"
                    }
                },
                "menuId": {
                    "type": "string"
                }
            }
        },
        "responses.RestaurantAddressResponse": {
            "type": "object",
            "properties": {
                "addressID": {
                    "type": "string"
                },
                "restaurantID": {
                    "type": "string"
                }
            }
        },
        "responses.RestaurantResponse": {
            "type": "object",
            "properties": {
                "restaurant": {
                    "type": "object",
                    "required": [
                        "name"
                    ],
                    "properties": {
                        "description": {
                            "type": "string"
                        },
                        "menu": {
                            "$ref": "#/definitions/responses.MenuResponse"
                        },
                        "name": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "responses.TokenResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "responses.UserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "object",
                    "properties": {
                        "email": {
                            "type": "string"
                        },
                        "role": {
                            "type": "integer"
                        },
                        "username": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "responses.UserRestaurantResponse": {
            "type": "object",
            "properties": {
                "restaurantID": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "utils.Error": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        },
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Version:     "0.1.0",
	Host:        "localhost:8585",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Restaurant backend API server",
	Description: "This is a Restaurant backend RESTful API documentation.",
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
	swag.Register(swag.Name, &s{})
}
