package handlers

import (
	"fmt"
	"hotel-checkin/internal/database"
	"hotel-checkin/internal/models"
	"net/http"
)

func CreateHotelForm(w http.ResponseWriter, r *http.Request) {
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
		Zip:        r.Form.Get("zip"),
		Landline:   r.Form.Get("landline"),
		OwnerName:  r.Form.Get("owner_name"),
		OwnerEmail: r.Form.Get("owner_email"),
	}
	fmt.Println(hotelData)
    fmt.Println(hotelData.Zip,"zip")
    err = database.New().InsertHotel(hotelData)
    if err != nil {
        fmt.Println(err)
        http.Error(w, "Error inserting hotel", http.StatusInternalServerError)
    }
}
