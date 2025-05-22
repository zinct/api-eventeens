package event

import (
	"context"
	"database/sql"
	"fmt"
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
	query := `INSERT INTO events (title, description, date) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, event.Title, event.Description, event.Date)
	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/event/mysql - Create - r.db.ExecContext: %w", err)
	}
	return nil
}

func (r *EventRepositoryMySQL) FindAll(ctx context.Context) ([]*entities.Event, error) {
	query := `SELECT * FROM events`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/event/mysql - FindAll - r.db.QueryContext: %w", err)
	}
	defer rows.Close()

	events := []*entities.Event{}
	for rows.Next() {
		var event entities.Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date)
		if err != nil {
			return nil, fmt.Errorf("internal/infrastructure/repositories/event/mysql - FindAll - rows.Scan: %w", err)
		}
		events = append(events, &event)
	}
	return events, nil
}

func (r *EventRepositoryMySQL) Find(ctx context.Context, id string) (*entities.Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	var event entities.Event
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date); err != nil {
		return nil, fmt.Errorf("internal/infrastructure/repositories/event/mysql - Find - r.db.QueryRowContext: %w", err)
	}

	return &event, nil
}

func (r *EventRepositoryMySQL) Update(ctx context.Context, event *entities.Event) error {
	query := `UPDATE events SET title = ?, description = ?, date = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, event.Title, event.Description, event.Date, event.ID)
	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/event/mysql - Update - r.db.ExecContext: %w", err)
	}
	return nil
}

func (r *EventRepositoryMySQL) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("internal/infrastructure/repositories/event/mysql - Delete - r.db.ExecContext: %w", err)
	}
	return nil
}
