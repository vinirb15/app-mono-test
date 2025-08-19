# Profile Service

Profile Service is a RESTful API written in Go using the Gin framework. It manages user profiles and follower relationships.

## API Endpoints
- `GET /healthy` — Health check
- `GET /me` — Get authenticated user profile (requires authentication)
- `GET /user/:email` — Get user by email
- `GET /user/id/:id` — Get user by ID
- `POST /followers` — Follow a user (requires authentication)
- `GET /followers/:userID` — List followers of a user (requires authentication)

## Features
- Health check endpoint
- Retrieve authenticated user profile
- Get user by email or ID
- Follow users and list followers

## Requirements
- Go 1.18 or newer
- Database supported by the `database.Connection` implementation
- Proper configuration via environment variables or config files

## Running

```bash
go run main.go
```

## Configuration
The service loads configuration using `config.LoadConfig()`. Ensure all required environment variables or configuration files are set (see the `config` package for details).

## License
MIT
