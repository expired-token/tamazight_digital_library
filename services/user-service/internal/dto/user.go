package dto

import "time"

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Country   string `json:"country"`
	// Languages []string       `json:"spoken_languages"`
	// Dialects  []string       `json:"spoken_dialects"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `json:"deleted_at"`
}

type CreateUpdateUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Country   string `json:"country"`
	// Languages []string       `json:"spoken_languages"`
	// Dialects  []string       `json:"spoken_dialects"`
}

/*type UpdateUserRequest struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Country   string `json:"country"`
	// Languages []string       `json:"spoken_languages"`
	// Dialects  []string       `json:"spoken_dialects"`
}*/
