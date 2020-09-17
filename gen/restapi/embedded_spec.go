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
  "swagger": "2.0",
  "info": {
    "title": "Secen Pick Server",
    "version": "1.0.0"
  },
  "paths": {
    "/dialog": {
      "get": {
        "operationId": "getDialog",
        "parameters": [
          {
            "type": "string",
            "description": "““, “anime“, “manga“, “book“のうちのどれか",
            "name": "genre",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "return dialogs",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/dialog"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postDialog",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "description": "セリフ投稿時に必要なパラメータ",
            "name": "content",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "author": {
                  "type": "string"
                },
                "content": {
                  "type": "string"
                },
                "link": {
                  "type": "string"
                },
                "style": {
                  "type": "string"
                },
                "title": {
                  "type": "string"
                },
                "user_id": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/dialog/{id}/comment": {
      "get": {
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "取得成功",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/comment"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "comment": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "comment": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "user": {
          "$ref": "#/definitions/user"
        }
      }
    },
    "dialog": {
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "link": {
          "type": "string"
        },
        "style": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "user_id": {
          "type": "integer"
        }
      }
    },
    "error": {
      "type": "string"
    },
    "user": {
      "type": "object",
      "properties": {
        "display_name": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Secen Pick Server",
    "version": "1.0.0"
  },
  "paths": {
    "/dialog": {
      "get": {
        "operationId": "getDialog",
        "parameters": [
          {
            "type": "string",
            "description": "““, “anime“, “manga“, “book“のうちのどれか",
            "name": "genre",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "オフセット。初めは0",
            "name": "offset",
            "in": "query",
            "required": true
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "取得する記録の数",
            "name": "limit",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "return dialogs",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/dialog"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "operationId": "postDialog",
        "parameters": [
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "description": "セリフ投稿時に必要なパラメータ",
            "name": "content",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "author": {
                  "type": "string"
                },
                "content": {
                  "type": "string"
                },
                "link": {
                  "type": "string"
                },
                "style": {
                  "type": "string"
                },
                "title": {
                  "type": "string"
                },
                "user_id": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "投稿成功",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/dialog/{id}/comment": {
      "get": {
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "取得成功",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                },
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/comment"
                  }
                }
              }
            }
          },
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "name": "comment",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "comment": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "400": {
            "description": "request error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "internal serever error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "comment": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "user": {
          "$ref": "#/definitions/user"
        }
      }
    },
    "dialog": {
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "content": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "link": {
          "type": "string"
        },
        "style": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "user_id": {
          "type": "integer"
        }
      }
    },
    "error": {
      "type": "string"
    },
    "user": {
      "type": "object",
      "properties": {
        "display_name": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        }
      }
    }
  }
}`))
}
