package functions

import (
	"database/sql"
	"errors"
	"fmt"
	"hotel-checkin/internal/models"
)

func IsAdmim(db *sql.DB, email, password string) (bool, error) {
	var adminPass string
	const query = "SELECT password FROM admin WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&adminPass)
	if err != nil {
		return false, errors.New("email or password is incorrect")
	}
	if adminPass != password {
		return false, errors.New("incorrect password")
	}
	return true, nil
}

func Getadminapikey(db *sql.DB, apiKey string) (models.Admin, error) {
	const getUser = `
        SELECT id,email,password,api_key,created_at,updated_at,name,department,isAuthorized
        FROM admin
        WHERE api_key = $1
        `
	var admin models.Admin
	err := db.QueryRow(getUser, apiKey).Scan(
        &admin.ID,
        &admin.Email,
        &admin.Password,
        &admin.ApiKey,
        &admin.CreatedAt,
        &admin.UpdatedAt,
        &admin.Name,
        &admin.Department,
        &admin.IsAuthorized,
    )
    if err != nil {
        return models.Admin{}, fmt.Errorf("error getting admin: %w", err)
    }
	return admin, nil
}

func GetAdminApiKeyByEmail(db *sql.DB, email string) (string, error) {
	const query = "SELECT api_key FROM admin WHERE email = $1"
	var apiKey string
	err := db.QueryRow(query, email).Scan(&apiKey)
	if err != nil {
		return "", err
	}
	return apiKey, nil
}
