# Google cloud compute with golang

**_This project was created for self education purpose_**
This project communicates with google cloud API.
Create multiple `computes` objects, each one reflects a specifc zone in GCE.
You can view the instances list, turn off or on your instances and more.

## ⚡️ How to run

1. Clone
2. Rename `.env.example` to `.env` and fill it with your environment values.
3. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

   - [golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations

4. Run project by this command:

```bash
make docker.run
```

or run manually by
```bash
 go run main.go
```
5. Under `client_api` you will find examples for creating API calls

## 📦 Used packages

| Name                                                                  | Version    | Type       |
| --------------------------------------------------------------------- | ---------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)                     | `v2.25.0`  | core       |
| [google-cloud-go](https://github.com/googleapis/google-cloud-go)      | `v2.25.0`  | core       |
| [Create Go App CLI](https://github.com/create-go-app/cli)             | `v2.25.0`  | template   |
| [gofiber/jwt](https://github.com/gofiber/jwt)                         | `v2.2.7`   | middleware |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)         | `v2.24.0`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.0`   | tests      |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt)                   | `v4.2.0`   | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.4.0`   | config     |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx)                       | `v1.3.4`   | database   |
| [jackc/pgx](https://github.com/jackc/pgx)                             | `v4.14.1`  | database   |
| [go-redis/redis](https://github.com/go-redis/redis)                   | `v8.11.4`  | cache      |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.7.8`   | utils      |
| [google/uuid](https://github.com/google/uuid)                         | `v1.3.0`   | utils      |
| [go-playground/validator](https://github.com/go-playground/validator) | `v10.10.0` | utils      |

## 🗄 File structure
Created by [Create Go App CLI](https://github.com/create-go-app/cli)
### ./app

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/repository` folder for describe `const` of your project
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/cache` folder with in-memory cache setup functions (by default, Redis)
- `./platform/database` folder with database setup functions (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

