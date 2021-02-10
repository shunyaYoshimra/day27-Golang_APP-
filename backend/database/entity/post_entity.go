package entity

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"required" json:"content"`
	Tags      string    `json:"tags"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
