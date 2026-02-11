package api

import "time"

type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Note struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type CreateNoteResponse struct {
	Note *Note `json:"note"`
}

type ListNotesRequest struct{}

type UploadNotesResponse struct {
	Created int32 `json:"created"`
}

type Ack struct {
	Message string `json:"message"`
}

// helper for consistent timestamp formatting
func RFC3339(t time.Time) string { return t.UTC().Format(time.RFC3339) }
