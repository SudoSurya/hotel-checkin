package functions

import (
	"context"
	"database/sql"
	"fmt"
	"hotel-checkin/internal/models"
	"hotel-checkin/internal/utils"
	"time"
)

func CreateHotel(db *sql.DB, hotel *models.Hotel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, `INSERT INTO hotels (
        id,
        hotel_name,
        city,
        state,
        country,
        zip,
        landline,
        password,
        owner_name,
        owner_email,
        created_at,
        updated_at
    ) VALUES (?,?,?, ?, ?, ?, ?, ?, ?, ?,?,?)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	_, err = stmt.ExecContext(ctx,
		utils.RandomString(),
		hotel.Name,
		hotel.City,
		hotel.State,
		hotel.Country,
		hotel.Zip,
		hotel.Landline,
		hotel.Password,
		hotel.OwnerName,
		hotel.OwnerEmail,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	if err != nil {
		return fmt.Errorf("error inserting hotel: %w", err)
	}
	fmt.Println("Hotel inserted successfully")
	return nil
}

func GetHotels(db *sql.DB) ([]models.Hotel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT * FROM hotels where status = 'progress'")
	if err != nil {
		return nil, fmt.Errorf("error querying hotels: %w", err)
	}
	defer rows.Close()

	var hotels []models.Hotel
	for rows.Next() {
		var hotel models.Hotel
		err := rows.Scan(
			&hotel.ID,
			&hotel.Name,
			&hotel.City,
			&hotel.State,
			&hotel.Country,
			&hotel.Zip,
			&hotel.Landline,
			&hotel.Password,
			&hotel.OwnerName,
			&hotel.OwnerEmail,
			&hotel.CreatedAt,
			&hotel.UpdatedAt,
			&hotel.ApiKey,
			&hotel.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning hotel: %w", err)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func IsHotelExist(db *sql.DB, email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var hotelID string
	err := db.QueryRowContext(ctx, "SELECT id FROM hotels WHERE owner_email = ?", email).Scan(&hotelID)
	return err == nil
}
