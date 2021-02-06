package entity

import "time"

type Profile struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Description string    `json:"description"`
	Subject     string    `gorm:"require" json:"subject"`
	UserID      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
