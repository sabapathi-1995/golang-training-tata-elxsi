package service

import (
	"context"
	"io"
	"log"

	"go-grpc-postgres-notes/internal/api"
	"go-grpc-postgres-notes/internal/db"
)

type NotesService struct {
	api.UnimplementedNotesServiceServer
	Store *db.Store
}

func NewNotesService(store *db.Store) *NotesService {
	return &NotesService{Store: store}
}

// Unary
func (s *NotesService) CreateNote(ctx context.Context, req *api.CreateNoteRequest) (*api.CreateNoteResponse, error) {
	row, err := s.Store.InsertNote(ctx, req.Title, req.Body)
	if err != nil {
		return nil, err
	}
	return &api.CreateNoteResponse{
		Note: &api.Note{
			ID:        row.ID.String(),
			Title:     row.Title,
			Body:      row.Body,
			CreatedAt: api.RFC3339(row.CreatedAt),
		},
	}, nil
}

// Server streaming
func (s *NotesService) ListNotes(_ *api.ListNotesRequest, stream api.NotesService_ListNotesServer) error {
	ctx := stream.Context()
	notes, err := s.Store.ListNotes(ctx, 100, 0) // demo: stream latest 100
	if err != nil {
		return err
	}
	for _, n := range notes {
		if err := stream.Send(&api.Note{
			ID:        n.ID.String(),
			Title:     n.Title,
			Body:      n.Body,
			CreatedAt: api.RFC3339(n.CreatedAt),
		}); err != nil {
			return err
		}
	}
	return nil
}

// Client streaming
func (s *NotesService) UploadNotes(stream api.NotesService_UploadNotesServer) error {
	ctx := stream.Context()
	var created int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&api.UploadNotesResponse{Created: created})
		}
		if err != nil {
			return err
		}
		if _, err := s.Store.InsertNote(ctx, req.Title, req.Body); err != nil {
			return err
		}
		created++
	}
}

// BiDi streaming
func (s *NotesService) ChatNotes(stream api.NotesService_ChatNotesServer) error {
	ctx := stream.Context()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		row, err := s.Store.InsertNote(ctx, req.Title, req.Body)
		if err != nil {
			return err
		}

		msg := "stored note id=" + row.ID.String()
		log.Printf("bidi: %s", msg)

		if err := stream.Send(&api.Ack{Message: msg}); err != nil {
			return err
		}
	}
}
