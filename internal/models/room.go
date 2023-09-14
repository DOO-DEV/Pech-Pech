package models

import "time"

type Room struct {
	Name        string
	Description string
	CreatedBy   string
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
