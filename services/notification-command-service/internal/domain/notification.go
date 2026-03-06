package domain

import "time"

type Notification struct {
	ID string
	UserID string
	Type string
	Message string
	Status string
	CreateAt time.Time
}