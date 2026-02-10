# Gin + GORM + Postgres CRUD (with Integration Tests)

This is a **simple CRUD service** for `User` using:
- Gin (HTTP)
- GORM (ORM)
- Postgres (DB)
- Integration tests using **Testcontainers** (real Postgres container per test)

## Endpoints
- POST   `/v1/users`
- GET    `/v1/users/:id`
- GET    `/v1/users`
- PUT    `/v1/users/:id`
- DELETE `/v1/users/:id`
- GET    `/health`

## Run locally with docker-compose Postgres
Start Postgres:
```bash
docker compose up -d
```

Run server:
```bash
export DB_DSN='host=localhost user=postgres password=postgres dbname=cruddb port=5432 sslmode=disable TimeZone=Asia/Kolkata'
go run ./cmd/server
```

Test quickly:
```bash
curl -s http://localhost:8080/health
curl -s -XPOST http://localhost:8080/v1/users -H 'content-type: application/json' -d '{"name":"JP","email":"jp@example.com"}'
```

## Integration tests (recommended)
These tests start a **real Postgres** via Testcontainers.

```bash
go test ./... -v
```

The main integration test is:
- `integration/TestUsersCRUD_Integration`

## Build Docker image (Mac M3 / arm64)
```bash
docker build -t gin-crud:arm64 .
```

Run (with your own Postgres DSN):
```bash
docker run --rm -p 8080:8080 \
  -e DB_DSN="host=host.docker.internal user=postgres password=postgres dbname=cruddb port=5432 sslmode=disable TimeZone=Asia/Kolkata" \
  gin-crud:arm64
```
