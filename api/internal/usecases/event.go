package usecases

import (
	"context"
	"goevents/internal/domain/entities"
	"goevents/internal/domain/repositories"
	"goevents/internal/domain/usecases"
)

type EventUsecase struct {
	repo repositories.EventRepository
}

func NewEventUsecase(repo repositories.EventRepository) usecases.EventUsecase {
	return &EventUsecase{repo: repo}
}

func (u *EventUsecase) Create(ctx context.Context, event *entities.Event) error {
	panic("not implemented")
}

func (u *EventUsecase) FindAll(ctx context.Context) ([]*entities.Event, error) {
	panic("not implemented")
}

func (u *EventUsecase) Find(ctx context.Context, id string) (*entities.Event, error) {
	panic("not implemented")
}

func (u *EventUsecase) Update(ctx context.Context, event *entities.Event) error {
	panic("not implemented")
}

func (u *EventUsecase) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}
