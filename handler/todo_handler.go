package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tak-Iwamoto/todo/application"
	"github.com/Tak-Iwamoto/todo/domain/entity"
	"github.com/google/uuid"
)

var _ http.Handler = (*ListTodoHandler)(nil)
var _ http.Handler = (*CreateTodoHandler)(nil)

// ListTodoHandler is get all todos
type ListTodoHandler struct {
	app *application.TodoApp
}

// NewListTodoHandler initializes ListTodoHandler
func NewListTodoHandler(a *application.TodoApp) *ListTodoHandler {
	return &ListTodoHandler{
		app: a,
	}
}

func (h *ListTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todos, err := h.app.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

type CreateTodoHandler struct {
	app *application.TodoApp
}

func NewCreateTodoHandler(a *application.TodoApp) *CreateTodoHandler {
	return &CreateTodoHandler{
		app: a,
	}
}

func (h *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t *entity.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	t.PublicID = uuid.New().String()
	if err := h.app.CreateTodo(r.Context(), t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
