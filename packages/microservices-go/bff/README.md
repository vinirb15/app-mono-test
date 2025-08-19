# BFF API

API Gateway / Backend For Frontend (BFF) for proxying requests to Auth, Profile, and Post services.

## Features
- User authentication (login, logout, signup, token refresh)
- Profile management (get profile, followers, follow user)
- Post management (feed, create, update, like, comment)
- Health check endpoint
- Swagger UI for API documentation

## Endpoints

### Auth Service
- `POST /login` — Login user
- `POST /logout` — Logout user
- `POST /signup` — Register new user
- `POST /refresh` — Refresh authentication token

### Profile Service
- `GET /me` — Get current user profile (auth required)
- `GET /followers/:user_id` — Get followers of a user (auth required)
- `POST /followers` — Follow a user (auth required)

### Posts Service
- `GET /feed` — Get user feed (auth required)
- `POST /posts` — Create a new post (auth required)
- `GET /posts/:post_id` — Get a post by ID (auth required)
- `PUT /posts/:post_id` — Update a post (auth required)
- `POST /likes` — Like a post (auth required)
- `POST /comments` — Comment on a post (auth required)

### Health Check
- `GET /healthy` — Returns `{ "message": "healthy" }`

### Swagger UI
- `GET /swagger/index.html` — API documentation

## Running

1. Configure environment variables as needed (see `config` package).
2. Run:
   ```sh
   go run main.go
   ```

## Dependencies
- [Gin](https://github.com/gin-gonic/gin)
- [Swaggo](https://github.com/swaggo/gin-swagger)

## Security
- Uses Bearer token authentication for protected endpoints.

## License
MIT
