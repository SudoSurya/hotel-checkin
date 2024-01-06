package server

import (
	"fmt"
	"hotel-checkin/internal/auth"
	"hotel-checkin/internal/models"
	"hotel-checkin/internal/server/handlers"
	"net/http"
)

type adminAuthedHandler func(http.ResponseWriter, *http.Request, models.Admin)

func (s *Server) createHotel(w http.ResponseWriter, r *http.Request) {
	handlers.CreateHotelForm(w, r, s.db)
}
func (s *Server) validateHotelEmail(w http.ResponseWriter, r *http.Request) {
	handlers.ValidateHotelEmail(w, r, s.db)
}

func (s *Server) AdminLogin(w http.ResponseWriter, r *http.Request) {
	handlers.AdminLogin(w, r, s.db)
}

func (s *Server) AdminLogout(w http.ResponseWriter, r *http.Request) {
	handlers.AdminLogout(w, r)
}
func (s *Server) getAdminDashboard(w http.ResponseWriter, r *http.Request, admin models.Admin) {
	handlers.GetAdminDashboard(w, r, s.db, admin)
}

func (s *Server) adminMiddlewareAuth(handler adminAuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.Getadminapikey(r)
		if err != nil {
			fmt.Println("auth api error", err)
			http.Error(w, "Unauthorized", http.StatusBadRequest)
			return
		}
		admin, err := s.db.GetadminByapikey(apiKey)
		if err != nil {
			fmt.Println("error getting admin", err)
			http.Error(w, "Unauthorized", http.StatusBadRequest)
			return
		}

		handler(w, r, admin)
	}
}
