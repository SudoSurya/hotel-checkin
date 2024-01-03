package server

import (
	"fmt"
	"hotel-checkin/cmd/web"
	"hotel-checkin/cmd/web/views"
	"hotel-checkin/internal/server/handlers"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/js/*", fileServer)
	r.Get("/", templ.Handler(views.Homepage()).ServeHTTP)
	r.Get("/hotel/register", templ.Handler(views.HotelRegisterForm()).ServeHTTP)
	r.Get("/hotel/login", templ.Handler(views.HotelLoginForm()).ServeHTTP)
	r.Post("/hotel/register/create", s.createHotel)
	r.Get("/delayed-redirect", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second) // Simulate processing or delay
		w.Header().Set("HX-Redirect", "/hotel/login")
		fmt.Fprintf(w, "Redirecting to login...")
	})
	r.Post("/hotel/email/validate", s.validateHotelEmail)
	return r
}

func (s *Server) createHotel(w http.ResponseWriter, r *http.Request) {
	handlers.CreateHotelForm(w, r, s.db)
}
func (s *Server) validateHotelEmail(w http.ResponseWriter, r *http.Request) {
    handlers.ValidateHotelEmail(w, r, s.db)
}
