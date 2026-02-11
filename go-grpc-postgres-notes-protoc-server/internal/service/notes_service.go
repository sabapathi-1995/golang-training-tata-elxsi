package service

import (
	"context"
	"io"
	"log"
	"math/rand/v2"
	"time"

	notesv1 "github.com/jitenpalaparthi/go-grpc-postgres-notes-protoc-server/gen/api/notes/v1"
	"github.com/jitenpalaparthi/go-grpc-postgres-notes-protoc-server/internal/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotesService struct {
	notesv1.UnimplementedNotesServiceServer
	Store *db.Store
}

func NewNotesService(store *db.Store) *NotesService {
	return &NotesService{Store: store}
}

// Unary
func (s *NotesService) CreateNote(ctx context.Context, req *notesv1.CreateNoteRequest) (*notesv1.CreateNoteResponse, error) {
	if req.GetTitle() == "" || req.GetBody() == "" {
		return nil, status.Error(codes.InvalidArgument, "title and body are required")
	}
	row, err := s.Store.InsertNote(ctx, req.GetTitle(), req.GetBody())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "db insert: %v", err)
	}

	return &notesv1.CreateNoteResponse{
		Note: &notesv1.Note{
			Id:        row.ID.String(),
			Title:     row.Title,
			Body:      row.Body,
			CreatedAt: row.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
		},
	}, nil
}

// Server streaming
func (s *NotesService) ListNotes(_ *notesv1.ListNotesRequest, stream notesv1.NotesService_ListNotesServer) error {
	ctx := stream.Context()
	notes, err := s.Store.ListNotes(ctx, 100, 0) // stream latest 100
	if err != nil {
		return status.Errorf(codes.Internal, "db list: %v", err)
	}
	for _, n := range notes {
		if err := stream.Send(&notesv1.Note{
			Id:        n.ID.String(),
			Title:     n.Title,
			Body:      n.Body,
			CreatedAt: n.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
		}); err != nil {
			return err
		}
	}
	return nil
}

// Client streaming
func (s *NotesService) UploadNotes(stream notesv1.NotesService_UploadNotesServer) error {
	ctx := stream.Context()
	var created int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&notesv1.UploadNotesResponse{Created: created})
		}
		if err != nil {
			return err
		}
		if req.GetTitle() == "" || req.GetBody() == "" {
			return status.Error(codes.InvalidArgument, "title and body are required")
		}
		if _, err := s.Store.InsertNote(ctx, req.GetTitle(), req.GetBody()); err != nil {
			return status.Errorf(codes.Internal, "db insert: %v", err)
		}
		created++
	}
}

// BiDi streaming
func (s *NotesService) ChatNotes(stream notesv1.NotesService_ChatNotesServer) error {
	ctx := stream.Context()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if req.GetTitle() == "" || req.GetBody() == "" {
			return status.Error(codes.InvalidArgument, "title and body are required")
		}

		row, err := s.Store.InsertNote(ctx, req.GetTitle(), req.GetBody())
		if err != nil {
			return status.Errorf(codes.Internal, "db insert: %v", err)
		}

		msg := "stored note id=" + row.ID.String()
		log.Printf("bidi: %s", msg)

		if err := stream.Send(&notesv1.Ack{Message: "Hey I have received message"}); err != nil {
			return err
		}
	}
}

// GenerateData(*EmptyRequest, grpc.ServerStreamingServer[GenerateResponse]) error

func (s *NotesService) GenerateData(_ *notesv1.EmptyRequest, stream grpc.ServerStreamingServer[notesv1.GenerateResponse]) error {

	for {
		num := rand.IntN(100000)
		stream.Send(&notesv1.GenerateResponse{Gen: int64(num)})
		time.Sleep(time.Second * 1)
	}

}
