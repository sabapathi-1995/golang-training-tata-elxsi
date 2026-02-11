# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS build
WORKDIR /src
RUN apk add --no-cache git ca-certificates protobuf

# Install protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

COPY go.mod ./
RUN go mod download

COPY . .

# Generate protobuf code into ./gen (protoc-based approach)
RUN mkdir -p gen && \
    protoc -I . \
      --go_out=./gen --go_opt=paths=source_relative \
      --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
      api/notes/v1/notes.proto

RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /out/grpc-server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=build /out/grpc-server /grpc-server
EXPOSE 50051
USER nonroot:nonroot
ENTRYPOINT ["/grpc-server"]
