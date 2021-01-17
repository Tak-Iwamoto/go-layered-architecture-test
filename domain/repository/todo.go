package repository

import (
	"context"
	"github.com/Tak-Iwamoto/todo/domain/entity"
)

// TodoRepository interface
type TodoRepository interface {
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, id int) (*entity.Todo, error)
	Create(ctx context.Context, todo *entity.Todo) error
}
