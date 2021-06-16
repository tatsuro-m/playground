package entity

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"autoIncrement;primaryKey"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
