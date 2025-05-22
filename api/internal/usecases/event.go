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
	return u.repo.Create(ctx, event)
}

func (u *EventUsecase) FindAll(ctx context.Context) ([]*entities.Event, error) {
	return u.repo.FindAll(ctx)
}

func (u *EventUsecase) Find(ctx context.Context, id string) (*entities.Event, error) {
	return u.repo.Find(ctx, id)
}

func (u *EventUsecase) Update(ctx context.Context, event *entities.Event) error {
	return u.repo.Update(ctx, event)
}

func (u *EventUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
