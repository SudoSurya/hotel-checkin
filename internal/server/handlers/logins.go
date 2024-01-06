package handlers

import (
	components "hotel-checkin/cmd/web/componets"
	"hotel-checkin/cmd/web/views"
	"hotel-checkin/internal/database"
	"hotel-checkin/internal/models"
	"net/http"
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
		Name:  "admin",
		Value: apiKey,
		Path:  "/",
        Secure: true,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Location", "/admin/dashboard")
}

func GetAdminDashboard(w http.ResponseWriter, r *http.Request, db database.Service, admin models.Admin) {
    err := components.AdminDashBoard(admin).Render(r.Context(), w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
