package database

import (
	"avito-intership-2022/models"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (pg *Postgres) GetBalance(id uuid.UUID) (models.Balance, error) {
	var balance float64

	if err := pg.DB.
		QueryRow(context.Background(), `SELECT balance FROM users_schema.user 
		WHERE id=$1`, id).
		Scan(&balance); err != nil {
		return models.Balance{}, err
	}

	return models.Balance{Balance: balance}, nil
}

func (pg *Postgres) AddBalance(dto models.UserAddBalance) (models.Balance, error) {
	var newBalance float64

	if err := pg.DB.
		QueryRow(context.Background(), `INSERT INTO users_schema.user (id, balance, reserved) 
		VALUES($1, $2, 0) 
		ON CONFLICT (id) 
		DO UPDATE SET balance = users_schema.user.balance + $2
		RETURNING balance`, dto.ID, dto.AddBalance).
		Scan(&newBalance); err != nil {
		return models.Balance{}, err
	}

	return models.Balance{Balance: newBalance}, nil
}

func (pg *Postgres) ReserveBalance(dto models.UserReserveAmount) (models.ReservedBalance, error) {
	var reservedBalance float64
	var balance float64

	tx, err := pg.DB.Begin(context.Background())
	if err != nil {
		return models.ReservedBalance{}, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), `UPDATE users_schema.user 
	SET balance = balance - $1, reserved = reserved + $1
	WHERE id = $2 
	RETURNING balance, reserved`, dto.ReservedAmount, dto.UserID).Scan(&balance, &reservedBalance); err != nil {
		return models.ReservedBalance{}, err
	}
	if balance < 0 {
		tx.Rollback(context.Background())
		return models.ReservedBalance{}, fmt.Errorf("balance cannot becomes negative")
	}

	if err := tx.QueryRow(context.Background(), `INSERT INTO orders_schema.order(id, date, service_id)
	VALUES($1, DEFAULT, $2)`, dto.OrderID, dto.ServiceID).Scan(); err != nil {
		return models.ReservedBalance{}, err
	}

	if err := tx.QueryRow(context.Background(), `INSERT INTO services_schema.service(id, title, amount)
	VALUES($1, $2, $3)`, dto.ServiceID, dto.Title, dto.ReservedAmount).Scan(); err != nil {
		return models.ReservedBalance{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return models.ReservedBalance{}, err
	}

	return models.ReservedBalance{ReservedBalance: reservedBalance}, nil
}

func (pg *Postgres) CommitAmount(dto models.UserCommitAmount) (models.ReservedBalance, error) {
	var reservedBalance float64

	tx, err := pg.DB.Begin(context.Background())
	if err != nil {
		return models.ReservedBalance{}, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), `UPDATE users_schema.user
	SET reserved = reserved - $1
	WHERE id = $2
	RETURNING reserved`, dto.CommitedAmount, dto.UserID).Scan(&reservedBalance); err != nil {
		return models.ReservedBalance{}, err
	}

	if reservedBalance < 0 {
		tx.Rollback(context.Background())
		return models.ReservedBalance{}, fmt.Errorf("reserved balance cannot becomes negative")
	}

	return models.ReservedBalance{ReservedBalance: reservedBalance}, nil
}

func (pg *Postgres) RestoreReservedAmount(dto models.UserRestoreReservedAmount) (models.Balance, error) {
	var currentBalance float64
	var reservedBalance float64

	tx, err := pg.DB.Begin(context.Background())
	if err != nil {
		return models.Balance{}, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), `UPDATE users_schema.user 
	SET reserved = reserved - $1, balance = balance + $1
	WHERE id = $2
	RETURNING balance, reserved`, dto.RestoredAmount, dto.UserID).Scan(&currentBalance, &reservedBalance); err != nil {
		return models.Balance{}, err
	}

	if reservedBalance < 0 {
		tx.Rollback(context.Background())
		return models.Balance{}, fmt.Errorf("reserved balance cannot becomes negative")
	}

	if err := tx.QueryRow(context.Background(), `DELETE FROM orders_schema.order WHERE id = $1`, dto.OrderID).Scan(); err != nil {
		return models.Balance{}, err
	}

	if err := tx.QueryRow(context.Background(), `DELETE FROM services_schema.service WHERE id = $1`, dto.ServiceID).Scan(); err != nil {
		return models.Balance{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return models.Balance{}, err
	}

	return models.Balance{Balance: currentBalance}, nil
}
