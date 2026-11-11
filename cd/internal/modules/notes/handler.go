package notes

import (
	"cd/internal/platform/response"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type NoteService interface {
	CreateNote(ctx context.Context, note *Note) error
	GetAllNotes(ctx context.Context) ([]Note, error)
	GetNoteByID(ctx context.Context, id int64) (*Note, error)
	UpdateNote(ctx context.Context, id int64, note *Note) (*Note, error)
	DeleteNote(ctx context.Context, id int64) error
}

type Handler struct {
	service NoteService
}

func NewHandler(service NoteService) *Handler {
	return &Handler{
		service: service,
	}
}

// CREATE NOTE
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON Request")
		return
	}

	err = h.service.CreateNote(r.Context(), &note)

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, "Note Created Successfully", note)
}

// GET ALL NOTES

func (h *Handler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.service.GetAllNotes(r.Context())

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch notes")
		return
	}
	response.JSON(w, http.StatusOK, "Notes fetched successfully", notes)
}

// GET NOTE BY ID
func (h *Handler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid note ID")
		return
	}
	note, err := h.service.GetNoteByID(r.Context(), id)

	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			response.Error(w, http.StatusNotFound, "Note not found")
			return
		}
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "Notes fetched successfully", note)

}

// UPDATE NOTE
func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	var note Note

	err = json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON request")
		return
	}

	updatedNote, err := h.service.UpdateNote(r.Context(), id, &note)

	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			response.Error(w, http.StatusNotFound, "Note not found")
			return
		}

		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, "Note updated successfully", updatedNote)
}

// DELETE NOTE
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	err = h.service.DeleteNote(r.Context(), id)

	if err != nil {
		if errors.Is(err, ErrNoteNotFound) {
			response.Error(w, http.StatusNotFound, "Note not found")
			return
		}

		response.Error(w, http.StatusInternalServerError, "Failed to delete note")
		return
	}

	response.JSON(w, http.StatusOK, "Note deleted successfully", nil)
}
