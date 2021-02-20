package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tak-Iwamoto/todo/handler"
)

// Run server
func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	router, err := NewRouter()
	if err != nil {
		return 1
	}
	s := handler.NewServer(8080, router)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		s.Stop(ctx)
		return 0
	case <-errCh:
		return 1
	}
}
