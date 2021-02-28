package table

import (
	"github.com/Tak-Iwamoto/todo/domain/entity"
)

// Todo is todo table schema
type Todo struct {
	Model
	UUID    string `json:"uuid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

// ConvertToEntity is convert to domain entity
func (t *Todo) ToEntity() *entity.Todo {
	return &entity.Todo{
		Id:        t.UUID,
		Title:     t.Title,
		Content:   t.Content,
		Done:      t.Done,
		CreatedAt: *t.CreatedAt,
	}
}
