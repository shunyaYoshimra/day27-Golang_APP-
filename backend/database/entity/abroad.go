package entity

import "time"

type Abroad struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Country     string    `json:"country"`
	College     string    `json:"college"`
	Description string    `json:"description"`
	UserID      int       `gorm:"required" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
