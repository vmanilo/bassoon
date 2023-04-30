package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const signalsChanSize = 2

type WithGracefulStop interface {
	GracefulStop()
}

func WithGracefulShutdown(ctx context.Context, resources ...WithGracefulStop) <-chan struct{} {
	signals := make(chan os.Signal, signalsChanSize)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	shutdownCompleted := make(chan struct{})

	go func() {
		select {
		case <-ctx.Done():
			log.Printf("catch ctx.Done() event: %v", ctx.Err())
		case sig := <-signals:
			log.Printf("catch %v signal", sig)
		}

		log.Println("ALL: graceful shutdown...")

		for _, res := range resources {
			res.GracefulStop()
		}

		time.Sleep(time.Second)

		log.Println("ALL: graceful shutdown... DONE")
		close(shutdownCompleted)
	}()

	return shutdownCompleted
}
