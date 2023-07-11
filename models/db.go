package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Balance  float64   `json:"balance"`
	Reserved float64   `json:"reserved"`
}

type Service struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Cost float64   `json:"cost"`
}

type Report struct {
	Date      time.Time `json:"date"`
	ServiceID uuid.UUID `json:"service_id"`
	UserID    uuid.UUID `json:"user_id"`
}
