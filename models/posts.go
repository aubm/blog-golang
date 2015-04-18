package models

import (
	"time"
)

type Post struct {
	ID        int64     `gorm:"primary_key"json:"id"`
	Title     string    `sql:"size:255"json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `sql:"type:text"json:"content"`
}
