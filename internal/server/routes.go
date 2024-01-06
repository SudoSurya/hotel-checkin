package server

import (
	"fmt"
	"hotel-checkin/cmd/web"
	"hotel-checkin/cmd/web/views"
	"hotel-checkin/internal/auth"
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
	// hotel routes
	hotelRouter := chi.NewRouter()
	hotelRouter.Get("/register", templ.Handler(views.HotelRegisterForm()).ServeHTTP)
	hotelRouter.Get("/login", templ.Handler(views.HotelLoginForm()).ServeHTTP)
	hotelRouter.Post("/register/create", s.createHotel)
	hotelRouter.Post("/email/validate", s.validateHotelEmail)
	r.Mount("/hotel", hotelRouter)
	// admin routes
	adminRouter := chi.NewRouter()
	adminRouter.With(auth.RedirectIfLoggedInMiddleware).Get("/login/form", templ.Handler(views.AdminLoginForm()).ServeHTTP)
	adminRouter.Post("/login", s.AdminLogin)
	adminRouter.With(auth.AdminAuthMiddleware).Get("/dashboard", s.adminMiddlewareAuth(s.getAdminDashboard))
	adminRouter.Get("/home", s.adminMiddlewareAuth(s.getAdminDashboard))
    adminRouter.Post("/logout", s.AdminLogout)
	r.Mount("/admin", adminRouter)
	// utils routes
	r.Get("/delayed-redirect", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second) // Simulate processing or delay
		w.Header().Set("HX-Redirect", "/hotel/login")
		fmt.Fprintf(w, "Redirecting to login...")
	})
	return r
}

