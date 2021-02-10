package entity

import "time"

type Image struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FileName  string    `json:"file_name"`
	Time      string    `json:"time"`
	PostID    int       `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
