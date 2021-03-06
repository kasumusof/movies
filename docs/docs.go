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
        "/character/{character_id}/": {
            "get": {
                "description": "get a particular character",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "get a particular character",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Character ID",
                        "name": "character_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CharacterResp"
                        }
                    }
                }
            }
        },
        "/movies/": {
            "get": {
                "description": "get all movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "get all movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.MoviesResp"
                        }
                    }
                }
            }
        },
        "/movies/{movie_id}/": {
            "get": {
                "description": "get a particular movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "get a particular movies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.MovieResp"
                        }
                    }
                }
            }
        },
        "/movies/{movie_id}/characters/{character_id}/": {
            "get": {
                "description": "get all  characters for a particular movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "get all  characters for a particular movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Character ID",
                        "name": "character_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CharactersResp"
                        }
                    }
                }
            }
        },
        "/movies/{movie_id}/comments/": {
            "get": {
                "description": "get all comments for a movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "get all comments for  a movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CommentsResp"
                        }
                    }
                }
            },
            "post": {
                "description": "add a comment to a movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "add a comment to a movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CommentResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "movie_id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "types.Character": {
            "type": "object",
            "properties": {
                "birth_year": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "edited": {
                    "type": "string"
                },
                "eye_color": {
                    "type": "string"
                },
                "films": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "gender": {
                    "type": "string"
                },
                "hair_color": {
                    "type": "string"
                },
                "height": {
                    "type": "string"
                },
                "mass": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "skin_color": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "types.CharacterResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/types.Characters"
                }
            }
        },
        "types.Characters": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "next": {},
                "previous": {},
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Character"
                    }
                }
            }
        },
        "types.CharactersResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/types.Characters"
                }
            }
        },
        "types.CommentResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/models.Comment"
                }
            }
        },
        "types.CommentsResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                }
            }
        },
        "types.Movie": {
            "type": "object",
            "properties": {
                "characters": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "comment": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created": {
                    "type": "string"
                },
                "director": {
                    "type": "string"
                },
                "edited": {
                    "type": "string"
                },
                "episode_id": {
                    "type": "integer"
                },
                "opening_crawl": {
                    "type": "string"
                },
                "producer": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "types.MovieResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/types.Movie"
                }
            }
        },
        "types.Movies": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "next": {},
                "previous": {},
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Movie"
                    }
                }
            }
        },
        "types.MoviesResp": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "result": {
                    "$ref": "#/definitions/types.Movies"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             ":{variable} port",
	BasePath:         "/documentation/",
	Schemes:          []string{},
	Title:            "MOVIES API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
