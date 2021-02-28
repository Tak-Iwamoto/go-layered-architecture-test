package table

import "time"

// Model is gorm model
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
