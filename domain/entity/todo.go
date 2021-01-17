package entity

import "time"

// Todo is todo entity
type Todo struct {
	ID        string    `json:"-"`
	PublicID  string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
}
