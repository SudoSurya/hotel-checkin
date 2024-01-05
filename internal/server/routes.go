package server

import (
	"fmt"
	"hotel-checkin/cmd/web"
	"hotel-checkin/cmd/web/views"
	"hotel-checkin/internal/auth"
	"hotel-checkin/internal/models"
	"hotel-checkin/internal/server/handlers"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type adminAuthedHandler func(http.ResponseWriter, *http.Request, models.Admin)

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
	r.With(auth.RedirectIfLoggedInMiddleware).Get("/admin/login/form", templ.Handler(views.AdminLoginForm()).ServeHTTP)
	r.Post("/admin/login", s.AdminLogin)
	r.With(auth.AdminAuthMiddleware).Get("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Admin Dashboard"))
	})
	return r
}

func (s *Server) createHotel(w http.ResponseWriter, r *http.Request) {
	handlers.CreateHotelForm(w, r, s.db)
}
func (s *Server) validateHotelEmail(w http.ResponseWriter, r *http.Request) {
	handlers.ValidateHotelEmail(w, r, s.db)
}

func (s *Server) AdminLogin(w http.ResponseWriter, r *http.Request) {
	handlers.AdminLogin(w, r, s.db)
}

func (s *Server) middlewareAuth(handler adminAuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.Getadminapikey(r)
		if err != nil {
			w.Write([]byte("Unauthorized"))
			return
		}
		admin, err := s.db.Getadminapikey(apiKey)
		if err != nil {
			w.Write([]byte("Unauthorized"))
			return
		}

		handler(w, r, admin)
	}
}
