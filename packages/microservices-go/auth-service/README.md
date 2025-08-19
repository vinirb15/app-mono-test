#Auth Service

Authentication service in Go using Gin and PostgreSQL.

## Endpoints

- `GET /healthy`: Checks if the service is running.
- `POST /signup`: Creates a new user.
- `POST /login`: Logs in and returns tokens.
- `POST /refresh`: Refreshes authentication tokens.
- `POST /logout`: Logs out the user.

## How to run

1. Set the required environment variables (see the configuration file in `config`).
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Run the service:
   ```
   go run main.go
   ```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin)
- [pgx](https://github.com/jackc/pgx)
- PostgreSQL

## Request example

```bash
curl -X POST http://localhost:8080/signup -d '{"username":"user","password":"pass"}' -H "Content-Type: application/json"
```
