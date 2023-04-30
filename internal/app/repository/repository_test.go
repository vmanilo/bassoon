package repository

import (
	"context"
	"testing"

	"bassoon/internal/app/model"

	"github.com/stretchr/testify/assert"
)

func TestStoreAndRetrieve(t *testing.T) {
	ctx := context.Background()

	repo := New()

	input := &model.Port{ID: "1", Code: "001"}
	assert.NoError(t, repo.CreatePort(ctx, input))
	input.ID = "reset"
	input.Code = "reset"

	port, err := repo.GetPort(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, &model.Port{ID: "1", Code: "001"}, port)
}
