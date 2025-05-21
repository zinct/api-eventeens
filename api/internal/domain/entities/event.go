package entities

import "time"

type Event struct {
	ID          string    `binding:"required"`
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Date        time.Time `binding:"required"`
}
