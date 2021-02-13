package entity

import "time"

type Favorite struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	PostID    int       `gorm:"required" json:"post_id"`
	UserID    int       `gorm:"required" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upadted_at"`
}
