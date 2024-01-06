package models

type Hotel struct {
	ID         string `json:"id"`
	Name       string `json:"hotel_name"`
	City       string `json:"city"`
	Country    string `json:"country"`
	State      string `json:"state"`
	Zip        int    `json:"zip"`
	Landline   int    `json:"landline"`
	OwnerName  string `json:"owner_name"`
	OwnerEmail string `json:"owner_email"`
	Status     string `json:"status"`
	ApiKey     string `json:"api_key"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type Admin struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Department   string `json:"department"`
	IsAuthorized bool   `json:"isAuthorized"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ApiKey       string `json:"api_key"`
}
