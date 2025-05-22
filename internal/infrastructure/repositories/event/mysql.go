package event

import (
	"context"
	"database/sql"
	"goevents/internal/domain/entities"
	"goevents/internal/domain/repositories"
)

type EventRepositoryMySQL struct {
	db *sql.DB
}

func NewEventRepositoryMySQL(db *sql.DB) repositories.EventRepository {
	return &EventRepositoryMySQL{db: db}
}

func (r *EventRepositoryMySQL) Create(ctx context.Context, event *entities.Event) error {
	panic("not implemented")
}

func (r *EventRepositoryMySQL) FindAll(ctx context.Context) ([]*entities.Event, error) {
	panic("not implemented")
}

func (r *EventRepositoryMySQL) Find(ctx context.Context, id string) (*entities.Event, error) {
	panic("not implemented")
}

func (r *EventRepositoryMySQL) Update(ctx context.Context, event *entities.Event) error {
	panic("not implemented")
}

func (r *EventRepositoryMySQL) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}
