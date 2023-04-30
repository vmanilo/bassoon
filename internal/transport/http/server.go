package http

import (
	"context"
	"log"
	"net/http"

	"bassoon/internal/app/model"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Service interface {
	StorePort(ctx context.Context, port *model.Port) error
	RetrievePort(ctx context.Context, portID string) (*model.Port, error)
}

type server struct {
	service Service
	srv     http.Server
}

func NewServer(service Service) *server {
	srv := &server{
		service: service,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Route("/v1/ports", func(r chi.Router) {
		r.Post("/", srv.createPort)
		r.Get("/{id}", srv.getPort)
	})

	srv.srv.Handler = router

	return srv
}

func (s *server) Serve(port string) error {
	s.srv.Addr = port

	return s.srv.ListenAndServe()
}

func (s *server) GracefulStop() {
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown: %v", err)

		return
	}

	log.Println("HTTP server gracefully shutdown")
}
