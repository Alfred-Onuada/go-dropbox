package routes

import (
	"net/http"

	"github.com/Alfred-Onuada/go-dropbox/internals/auth"
	"github.com/Alfred-Onuada/go-dropbox/internals/handlers"
	"github.com/Alfred-Onuada/go-dropbox/internals/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /files", handlers.GetFiles)

	mux.HandleFunc("POST /files", handlers.AddFile)

	// TODO:
	mux.HandleFunc("PATCH /files/{fileName}", handlers.UpdateFile)

	// TODO:
	mux.HandleFunc("DELETE /files/{fileName}", handlers.DeleteFile)

	mux.HandleFunc("POST /auth/login", auth.LoginHandler)

	mux.HandleFunc("POST /auth/register", auth.RegisterHandler)

	mux.Handle("GET /auth/ping" , middleware.JWTMiddleware(http.HandlerFunc(auth.TestAuth)))
}
