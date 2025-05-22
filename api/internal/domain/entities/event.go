package entities

import "time"

type Event struct {
	ID          string
	Title       string
	Description string
	Date        time.Time
}
