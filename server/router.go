package server

import (
	"net/http"

	"github.com/Tak-Iwamoto/todo/di"
)

// NewRouter returns router
func NewRouter() (*http.ServeMux, error) {
	mux := http.NewServeMux()
	listHandler, err := di.InitListTodoHandler()
	if err != nil {
		return nil, err
	}
	createHandler, err := di.InitCreateTodoHandler()
	if err != nil {
		return nil, err
	}
	mux.Handle("/list", listHandler)
	mux.Handle("/create", createHandler)
	return mux, nil
}
