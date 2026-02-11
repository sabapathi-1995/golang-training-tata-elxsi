.PHONY: proto build run test docker

proto:
	@echo "Generating protobuf/grpc code into ./gen ..."
	protoc -I . \
	  --go_out=./gen --go_opt=paths=source_relative \
	  --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
	  api/notes/v1/notes.proto

build:
	go build ./...

run:
	go run ./cmd/server

test:
	go test ./...

docker:
	docker build -t grpc-notes:local .
