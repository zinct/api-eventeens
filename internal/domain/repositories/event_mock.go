package repositories

import (
	"context"
	"goevents/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	mock.Mock
}

func (m *EventRepositoryMock) Create(ctx context.Context, event *entities.Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *EventRepositoryMock) FindAll(ctx context.Context) ([]*entities.Event, error) {
	args := m.Called(ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Event), nil
}

func (m *EventRepositoryMock) Find(ctx context.Context, id string) (*entities.Event, error) {
	args := m.Called(ctx, id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Event), nil
}

func (m *EventRepositoryMock) Update(ctx context.Context, event *entities.Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *EventRepositoryMock) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
