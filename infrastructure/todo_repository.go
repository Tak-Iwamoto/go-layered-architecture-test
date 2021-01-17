package infrastructure

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Tak-Iwamoto/todo/domain/entity"
)

// TodoRepository is repository for todo
type TodoRepository struct {
	db *sql.DB
}

// NewTodoRepository initializes TodoRepository
func NewTodoRepository() (*TodoRepository, error) {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &TodoRepository{
		db: db,
	}, nil
}

// Create todo
func (r *TodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	ins, err := r.db.Prepare("insert into todos (public_id, title, content, done, created_at) values (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = ins.Exec(todo.PublicID, todo.Title, todo.Content, false, time.Now())
	return err
}

// GetAll todos
func (r *TodoRepository) GetAll(ctx context.Context) ([]entity.Todo, error) {
	rows, err := r.db.Query("select public_id, title, content, done from todos")
	if err != nil {
		return nil, err
	}
	var todos []entity.Todo
	for rows.Next() {
		todo := entity.Todo{}
		if err := rows.Scan(&todo.PublicID, &todo.Title, &todo.Content, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// GetTodoByID get todo by ID
func (r *TodoRepository) GetTodoByID(ctx context.Context, id int) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.QueryRow("select * from todos where ?", id).Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Done)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// Close is close db connection
func (r *TodoRepository) Close() error {
	return r.db.Close()
}
