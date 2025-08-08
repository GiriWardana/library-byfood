// In models/book.go
package models

import (
	"time"
)

type Book struct {
	Id        uint       `json:"id" gorm:"primaryKey"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Year      uint       `json:"year"`
}
