# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed if you have any go modules from VCS)
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static-ish binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o user-service .

# ---------- Runtime stage ----------
FROM gcr.io/distroless/static-debian12

WORKDIR /

COPY --from=builder /app/user-service /user-service

# Your app listens on APP_PORT (default 8083 via flag only if env empty)
ENV APP_PORT=8083

EXPOSE 8083

ENTRYPOINT ["/user-service"]