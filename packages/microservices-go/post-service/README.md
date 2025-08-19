# Post Service

A RESTful API service for managing posts, likes, and comments, built with Go and Gin.

## Features
- Health check endpoint
- Feed retrieval
- Post creation, update, and detail retrieval
- Like and comment creation
- JWT-based authentication middleware

## Endpoints
- `GET /healthy` — Health check
- `GET /feed` — Get posts feed (auth required)
- `GET /posts/:postID` — Get post details
- `POST /posts` — Create a new post (auth required)
- `PUT /posts/:postID` — Update a post (auth required)
- `POST /likes` — Like a post (auth required)
- `POST /comments` — Comment on a post (auth required)

## Configuration
Configuration is loaded via the `config` package. Set environment variables as needed (see `config` package for details).

## Running
```sh
cd post-service
go run main.go
```

## Dependencies
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) (via `database` package)

## Project Structure
- `main.go` — Entry point
- `config/` — Configuration loading
- `database/` — Database connection
- `handlers/` — HTTP handlers
- `middleware/` — Authentication middleware

## License
MIT
