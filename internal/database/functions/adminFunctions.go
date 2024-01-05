package functions

import (
	"database/sql"
	"errors"
	"fmt"
	"hotel-checkin/internal/models"
)

func IsAdmim(db *sql.DB, email, password string) (bool, error) {
	var adminPass string
	fmt.Println(email, password)
	const query = "SELECT password FROM admin WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&adminPass)
    if err != nil {
        return false, errors.New("email is incorrect")
    }
    if adminPass != password {
        return false, errors.New("password is incorrect")
    }
    return true, nil
}

func Getadminapikey(db *sql.DB, apiKey string) (models.Admin, error) {
		const getUser = `
        SELECT id, created_at, updated_at, name, api_key
        FROM admin
        WHERE api_key = $1
        `
        var admin models.Admin
        err := db.QueryRow(getUser, apiKey).Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt, &admin.Name, &admin.ApiKey)
        if err != nil {
            return models.Admin{}, err
        }
        return admin, nil
}

