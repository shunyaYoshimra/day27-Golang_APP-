package entity

import "time"

type Article struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"required" json:"title"`
	UserID    int       `gorm:"required" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ArticleLine struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Line      string    `gorm:"required,size:255" json:"line"`
	ArticleID int       `gorm:"required" json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BoldNumber struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Number    int       `gorm:"required" json:"number"`
	ArticleID int       `gorm:"required" json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LinkNumber struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Number    int       `gorm:"required" json:"number"`
	ArticleID int       `gorm:"required" json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
