package model

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique;notnull"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Country   string `json:"country"` // TODO - data from catalog service
	// Languages []string       `json:"spoken_languages"` // TODO - data from catalog service
	// Dialects  []string       `json:"spoken_dialects"`  // TODO - data from catalog service
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
