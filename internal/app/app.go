package app

import (
	"context"
	"errors"
	"log"
	"net/http"

	"bassoon/config"
	"bassoon/internal/app/model"
	"bassoon/internal/app/repository"
	"bassoon/internal/app/service"
	"bassoon/internal/parser"
	rest "bassoon/internal/transport/http"
)

type Service interface {
	StorePort(ctx context.Context, port *model.Port) error
}

type App struct {
	version string
	cfg     *config.Config
}

func New(version string, cfg *config.Config) *App {
	return &App{
		version: version,
		cfg:     cfg,
	}
}

func (a *App) Run(ctx context.Context) error {
	srv := service.New(repository.New())

	dbTask := Run(ctx, func(c context.Context) <-chan struct{} {
		shutdown := make(chan struct{})

		go func() {
			if err := populateDatabase(c, srv, a.cfg.PortsFilepath); err != nil {
				log.Fatalf("failed to populate database: %v", err)
			}
			close(shutdown)
		}()

		return shutdown
	})

	server := rest.NewServer(srv)

	gracefulShutdown := WithGracefulShutdown(ctx, server, dbTask)

	log.Printf("app.version=%s | starting server on port%s\n", a.version, a.cfg.HTTPPort)

	if err := server.Serve(a.cfg.HTTPPort); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server failed to start: %v", err)
	}

	<-gracefulShutdown

	return nil
}

func populateDatabase(ctx context.Context, srv Service, filepath string) error {
	log.Println("start populating database")

	ch, err := parser.ParseFile(ctx, filepath)
	if err != nil {
		return err
	}

	for port := range ch {
		if err := srv.StorePort(ctx, port); err != nil {
			return err
		}
	}

	log.Println("populate database - DONE")

	return nil
}
