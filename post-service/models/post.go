package models

import "time"

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	AuthorId uint `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}