package notes

import (
	"errors"
	"time"
)

var ErrNoteNotFound = errors.New("note not found")

type Note struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
