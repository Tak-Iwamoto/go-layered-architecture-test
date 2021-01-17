package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// Server is http server struct
type Server struct {
	server *http.Server
}

// NewServer initializes http server
func NewServer(port int, mux *http.ServeMux) *Server {
	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

// Start http server
func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}
	return nil
}

// Stop http server
func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop: %w", err)
	}
	return nil
}
