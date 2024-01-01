package models

type Hotel struct {
	Name       string `json:"hotel_name"`
	City       string `json:"city"`
	Country    string `json:"country"`
	State      string `json:"state"`
	Zip        string `json:"zip"`
	Landline   string `json:"landline"`
	OwnerName  string `json:"owner_name"`
	OwnerEmail string `json:"owner_email"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
}
