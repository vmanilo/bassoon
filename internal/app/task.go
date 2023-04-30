package app

import (
	"context"
	"log"
)

type task struct {
	cancel   context.CancelFunc
	shutdown <-chan struct{}
}

func Run(cxt context.Context, f func(context.Context) <-chan struct{}) *task {
	nCtx, cancel := context.WithCancel(cxt)

	return &task{
		cancel:   cancel,
		shutdown: f(nCtx),
	}
}

func (t *task) GracefulStop() {
	log.Println("task: graceful shutdown...")
	t.cancel()
	<-t.shutdown
	log.Println("task: graceful shutdown... DONE")
}
