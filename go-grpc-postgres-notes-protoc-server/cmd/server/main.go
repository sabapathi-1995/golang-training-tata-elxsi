package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	notesv1 "github.com/jitenpalaparthi/go-grpc-postgres-notes-protoc-server/gen/api/notes/v1"
	"github.com/jitenpalaparthi/go-grpc-postgres-notes-protoc-server/internal/db"
	"github.com/jitenpalaparthi/go-grpc-postgres-notes-protoc-server/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := getenv("GRPC_ADDR", ":50052")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	store, err := db.NewStore(ctx)
	if err != nil {
		log.Fatalf("db init: %v", err)
	}
	defer store.Close()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
			Timeout:           20 * time.Second,
			Time:              2 * time.Minute,
		}),
	)

	notesv1.RegisterNotesServiceServer(s, service.NewNotesService(store))
	reflection.Register(s)

	go func() {
		log.Printf("gRPC server listening on %s", addr)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Printf("shutting down...")
	s.GracefulStop()
	log.Printf("bye")
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
