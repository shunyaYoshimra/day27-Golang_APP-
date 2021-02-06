package entity

import "time"

type Occupation struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Kind        string    `json:"kind"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	UserID      int       `gorm:"required" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
