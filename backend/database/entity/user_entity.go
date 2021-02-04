package entity

import "time"

type User struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `gorm:"required" json:"name"`
	Email     string     `gorm:"required" json:"email"`
	Password  string     `gorm:"required" json:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
