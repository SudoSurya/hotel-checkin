package handlers

import (
	"fmt"
	"hotel-checkin/cmd/web/views"
	"hotel-checkin/internal/database"
	"hotel-checkin/internal/models"
	"hotel-checkin/internal/utils"
	"net/http"
)

func CreateHotelForm(w http.ResponseWriter, r *http.Request, s database.Service) {
	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
		http.Error(w, "Error parsing form", http.StatusBadRequest)
	}
	// Do something with form data
	hotelData := models.Hotel{
		Name:       r.Form.Get("hotel_name"),
		City:       r.Form.Get("city"),
		State:      r.Form.Get("state"),
		Country:    r.Form.Get("country"),
		Zip:        utils.Str_int(r.Form.Get("zip")),
		Landline:   utils.Str_int(r.Form.Get("landline")),
		OwnerName:  r.Form.Get("owner_name"),
		OwnerEmail: r.Form.Get("owner_email"),
		Password:   r.Form.Get("password"),
	}
	isHotelExist := s.IsHotelExist(hotelData.OwnerEmail)
	if isHotelExist {
		component := views.Response(false, "Hotel already exist")
		_ = component.Render(r.Context(), w)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = s.InsertHotel(hotelData)
	if err != nil {
		fmt.Println(err)
		component := views.Response(false, "Error inserting hotel")
		_ = component.Render(r.Context(), w)
		return
	}
	component := views.Response(true, "Hotel inserted successfully")
	_ = component.Render(r.Context(), w)
}
