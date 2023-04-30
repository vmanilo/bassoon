package repository

import (
	"context"
	"errors"

	"bassoon/internal/app/model"
)

var ErrPortNotFound = errors.New("port not found")

type repository struct {
	db map[string]*model.Port
}

func New() *repository {
	return &repository{
		db: make(map[string]*model.Port),
	}
}

func (r *repository) IsPortExists(_ context.Context, portID string) (bool, error) {
	_, exists := r.db[portID]

	return exists, nil
}

func (r *repository) CreatePort(_ context.Context, port *model.Port) error {
	clone := *port
	r.db[port.ID] = &clone

	return nil
}

func (r *repository) UpdatePort(_ context.Context, port *model.Port) error {
	clone := *port
	r.db[port.ID] = &clone

	return nil
}

func (r *repository) GetPort(_ context.Context, portID string) (*model.Port, error) {
	port, exists := r.db[portID]
	if !exists {
		return nil, ErrPortNotFound
	}

	clone := *port

	return &clone, nil
}
