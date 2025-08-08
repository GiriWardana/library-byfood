package models

type CleanURL struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Path   string `json:"path" gorm:"unique;not null"` // e.g., "/my-book"
	Target string `json:"target" gorm:"not null"`      // e.g., "/books/123"
}
