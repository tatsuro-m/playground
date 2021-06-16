package entity

import "time"

type Post struct {
	ID        uint   `json:"id" gorm:"autoIncrement;primaryKey"`
	Title     string `json:"title" binding:"required" gorm:"not null"`
	Content   string `json:"content" binding:"required" gorm:"not null"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
