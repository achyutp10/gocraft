package notes

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

// CREATE

func (s *Service) CreateNote(ctx context.Context, note *Note) error {
	note.Title = strings.TrimSpace(note.Title)
	note.Content = strings.TrimSpace(note.Content)

	if note.Title == "" {
		return errors.New("Title is required")
	}
	if note.Content == "" {
		return errors.New("Content is required")
	}
	if len(note.Title) > 255 {
		return errors.New("title cannot exceed 255 characters")
	}
	return s.repository.Create(ctx, note)
}

// READ ALL
func (s *Service) GetAllNotes(ctx context.Context) ([]Note, error) {
	return s.repository.GetAll(ctx)
}

// READ BY ID
func (s *Service) GetNoteByID(ctx context.Context, id int64) (*Note, error) {
	if id <= 0 {
		return nil, errors.New("Invalid Note ID")
	}
	return s.repository.GetByID(ctx, id)
}

// UPDATE
func (s *Service) UpdateNote(ctx context.Context, id int64, note *Note) (*Note, error) {
	if id <= 0 {
		return nil, errors.New("Invalid Note ID")
	}
	note.Title = strings.TrimSpace(note.Title)
	note.Content = strings.TrimSpace(note.Content)

	if note.Title == "" {
		return nil, errors.New("Title is required")
	}
	if note.Content == "" {
		return nil, errors.New("Content is required")
	}
	if len(note.Title) > 255 {
		return nil, errors.New("title cannot exceed 255 characters")
	}

	return s.repository.Update(ctx, id, note)

}

// DELETE
func (s *Service) DeleteNote(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("Invalid Note ID")
	}
	return s.repository.Delete(ctx, id)
}
