//+build wireinject

package di

import (
	"github.com/Tak-Iwamoto/todo/application"
	"github.com/Tak-Iwamoto/todo/domain/repository"
	"github.com/Tak-Iwamoto/todo/handler"
	"github.com/Tak-Iwamoto/todo/infrastructure"
	"github.com/google/wire"
)

func InitListTodoHandler() (*handler.ListTodoHandler, error) {
	wire.Build(handler.NewListTodoHandler, application.NewTodoApp, infrastructure.NewTodoRepository, wire.Bind(new(repository.TodoRepository), new(*infrastructure.TodoRepository)))
	return &handler.ListTodoHandler{}, nil
}

func InitCreateTodoHandler() (*handler.CreateTodoHandler, error) {
	wire.Build(handler.NewCreateTodoHandler, application.NewTodoApp, infrastructure.NewTodoRepository, wire.Bind(new(repository.TodoRepository), new(*infrastructure.TodoRepository)))
	return &handler.CreateTodoHandler{}, nil
}
