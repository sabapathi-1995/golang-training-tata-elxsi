# Go gRPC + Postgres (protoc-based) — Server Only

Single server project that demonstrates:
- gRPC Unary + Server streaming + Client streaming + BiDi streaming
- Postgres persistence (pgxpool)
- Dockerfile + docker-compose
- **protoc generation approach**
  - Local: `make proto` (requires protoc + plugins)
  - Docker: generation happens automatically during image build

## Run with Docker

```bash
docker compose up --build
```

Server listens on: `localhost:50051`

## Test without writing a client (grpcurl)

List services:
```bash
grpcurl -plaintext localhost:50051 list
```

Describe service:
```bash
grpcurl -plaintext localhost:50051 describe notes.v1.NotesService
```

Unary CreateNote:
```bash
grpcurl -plaintext -d '{"title":"hello","body":"from grpcurl"}' localhost:50051 notes.v1.NotesService/CreateNote
```

Server streaming ListNotes:
```bash
grpcurl -plaintext -d '{}' localhost:50051 notes.v1.NotesService/ListNotes
```

Client streaming UploadNotes:
```bash
printf '{"title":"n1","body":"b1"}\n{"title":"n2","body":"b2"}\n' | grpcurl -plaintext -d @ localhost:50051 notes.v1.NotesService/UploadNotes
```

BiDi ChatNotes:
```bash
printf '{"title":"c1","body":"b1"}\n{"title":"c2","body":"b2"}\n' | grpcurl -plaintext -d @ localhost:50051 notes.v1.NotesService/ChatNotes
```

## Protobuf generation locally

Install:
- protoc
- protoc-gen-go
- protoc-gen-go-grpc

Then:
```bash
make proto
go run ./cmd/server
```

Generated files go into `./gen/` (and are created during Docker build too).

## Structure
- `api/notes/v1/notes.proto` — API contract
- `cmd/server` — gRPC server entrypoint
- `internal/service` — gRPC implementation (uses generated stubs)
- `internal/db` — Postgres access
- `scripts/init.sql` — DB schema

### Docker 
docker run -d \
  --name pg \
  -p 5432:5432 \
  --network demo-network \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=notesdb \
  -v $(pwd)/scripts/:/docker-entrypoint-initdb.d \
  postgres:16