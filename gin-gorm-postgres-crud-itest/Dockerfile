FROM golang:1.22-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git ca-certificates && update-ca-certificates
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server ./cmd/server

FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=builder /app/server /server
ENV ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/server"]
