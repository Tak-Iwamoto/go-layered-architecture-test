package infrastructure

import (
	"context"

	"gorm.io/gorm"

	"github.com/Tak-Iwamoto/todo/domain/entity"
	"github.com/Tak-Iwamoto/todo/infrastructure/table"
)

// TodoRepository is repository for todo
type TodoRepository struct {
	db *gorm.DB
}

// NewTodoRepository initializes TodoRepository
func NewTodoRepository() (*TodoRepository, error) {
	db, err := NewGormDB()
	if err != nil {
		return nil, err
	}
	return &TodoRepository{
		db: db,
	}, nil
}

// Create todo
func (r *TodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	t := &table.Todo{
		UUID:    todo.Id,
		Title:   todo.Title,
		Content: todo.Content,
	}
	return r.db.Create(t).Error
}

// GetAll todos
func (r *TodoRepository) GetAll(ctx context.Context) ([]entity.Todo, error) {
	var todos []table.Todo
	r.db.Find(&todos)
	var result []entity.Todo
	for _, t := range todos {
		result = append(result, *t.ToEntity())
	}
	return result, nil
}

// GetTodoByID get todo by ID
func (r *TodoRepository) GetTodoByID(ctx context.Context, id int) (*entity.Todo, error) {
	var todo *table.Todo
	r.db.First(&todo, "uuid = ?", id)
	return todo.ToEntity(), nil
}
