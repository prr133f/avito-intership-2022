package view

import (
	"avito-intership-2022/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/rs/zerolog/log"
)

// @Summary		Get balance of user
// @Description	Returns balance of user by id
// @Tags			Balance
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"uuid of user"
// @Success		200	{object}	models.GetBalance
// @Failure		400	{object}	models.Error
// @Router			/balance/{id} [get]
func (v *View) GetBalance(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	balance, err := v.Logic.GetBalance(id)
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(balance)
}

// @Summary		Add balance to user
// @Description	Returns new balance of user
// @Tags			Balance
// @Accept			json
// @Produce		json
// @Param			dto	body		models.UserAddBalance	true	"DTO for adding user balance"
// @Success		200	{object}	models.AddBalance
// @Failure		400	{object}	models.Error
// @Router			/balance [put]
func (v *View) AddBalance(c *fiber.Ctx) error {
	dto := new(models.UserAddBalance)

	if err := c.BodyParser(dto); err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	if dto.AddBalance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{Text: "add_balance cannot be negative"})
	}

	newBalance, err := v.Logic.AddBalance(*dto)
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(newBalance)
}

// @Summary		Reserve part of user balance
// @Description	Reserves a part of user balance
// @Tags			Balance
// @Accept			json
// @Produce		json
// @Param			dto	body		models.UserReserveAmount	true	"DTO for reserving part of user balance"
// @Success		200	{object}	models.ReserveBalance
// @Failure		400	{object}	models.Error
// @Router			/reserve [put]
func (v *View) ReserveAmount(c *fiber.Ctx) error {
	dto := new(models.UserReserveAmount)

	if err := c.BodyParser(dto); err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	reservedBalance, err := v.Logic.ReserveAmount(*dto)
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(reservedBalance)
}

// @Summary		Commit reserved amount
// @Description	Commites previous reserved amount
// @Tags			Balance
// @Accept			json
// @Produce		json
// @Param			dto	body		models.UserCommitAmount	true	"DTO for commiting amount"
// @Success		200	{object}	models.ReserveBalance
// @Failure		400	{object}	models.Error
// @Router			/commit [put]
func (v *View) CommitAmount(c *fiber.Ctx) error {
	dto := new(models.UserCommitAmount)

	if err := c.BodyParser(dto); err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	reservedBalance, err := v.Logic.CommitAmount(*dto)
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(reservedBalance)
}

// @Summary		Restore part of reserved user balance
// @Description	Restores a part of reserved user balance
// @Tags			Balance
// @Accept			json
// @Produce		json
// @Param			dto	body		models.UserRestoreReservedAmount	true	"DTO for restoring part of reserved user balance"
// @Success		200	{object}	models.GetBalance
// @Failure		400	{object}	models.Error
// @Router			/restore [put]
func (v *View) RestoreAmount(c *fiber.Ctx) error {
	dto := new(models.UserRestoreReservedAmount)

	if err := c.BodyParser(dto); err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	currentBalance, err := v.Logic.RestoreReservedAmount(*dto)
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(currentBalance)
}
