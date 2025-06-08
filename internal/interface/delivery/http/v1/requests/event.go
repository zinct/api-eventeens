package requests

import "time"

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

type UpdateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required`
}
