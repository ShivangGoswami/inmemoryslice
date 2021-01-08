// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Nokia Rest api",
    "title": "Nokia Rest api",
    "contact": {
      "name": "Shivang Goswami",
      "email": "shivang.goswami@outlook.com"
    },
    "version": "1.0"
  },
  "basePath": "/v1/rest",
  "paths": {
    "/add": {
      "post": {
        "description": "store data",
        "tags": [
          "nokiaapi"
        ],
        "summary": "Store Data",
        "parameters": [
          {
            "name": "input",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/inputParam"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "status",
            "schema": {
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        },
        "x-visibility": "public"
      }
    },
    "/delete": {
      "delete": {
        "description": "Delete data",
        "tags": [
          "nokiaapi"
        ],
        "summary": "Delete Data",
        "parameters": [
          {
            "name": "input",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/inputParam"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "status",
            "schema": {
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        },
        "x-visibility": "public"
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "http error code",
          "type": "integer",
          "format": "int64",
          "x-example": 401
        },
        "message": {
          "description": "http error message",
          "type": "string",
          "x-example": "forbidden"
        }
      },
      "x-visibility": "public"
    },
    "inputParam": {
      "type": "object",
      "required": [
        "input"
      ],
      "properties": {
        "input": {
          "type": "string"
        }
      }
    },
    "status": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations restapi",
      "name": "nokiaapi"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Nokia Rest api",
    "title": "Nokia Rest api",
    "contact": {
      "name": "Shivang Goswami",
      "email": "shivang.goswami@outlook.com"
    },
    "version": "1.0"
  },
  "basePath": "/v1/rest",
  "paths": {
    "/add": {
      "post": {
        "description": "store data",
        "tags": [
          "nokiaapi"
        ],
        "summary": "Store Data",
        "parameters": [
          {
            "name": "input",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/inputParam"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "status",
            "schema": {
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        },
        "x-visibility": "public"
      }
    },
    "/delete": {
      "delete": {
        "description": "Delete data",
        "tags": [
          "nokiaapi"
        ],
        "summary": "Delete Data",
        "parameters": [
          {
            "name": "input",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/inputParam"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "status",
            "schema": {
              "$ref": "#/definitions/status"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        },
        "x-visibility": "public"
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "http error code",
          "type": "integer",
          "format": "int64",
          "x-example": 401
        },
        "message": {
          "description": "http error message",
          "type": "string",
          "x-example": "forbidden"
        }
      },
      "x-visibility": "public"
    },
    "inputParam": {
      "type": "object",
      "required": [
        "input"
      ],
      "properties": {
        "input": {
          "type": "string"
        }
      }
    },
    "status": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Operations restapi",
      "name": "nokiaapi"
    }
  ]
}`))
}