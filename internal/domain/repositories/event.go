package repositories

import (
	"context"
	"goevents/internal/domain/entities"
)

type EventRepository interface {
	Create(ctx context.Context, event *entities.Event) error
	FindAll(ctx context.Context) ([]*entities.Event, error)
	Find(ctx context.Context, id string) (*entities.Event, error)
	Update(ctx context.Context, event *entities.Event) error
	Delete(ctx context.Context, id string) error
}
