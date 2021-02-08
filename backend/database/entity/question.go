package entity

import "time"

type Question struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"required" json:"title"`
	Content   string    `gorm:"required" json:"content"`
	About     string    `gorm:"required" json:"about"`
	Status    bool      `gorm:"required" json:"status"`
	UserID    int       `gorm:"required" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
