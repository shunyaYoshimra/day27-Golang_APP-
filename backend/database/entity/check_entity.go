package entity

import "time"

type Check struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	ArticleID int       `gorm:"required" json:"article_id"`
	UserID    int       `gorm:"required" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upadted_at"`
}
