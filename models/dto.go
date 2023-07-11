package models

import "github.com/google/uuid"

type UserAddBalance struct {
	ID         uuid.UUID `json:"id"`
	AddBalance float64   `json:"add_balance"`
}

type UserReserveAmount struct {
	ID             uuid.UUID `json:"id"`
	ReservedAmount float64   `json:"reserved_amount"`
}
