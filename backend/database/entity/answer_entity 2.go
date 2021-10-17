package entity

import "time"

type Answer struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Content    string    `gorm:"required" json:"content"`
	UserID     int       `json:"user_id"`
	QuestionID int       `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
