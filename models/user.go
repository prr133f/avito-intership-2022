package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Balance  float64   `json:"balance"`
	Reserved float64   `json:"reserved"`
}

type UserAddBalance struct {
	ID         uuid.UUID `json:"id"`
	AddBalance float64   `json:"add_balance"`
}
