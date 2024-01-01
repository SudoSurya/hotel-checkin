package server

import (
	"hotel-checkin/cmd/web"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/js/*", fileServer)
	r.Get("/", templ.Handler(web.Homepage()).ServeHTTP)
    r.Get("/login", templ.Handler(web.LoginForm()).ServeHTTP)
	return r
}
