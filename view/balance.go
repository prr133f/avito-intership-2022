package view

import (
	"avito-intership-2022/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

//	@Summary		Get balance of user
//	@Description	Returns balance of user by id
//	@Tags			Balance
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"uuid of user"
//	@Success		200	{object}	map[string]float64{}
//	@Failure		400	{object}	map[string]string{}
//	@Router			/balance/{id} [get]
func (v *View) GetBalance(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	balance, err := v.Pg.GetBalance(id)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]float64{
		"balance": balance,
	})
}

//	@Summary		Add balance to user
//	@Description	Returns new balance of user
//	@Tags			Balance
//	@Accept			json
//	@Produce		json
//	@Param			dto	body		models.UserAddBalance	true	"DTO for adding user balance"
//	@Success		200	{object}	map[string]float64{}
//	@Failure		400	{object}	map[string]string{}
//	@Router			/balance [put]
func (v *View) AddBalance(c *fiber.Ctx) error {
	dto := new(models.UserAddBalance)

	if err := c.BodyParser(dto); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	newBalance, err := v.Pg.AddBalance(*dto)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]float64{
		"new_balance": newBalance,
	})
}
