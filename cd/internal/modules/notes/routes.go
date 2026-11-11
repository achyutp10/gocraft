package notes

import "net/http"

func RegisterRoutes(
	mux *http.ServeMux,
	handler *Handler,
) {

	mux.HandleFunc(
		"POST /api/v1/notes",
		handler.CreateNote,
	)

	mux.HandleFunc(
		"GET /api/v1/notes",
		handler.GetAllNotes,
	)

	mux.HandleFunc(
		"GET /api/v1/notes/{id}",
		handler.GetNoteByID,
	)

	mux.HandleFunc(
		"PUT /api/v1/notes/{id}",
		handler.UpdateNote,
	)

	mux.HandleFunc(
		"DELETE /api/v1/notes/{id}",
		handler.DeleteNote,
	)
}
