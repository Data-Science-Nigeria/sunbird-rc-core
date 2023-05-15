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
  "swagger": "2.0",
  "info": {
    "description": "Bulk Issuance API",
    "title": "Bulk Issuance",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/uploads": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "uploadedFiles"
        ],
        "summary": "get uploaded files",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/v1/{id}/report": {
      "get": {
        "produces": [
          "application/json",
          "application/octet-stream"
        ],
        "tags": [
          "downloadFileReport"
        ],
        "summary": "download the success and error report of file uploaded",
        "parameters": [
          {
            "type": "integer",
            "description": "File name",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/FileDownload"
            },
            "headers": {
              "Content-Disposition": {
                "type": "string"
              }
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/v1/{schemaName}/sample-csv": {
      "get": {
        "produces": [
          "application/octet-stream",
          "application/json"
        ],
        "tags": [
          "sampleTemplate"
        ],
        "summary": "get sample template",
        "parameters": [
          {
            "type": "string",
            "description": "schema name",
            "name": "schemaName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/SampleTemplateResponse"
            },
            "headers": {
              "Content-Disposition": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/v1/{schemaName}/upload": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "uploadAndCreateRecords"
        ],
        "summary": "upload the file and create records",
        "parameters": [
          {
            "type": "file",
            "description": "Certification data in the form of csv",
            "name": "file",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Schema name you're uploading for",
            "name": "schemaName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/CreateRecordResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateRecordResponse": {
      "type": "object"
    },
    "FileDownload": {
      "type": "object"
    },
    "Group": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "JWTClaimBody": {
      "type": "object",
      "properties": {
        "preferred_username": {
          "type": "string"
        },
        "realm_access": {
          "type": "object",
          "properties": {
            "roles": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        },
        "resource_access": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/Group"
          }
        },
        "scope": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
        },
        "user_id": {
          "type": "string"
        }
      }
    },
    "SampleTemplateResponse": {
      "type": "object"
    },
    "UploadedFiles": {
      "type": "object"
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://dummy.oauth.net/auth",
      "tokenUrl": "https://dumy.oauth.net/token",
      "scopes": {
        "Issuer": "scope of Issuer"
      }
    }
  },
  "security": [
    {
      "hasRole": [
        "Issuer"
      ]
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
  "swagger": "2.0",
  "info": {
    "description": "Bulk Issuance API",
    "title": "Bulk Issuance",
    "version": "1.0.0"
  },
  "paths": {
    "/v1/uploads": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "uploadedFiles"
        ],
        "summary": "get uploaded files",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/v1/{id}/report": {
      "get": {
        "produces": [
          "application/json",
          "application/octet-stream"
        ],
        "tags": [
          "downloadFileReport"
        ],
        "summary": "download the success and error report of file uploaded",
        "parameters": [
          {
            "type": "integer",
            "description": "File name",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/FileDownload"
            },
            "headers": {
              "Content-Disposition": {
                "type": "string"
              }
            }
          },
          "403": {
            "description": "Forbidden",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/v1/{schemaName}/sample-csv": {
      "get": {
        "produces": [
          "application/json",
          "application/octet-stream"
        ],
        "tags": [
          "sampleTemplate"
        ],
        "summary": "get sample template",
        "parameters": [
          {
            "type": "string",
            "description": "schema name",
            "name": "schemaName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/SampleTemplateResponse"
            },
            "headers": {
              "Content-Disposition": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/v1/{schemaName}/upload": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "uploadAndCreateRecords"
        ],
        "summary": "upload the file and create records",
        "parameters": [
          {
            "type": "file",
            "description": "Certification data in the form of csv",
            "name": "file",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Schema name you're uploading for",
            "name": "schemaName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/CreateRecordResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "404": {
            "description": "Not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateRecordResponse": {
      "type": "object"
    },
    "FileDownload": {
      "type": "object"
    },
    "Group": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "JWTClaimBody": {
      "type": "object",
      "properties": {
        "preferred_username": {
          "type": "string"
        },
        "realm_access": {
          "type": "object",
          "properties": {
            "roles": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        },
        "resource_access": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/Group"
          }
        },
        "scope": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
        },
        "user_id": {
          "type": "string"
        }
      }
    },
    "JWTClaimBodyRealmAccess": {
      "type": "object",
      "properties": {
        "roles": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "SampleTemplateResponse": {
      "type": "object"
    },
    "UploadedFiles": {
      "type": "object"
    }
  },
  "securityDefinitions": {
    "hasRole": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://dummy.oauth.net/auth",
      "tokenUrl": "https://dumy.oauth.net/token",
      "scopes": {
        "Issuer": "scope of Issuer"
      }
    }
  },
  "security": [
    {
      "hasRole": [
        "Issuer"
      ]
    }
  ]
}`))
}
