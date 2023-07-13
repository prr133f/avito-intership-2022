package domain

import (
	"avito-intership-2022/models"
	"time"
)

func (l *Logic) MonthlyReport(beginDate, endDate time.Time) ([]models.MonthlyReport, error) {
	return l.Pg.MonthlyReport(beginDate, endDate)
}
