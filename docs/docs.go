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
        "/add_product": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addProduct"
                ],
                "summary": "addProduct",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductCreateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Added product",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/add_user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addUser"
                ],
                "summary": "addUser",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.UserAddParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Added user",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/create_order": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderCreate"
                ],
                "summary": "Order Create",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateOrderParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "create order",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/del_order": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deleteOrder"
                ],
                "summary": "delete order by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete order by id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/del_product": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deleteProduct"
                ],
                "summary": "delete product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete product",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/del_user": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deleteUserByName"
                ],
                "summary": "delete User by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete user by name",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/edit_product": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "updateProductPrice"
                ],
                "summary": "update product price",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductUpdateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update product",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/edit_user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "updateUser"
                ],
                "summary": "updateUser",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.UserUpdateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update user name",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order_add_product": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OrderAddProduct"
                ],
                "summary": "Order add product",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductAddOrderParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "create order",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order_recount": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orderRecount"
                ],
                "summary": "order recount",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.OrderRecountParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order after recount",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order_update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orderUpdateQuantity"
                ],
                "summary": "order update",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.OrderProductUpdateParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.ShopOrder"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getOrders"
                ],
                "summary": "get orders",
                "responses": {
                    "200": {
                        "description": "Get orders",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/orders_products": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getOrdersFull"
                ],
                "summary": "get orders full",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.OrdersProductsFullRow"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getProductByName"
                ],
                "summary": "get Product by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get product by name",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getProducts"
                ],
                "summary": "get products",
                "responses": {
                    "200": {
                        "description": "Get products",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/products_prices": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getProductsForEange"
                ],
                "summary": "get Products for range",
                "parameters": [
                    {
                        "description": "Модель которую принимает метод",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.ProductGetRangePriceParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "get products",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/statistic": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getUsersStatistic"
                ],
                "summary": "get users statistic",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repository.UsersStatisticRow"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getUserByParameters"
                ],
                "summary": "get User by parameters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get user by parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getUsers"
                ],
                "summary": "get users",
                "responses": {
                    "200": {
                        "description": "Get users",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateOrderParams": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.Product"
                    }
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "handler.OrderProductUpdateParams": {
            "type": "object",
            "properties": {
                "order_id": {
                    "description": "nolint: tagliatelle",
                    "type": "integer"
                },
                "product_id": {
                    "description": "nolint: tagliatelle",
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                }
            }
        },
        "handler.OrderRecountParams": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "integer"
                }
            }
        },
        "handler.OrdersProductsFullRow": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "order_id": {
                    "description": "nolint: tagliatelle",
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                }
            }
        },
        "handler.Product": {
            "type": "object",
            "properties": {
                "product": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                }
            }
        },
        "handler.ProductAddOrderParams": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "orderId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                }
            }
        },
        "handler.ProductCreateParams": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "handler.ProductGetRangePriceParams": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number"
                },
                "price_2": {
                    "description": "nolint: tagliatelle",
                    "type": "number"
                }
            }
        },
        "handler.ProductUpdateParams": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "handler.ShopOrder": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "order_date": {
                    "description": "nolint: tagliatelle",
                    "type": "string"
                },
                "total_amount": {
                    "description": "nolint: tagliatelle",
                    "type": "number"
                },
                "user_id": {
                    "description": "nolint: tagliatelle",
                    "type": "integer"
                }
            }
        },
        "repository.UserAddParams": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "repository.UserUpdateParams": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "name_2": {
                    "type": "string"
                }
            }
        },
        "repository.UsersStatisticRow": {
            "type": "object",
            "properties": {
                "avr price": {
                    "type": "number"
                },
                "total_orders": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "0.0.0.0:8080/",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "HW15: shop",
	Description:      "API Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
