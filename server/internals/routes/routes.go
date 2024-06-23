package routes

import (
	"net/http"

	"github.com/Alfred-Onuada/go-dropbox/internals/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /files", handlers.GetFiles)

	mux.HandleFunc("POST /files", handlers.AddFile)

	// TODO:
	mux.HandleFunc("PATCH /files/{fileName}", handlers.UpdateFile)

	// TODO:
	mux.HandleFunc("DELETE /files/{fileName}", handlers.DeleteFile)
}
