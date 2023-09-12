package models

import "time"

type Room struct {
	ID          string
	Description string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   User
}
