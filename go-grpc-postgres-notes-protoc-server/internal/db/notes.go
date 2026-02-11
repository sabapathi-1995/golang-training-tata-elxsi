package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type NoteRow struct {
	ID        uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
}

func (s *Store) InsertNote(ctx context.Context, title, body string) (NoteRow, error) {
	id := uuid.New()
	now := time.Now().UTC()

	_, err := s.Pool.Exec(ctx,
		`INSERT INTO notes(id, title, body, created_at) VALUES($1,$2,$3,$4)`,
		id, title, body, now,
	)
	if err != nil {
		return NoteRow{}, err
	}
	return NoteRow{ID: id, Title: title, Body: body, CreatedAt: now}, nil
}

func (s *Store) ListNotes(ctx context.Context, limit, offset int) ([]NoteRow, error) {
	rows, err := s.Pool.Query(ctx,
		`SELECT id, title, body, created_at FROM notes ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]NoteRow, 0, limit)
	for rows.Next() {
		var r NoteRow
		if err := rows.Scan(&r.ID, &r.Title, &r.Body, &r.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}
