package server

import (
	"hotel-checkin/cmd/web"
	"hotel-checkin/internal/server/handlers"
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
	r.Get("/hotel/register", templ.Handler(web.HotelRegisterForm()).ServeHTTP)
	r.Post("/hotel/register/create", s.createHotel)
	return r
}

func (s *Server) createHotel(w http.ResponseWriter, r *http.Request) {
	handlers.CreateHotelForm(w, r, s.db)
}
