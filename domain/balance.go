package domain

import (
	"avito-intership-2022/models"
	"context"

	"github.com/google/uuid"
)

func (pg *Postgres) GetBalance(id uuid.UUID) (float64, error) {
	var balance float64

	if err := pg.DB.
		QueryRow(context.Background(), "SELECT balance FROM users_schema.user WHERE id=$1", id).
		Scan(&balance); err != nil {
		return 0, nil
	}

	return balance, nil
}

func (pg *Postgres) AddBalance(dto models.UserAddBalance) (float64, error) {
	var newBalance float64

	if err := pg.DB.
		QueryRow(context.Background(), "UPDATE users_schema.user SET balance = balance + $1 WHERE id = $2 RETURNING balance", dto.AddBalance, dto.ID).
		Scan(&newBalance); err != nil {
		return 0, err
	}

	return newBalance, nil
}
