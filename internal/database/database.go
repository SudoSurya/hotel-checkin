package database

import (
	"context"
	"database/sql"
	"fmt"
	"hotel-checkin/internal/database/functions"
	"hotel-checkin/internal/models"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string
	InsertHotel(hotel models.Hotel) error
    IsHotelExist(email string) bool
}

type ConfigDB struct {
	db *sql.DB
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &ConfigDB{db: db}
	return s
}

func (s *ConfigDB) InsertHotel(hotel models.Hotel) error {
	err := functions.CreateHotel(s.db, &hotel)
	if err != nil {
		return fmt.Errorf("error inserting hotel: %w", err)
	}
	return nil
}

func (s *ConfigDB) Health() map[string]string {
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

func (s *ConfigDB) IsHotelExist(email string) bool {
    return functions.IsHotelExist(s.db, email)
}

