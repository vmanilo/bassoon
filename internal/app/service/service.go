package service

import (
	"context"
	"errors"

	"bassoon/internal/app/model"
)

var ErrInvalidData = errors.New("invalid data")

type Repository interface {
	GetPort(ctx context.Context, portID string) (*model.Port, error)
	IsPortExists(ctx context.Context, portID string) (bool, error)
	CreatePort(ctx context.Context, port *model.Port) error
	UpdatePort(ctx context.Context, port *model.Port) error
}

type service struct {
	store Repository
}

func New(store Repository) *service {
	return &service{store: store}
}

func (s *service) StorePort(ctx context.Context, port *model.Port) error {
	if port == nil || port.ID == "" {
		return ErrInvalidData
	}

	exists, err := s.store.IsPortExists(ctx, port.ID)
	if err != nil {
		return err
	}

	if exists {
		return s.store.UpdatePort(ctx, port)
	}

	return s.store.CreatePort(ctx, port)
}

func (s *service) RetrievePort(ctx context.Context, portID string) (*model.Port, error) {
	return s.store.GetPort(ctx, portID)
}
