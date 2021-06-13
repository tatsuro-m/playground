package entity

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primarykey" gorm:"autoIncrement"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
