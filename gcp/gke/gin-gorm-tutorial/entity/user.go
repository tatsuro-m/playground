package entity

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"first_name" gorm:"not null" binding:"required"`
	LastName  string `json:"last_name" gorm:"not null" binding:"required"`
	Posts     []Post
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
