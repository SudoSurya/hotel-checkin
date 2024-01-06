package handlers

import (
	"hotel-checkin/cmd/web/views"
	"hotel-checkin/internal/database"
	"hotel-checkin/internal/models"
	"net/http"
	"time"
)

func AdminLogin(w http.ResponseWriter, r *http.Request, db database.Service) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	if email == "" || password == "" {
		errElement := views.ErrorAdminLogin("email and password are required")
		_ = errElement.Render(r.Context(), w)
		http.Error(w, "email and password are required", http.StatusBadRequest)
		return
	}
	isAuthorized, err := db.IsAdmin(email, password)
	if err != nil {
		errElement := views.ErrorAdminLogin(err.Error())
		_ = errElement.Render(r.Context(), w)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !isAuthorized {
		errElement := views.ErrorAdminLogin("email or password is incorrect")
		_ = errElement.Render(r.Context(), w)
		http.Error(w, "email or password is incorrect", http.StatusBadRequest)
		return
	}
	apiKey, err := db.GetAdminApiKeyByEmail(email)
	if err != nil {
		errElement := views.ErrorAdminLogin(err.Error())
		_ = errElement.Render(r.Context(), w)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{
		Name:   "admin",
		Value:  apiKey,
		Path:   "/",
		Secure: true,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Location", "/admin/dashboard")
}

func AdminLogout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "admin",
		Value:   "",
		Path:    "/",
		Secure:  true,
		MaxAge:  -1,                               // Set MaxAge to a negative value to expire the cookie immediately
		Expires: time.Now().Add(-1 * time.Minute), // Set Expires to a time in the past
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Location", "/admin/login/form")
}

func GetAdminDashboard(w http.ResponseWriter, r *http.Request, db database.Service, admin models.Admin) {
	err := views.AdminDashBoard(admin).Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
