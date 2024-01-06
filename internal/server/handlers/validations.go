package handlers

import (
	components "hotel-checkin/cmd/web/components"
	"hotel-checkin/internal/database"
	"hotel-checkin/internal/models"
	"net/http"

	"github.com/a-h/templ"
)

func ValidateHotelEmail(w http.ResponseWriter, r *http.Request, s database.Service) {
	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
		http.Error(w, "Error parsing form", http.StatusBadRequest)
	}
	email := r.Form.Get("owner_email")
	isHotelExist := s.IsHotelExist(email)
	if isHotelExist {
		component := components.InputElement(validEmailElement(email, "text-red-500", "email already exist"))
		_ = component.Render(r.Context(), w)
		return
	}
	component := components.InputElement(validEmailElement(email, "text-green-500", "valid email"))
	_ = component.Render(r.Context(), w)
}

func validEmailElement(email, color, message string) models.InputProps {
	return models.InputProps{
		Type:        "email",
		Name:        "owner_email",
		IsRequired:  true,
		PlaceHolder: "Enter email",
		Attributes: templ.Attributes{
			"value":   email,
			"hx-post": "/hotel/email/validate",
		},
		MainAttributes: templ.Attributes{
			"hx-target": "this",
			"hx-swap":   "outerHTML",
		},
		StatsMessage:     message,
		StatsMessageColor: color,
	}
}
