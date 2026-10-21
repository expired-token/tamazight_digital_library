package model

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Language    string    `json:"language"`
	Dialect     string    `json:"dialect"`
	Description string    `json:"description"`
	Content     any       `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
