package entity

import "time"

type Contact struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Media     string    `json:"media"`
	Account   string    `json:"account"`
	UserID    int       `gorm:"required" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
