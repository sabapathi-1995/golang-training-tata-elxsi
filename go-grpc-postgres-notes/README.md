# Go gRPC + Postgres (Docker) — Unary + Client Streaming + Server Streaming + BiDi

This is a **single Go project** that demonstrates:
- **gRPC server** in Go
- **PostgreSQL** persistence (pgxpool)
- **Dockerfile** (multi-stage) + **docker-compose** (server + postgres)
- gRPC RPC types:
  - Unary
  - Server streaming
  - Client streaming
  - Bidirectional streaming

> Note: To keep this project **fully self-contained without protoc**, it uses a **JSON gRPC codec** (still gRPC transport/streams),
> and includes a `.proto` file as an API contract/reference. In real projects, you would generate stubs via `protoc`.

## Services

`NotesService` stores simple notes in Postgres.

RPCs:
- Unary: `CreateNote`
- Server streaming: `ListNotes`
- Client streaming: `UploadNotes`
- BiDi streaming: `ChatNotes`

## Run with Docker

```bash
docker compose up --build
```

Server listens on: `localhost:50051`

## Run the client locally

In another terminal:

```bash
go run ./cmd/client -addr localhost:50051
```

It will:
1) Create a note (unary)
2) Upload 3 notes (client streaming)
3) List notes (server streaming)
4) Chat notes (bidi) — sends a few notes and receives ACKs

## DB schema

Created automatically on first start via `docker-entrypoint-initdb.d/init.sql`:
- `notes(id uuid, title, body, created_at)`

## Files you care about

- `cmd/server/main.go` — gRPC server + db wiring
- `cmd/client/main.go` — demo client for all RPC types
- `internal/api/notes_grpc.go` — gRPC service + client stub (manual, no protoc)
- `internal/grpcjson/codec.go` — JSON codec for gRPC
- `internal/db/db.go` — Postgres pool + queries
- `docker-compose.yml` — postgres + server
- `Dockerfile` — server container
