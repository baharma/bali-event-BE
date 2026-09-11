# Event Bali Backend

Backend service for the Event Bali application. The project is written in Go and currently provides the base configuration, PostgreSQL connection, database migration command, and reusable JSON response helpers.

> The project is still under development. The API entry point in `cmd/api/main.go` does not start an HTTP server yet.

## Tech Stack

- **Go 1.26.1** - programming language
- **Gin** - HTTP web framework and JSON responses
- **PostgreSQL** - relational database
- **GORM** - ORM and database migration
- **godotenv** - loads environment variables from `.env`

## Main Libraries

| Library | Purpose |
| --- | --- |
| [`gin-gonic/gin`](https://github.com/gin-gonic/gin) | HTTP routing and API responses |
| [`gin-contrib/cors`](https://github.com/gin-contrib/cors) | Cross-Origin Resource Sharing middleware |
| [`gorm.io/gorm`](https://gorm.io/) | Object-relational mapping and migrations |
| [`gorm.io/driver/postgres`](https://github.com/go-gorm/postgres) | PostgreSQL driver for GORM |
| [`joho/godotenv`](https://github.com/joho/godotenv) | Loads local environment configuration |
| [`golang-jwt/jwt`](https://github.com/golang-jwt/jwt) | JSON Web Token support |
| [`mongo-driver`](https://github.com/mongodb/mongo-go-driver) | MongoDB driver, available but not currently used |

Other entries in `go.mod` are mostly indirect dependencies required by these libraries.

## Project Structure

```text
event-bali-be/
├── cmd/
│   ├── api/
│   │   └── main.go          # API application entry point
│   └── migrate/
│       └── main.go          # Database migration command
├── internal/
│   ├── config/
│   │   └── config.go        # Environment and application configuration
│   └── database/
│       └── postgres.go      # PostgreSQL and GORM connection setup
├── pkg/
│   └── response/
│       └── response.go      # Standard success and error JSON responses
├── .example.env             # Example environment variables
├── go.mod                   # Go module and dependency definitions
├── go.sum                   # Dependency checksums
└── README.md                # Project documentation
```

## Prerequisites

- Go 1.26.1 or compatible version
- PostgreSQL

## Getting Started

1. Clone the repository and enter the project directory.

2. Create your local environment file:

   ```bash
   cp .example.env .env
   ```

3. Update `.env` with your PostgreSQL credentials:

   ```env
   APP_PORT=8080
   APP_ENV=development
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=secret
   DB_NAME=myapp
   ```

4. Download the dependencies:

   ```bash
   go mod download
   ```

5. Run the database migration:

   ```bash
   go run ./cmd/migrate
   ```

6. Run the API application:

   ```bash
   go run ./cmd/api
   ```

The API command currently exits immediately because the HTTP server and routes have not been implemented.

## Environment Variables

| Variable | Default | Description |
| --- | --- | --- |
| `APP_PORT` | `8080` | HTTP server port |
| `APP_ENV` | `development` | Application environment |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | PostgreSQL username |
| `DB_PASSWORD` | `secret` | PostgreSQL password |
| `DB_NAME` | `myapp` | PostgreSQL database name |

## Database Connection Pool

The PostgreSQL connection is configured with:

- Maximum idle connections: `10`
- Maximum open connections: `100`
- Maximum connection lifetime: `30 minutes`

## API Response Format

Successful response:

```json
{
  "status": "success",
  "message": "Request completed successfully",
  "data": {}
}
```

Error response:

```json
{
  "status": "error",
  "message": "Something went wrong"
}
```
