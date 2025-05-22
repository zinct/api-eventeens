//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"

	"goevents/internal/infrastructure/repositories/event"
	"goevents/internal/interface/delivery/http/controllers"
	"goevents/internal/usecases"
)

func InitializeEventController(db *sql.DB) *controllers.EventController {
	wire.Build(
		event.NewEventRepositoryMySQL,
		usecases.NewEventUsecase,
		controllers.NewEventController,
	)
	return nil
}
