package application

import (
	"context"

	"github.com/Tak-Iwamoto/todo/domain/entity"
	"github.com/Tak-Iwamoto/todo/domain/repository"
)

// TodoApp is todo application
type TodoApp struct {
	repo repository.TodoRepository
}

// NewTodoApp initializes TodoApp
func NewTodoApp(repo repository.TodoRepository) *TodoApp {
	return &TodoApp{
		repo: repo,
	}
}

// GetAll gets all todos
func (t *TodoApp) GetAll(ctx context.Context) ([]entity.Todo, error) {
	return t.repo.GetAll(ctx)
}

// GetTodoByID gets todo by id
func (t *TodoApp) GetTodoByID(ctx context.Context, id int) (*entity.Todo, error) {
	return t.repo.GetTodoByID(ctx, id)
}

// CreateTodo creates todo
func (t *TodoApp) CreateTodo(ctx context.Context, todo *entity.Todo) error {
	return t.repo.Create(ctx, todo)
}
