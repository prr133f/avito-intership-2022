package view

import (
	"avito-intership-2022/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// @Summary		Get monthly report
// @Description	Returns monthly margin report
// @Tags			Report
// @Accept			json
// @Produce		json
// @Param			date	path		string	true	"date in format YYYY-MM"
// @Success		200		{object}	[]models.MonthlyReport
// @Failure		400		{object}	models.Error
// @Router			/report/{date} [put]
func (v *View) MonthlyReport(c *fiber.Ctx) error {
	stringDate := c.Params("date")
	beginDate, err := time.Parse("2006-01", stringDate)
	if err != nil {
		return err
	}
	endDate := beginDate.AddDate(0, 1, 0)

	report, err := v.Logic.MonthlyReport(beginDate, endDate)
	if err != nil {
		log.Error().Err(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.Error{
			Text: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(report)
}
