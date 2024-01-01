package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"hotel-checkin/internal/models"
	"log"
	"os"
	"time"
)

type Service interface {
	Health() map[string]string
	InsertHotel(hotel models.Hotel) error
}

type service struct {
	db *sql.DB
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
	db, err := sql.Open("libsql", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func randomString() string {
	return fmt.Sprintf("%d", time.Now().Nanosecond())
}

func (s *service) InsertHotel(hotel models.Hotel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO hotels (
        id,
        hotel_name,
        city,
        state,
        country,
        zip,
        landline,
        owner_name,
        owner_email,
        created_at,
        updated_at
    ) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?,?,?)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}

	_, err = stmt.ExecContext(ctx,
        randomString(),
		hotel.Name,
		hotel.City,
		hotel.State,
		hotel.Country,
		hotel.Zip,
		hotel.Landline,
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

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
