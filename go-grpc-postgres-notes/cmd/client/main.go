package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"go-grpc-postgres-notes/internal/api"
	"go-grpc-postgres-notes/internal/grpcjson"

	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "gRPC address")
	flag.Parse()

	grpcjson.Register()

	conn, err := grpc.Dial(
		*addr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(grpcjson.Codec{})),
	)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	c := api.NewNotesServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("1) Unary CreateNote")
	u, err := c.CreateNote(ctx, &api.CreateNoteRequest{Title: "hello", Body: "from unary"})
	must(err)
	fmt.Printf("created: %+v\n\n", u.Note)

	fmt.Println("2) Client streaming UploadNotes (3 notes)")
	up, err := c.UploadNotes(ctx)
	must(err)
	for i := 1; i <= 3; i++ {
		must(up.Send(&api.CreateNoteRequest{
			Title: fmt.Sprintf("stream note %d", i),
			Body:  "sent via client streaming",
		}))
	}
	upResp, err := up.CloseAndRecv()
	must(err)
	fmt.Printf("uploaded created=%d\n\n", upResp.Created)

	fmt.Println("3) Server streaming ListNotes (latest 100)")
	ls, err := c.ListNotes(ctx, &api.ListNotesRequest{})
	must(err)
	for {
		n, err := ls.Recv()
		if err == io.EOF {
			break
		}
		must(err)
		fmt.Printf("- %s | %s\n", n.ID, n.Title)
	}
	fmt.Println()

	fmt.Println("4) BiDi streaming ChatNotes (send 3, receive ACKs)")
	chat, err := c.ChatNotes(ctx)
	must(err)

	// receive in goroutine
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			ack, err := chat.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("recv err: %v", err)
				return
			}
			fmt.Printf("ACK: %s\n", ack.Message)
		}
	}()

	for i := 1; i <= 3; i++ {
		must(chat.Send(&api.CreateNoteRequest{
			Title: fmt.Sprintf("bidi note %d", i),
			Body:  "sent via bidi streaming",
		}))
	}
	must(chat.CloseSend())
	<-done
	fmt.Println("\nDone.")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
