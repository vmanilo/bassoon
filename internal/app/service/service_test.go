package service

import (
	"context"
	"errors"
	"testing"

	"bassoon/internal/app/model"
	"bassoon/internal/app/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	repository *mocks.MockRepository
}

func (s *ServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repository = mocks.NewMockRepository(s.ctrl)
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestPublisherService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestStorePort() {
	ctx := context.Background()

	cases := []struct {
		name            string
		req             *model.Port
		expectFuncCalls func(s *ServiceSuite)
		err             error
	}{
		{
			name:            "invalid request",
			req:             &model.Port{},
			expectFuncCalls: func(s *ServiceSuite) {},
			err:             errors.New("invalid data"),
		},
		{
			name: "valid request - create",
			req:  &model.Port{ID: "test", Code: "001"},
			expectFuncCalls: func(s *ServiceSuite) {
				s.repository.EXPECT().
					IsPortExists(ctx, "test").
					Return(false, nil).
					Times(1)

				s.repository.EXPECT().
					CreatePort(ctx, &model.Port{ID: "test", Code: "001"}).
					Return(nil).
					Times(1)
			},
			err: nil,
		},
		{
			name: "valid request - update",
			req:  &model.Port{ID: "test", Code: "001"},
			expectFuncCalls: func(s *ServiceSuite) {
				s.repository.EXPECT().
					IsPortExists(ctx, "test").
					Return(true, nil).
					Times(1)

				s.repository.EXPECT().
					UpdatePort(ctx, &model.Port{ID: "test", Code: "001"}).
					Return(nil).
					Times(1)
			},
			err: nil,
		},
	}

	for _, c := range cases {
		s.Run(c.name, func() {
			c.expectFuncCalls(s)

			srv := New(s.repository)
			err := srv.StorePort(ctx, c.req)
			assertError(s.T(), err, c.err)
		})
	}
}

func assertError(t *testing.T, err, expected error) {
	t.Helper()

	if expected != nil {
		assert.EqualError(t, err, expected.Error())
	} else {
		assert.NoError(t, err)
	}
}
