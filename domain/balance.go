package domain

import (
	"avito-intership-2022/models"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (pg *Postgres) GetBalance(id uuid.UUID) (models.GetBalance, error) {
	var balance float64

	if err := pg.DB.
		QueryRow(context.Background(), `SELECT balance FROM users_schema.user 
		WHERE id=$1`, id).
		Scan(&balance); err != nil {
		return models.GetBalance{}, nil
	}

	return models.GetBalance{Balance: balance}, nil
}

func (pg *Postgres) AddBalance(dto models.UserAddBalance) (models.AddBalance, error) {
	var newBalance float64

	if err := pg.DB.
		QueryRow(context.Background(), `INSERT INTO users_schema.user (id, balance, reserved) 
		VALUES($1, $2, 0) 
		ON CONFLICT (id) 
		DO UPDATE SET balance = users_schema.user.balance + $2
		RETURNING balance`, dto.ID, dto.AddBalance).
		Scan(&newBalance); err != nil {
		return models.AddBalance{}, err
	}

	return models.AddBalance{NewBalance: newBalance}, nil
}

func (pg *Postgres) ReserveAmount(dto models.UserReserveAmount) (models.ReserveBalance, error) {
	var reservedBalance float64
	var balance float64

	tx, err := pg.DB.Begin(context.Background())
	if err != nil {
		return models.ReserveBalance{}, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), `UPDATE users_schema.user 
	SET balance = balance - $1, reserved = reserved + $1
	WHERE id = $2 
	RETURNING balance, reserved`, dto.ReservedAmount, dto.UserID).Scan(&balance, &reservedBalance); err != nil {
		return models.ReserveBalance{}, err
	}
	if balance < 0 {
		tx.Rollback(context.Background())
		return models.ReserveBalance{}, fmt.Errorf("balance cannot becomes negative")
	}

	if err := tx.QueryRow(context.Background(), `INSERT INTO orders_schema.order(id, date, service_id)
	VALUES($1, DEFAULT, $2)`, dto.OrderID, dto.ServiceID).Scan(); err != nil {
		return models.ReserveBalance{}, err
	}

	if err := tx.QueryRow(context.Background(), `INSERT INTO services_schema.service(id, title, amount)
	VALUES($1, $2, $3)`, dto.ServiceID, dto.Title, dto.ReservedAmount).Scan(); err != nil {
		return models.ReserveBalance{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return models.ReserveBalance{}, err
	}

	return models.ReserveBalance{ReservedBalance: reservedBalance}, nil
}

func (pg *Postgres) CommitAmount(dto models.UserCommitAmount) (models.ReserveBalance, error) {
	var reservedBalance float64

	tx, err := pg.DB.Begin(context.Background())
	if err != nil {
		return models.ReserveBalance{}, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), `UPDATE users_schema.user
	SET reserved = reserved - $1
	WHERE id = $2
	RETURNING reserved`, dto.CommitedAmount, dto.UserID).Scan(&reservedBalance); err != nil {
		return models.ReserveBalance{}, err
	}

	if reservedBalance < 0 {
		tx.Rollback(context.Background())
		return models.ReserveBalance{}, fmt.Errorf("reserved balance cannot becomes negative")
	}

	return models.ReserveBalance{ReservedBalance: reservedBalance}, nil
}

func (pg *Postgres) RestoreReservedAmount(dto models.UserRestoreReservedAmount) (models.GetBalance, error) {
	var currentBalance float64
	var reservedBalance float64

	tx, err := pg.DB.Begin(context.Background())
	if err != nil {
		return models.GetBalance{}, err
	}
	defer tx.Rollback(context.Background())

	if err := tx.QueryRow(context.Background(), `UPDATE users_schema.user 
	SET reserved = reserved - $1, balance = balance + $1
	WHERE id = $2
	RETURNING balance, reserved`, dto.RestoredAmount, dto.UserID).Scan(&currentBalance, &reservedBalance); err != nil {
		return models.GetBalance{}, err
	}

	if reservedBalance < 0 {
		tx.Rollback(context.Background())
		return models.GetBalance{}, fmt.Errorf("reserved balance cannot becomes negative")
	}

	if err := tx.QueryRow(context.Background(), `DELETE FROM orders_schema.order WHERE id = $1`, dto.OrderID).Scan(); err != nil {
		return models.GetBalance{}, err
	}

	if err := tx.QueryRow(context.Background(), `DELETE FROM services_schema.service WHERE id = $1`, dto.ServiceID).Scan(); err != nil {
		return models.GetBalance{}, err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return models.GetBalance{}, err
	}

	return models.GetBalance{Balance: currentBalance}, nil
}
