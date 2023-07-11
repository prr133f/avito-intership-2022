package models

import "github.com/google/uuid"

type UserAddBalance struct {
	ID         uuid.UUID `json:"id"`
	AddBalance float64   `json:"add_balance"`
}

type UserReserveAmount struct {
	OrderID        uuid.UUID `json:"order_id"`
	UserID         uuid.UUID `json:"user_id"`
	ServiceID      uuid.UUID `json:"service_id"`
	ReservedAmount float64   `json:"reserved_amount"`
	Title          string    `json:"title"`
}

type UserCommitAmount struct {
	UserID         uuid.UUID `json:"user_id"`
	ServiceID      uuid.UUID `json:"service_id"`
	OrderID        uuid.UUID `json:"order_id"`
	CommitedAmount float64   `json:"commited_amount"`
}

type UserRestoreReservedAmount struct {
	OrderID        uuid.UUID `json:"order_id"`
	UserID         uuid.UUID `json:"user_id"`
	ServiceID      uuid.UUID `json:"service_id"`
	RestoredAmount float64   `json:"restored_amount"`
}
