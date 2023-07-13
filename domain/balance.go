package domain

import (
	"avito-intership-2022/models"

	"github.com/google/uuid"
)

func (l *Logic) GetBalance(id uuid.UUID) (models.Balance, error) {
	return l.Pg.GetBalance(id)
}

func (l *Logic) AddBalance(dto models.UserAddBalance) (models.Balance, error) {
	return l.Pg.AddBalance(dto)
}

func (l *Logic) ReserveAmount(dto models.UserReserveAmount) (models.ReservedBalance, error) {
	return l.Pg.ReserveBalance(dto)
}

func (l *Logic) CommitAmount(dto models.UserCommitAmount) (models.ReservedBalance, error) {
	return l.Pg.CommitAmount(dto)
}

func (l *Logic) RestoreReservedAmount(dto models.UserRestoreReservedAmount) (models.Balance, error) {
	return l.Pg.RestoreReservedAmount(dto)
}
