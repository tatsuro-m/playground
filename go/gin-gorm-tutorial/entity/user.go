package entity

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primarykey" gorm:"autoIncrement"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
