package domain

import (
	"avito-intership-2022/models"
	"context"
	"time"
)

func (pg *Postgres) MonthlyReport(beginDate, endDate time.Time) ([]models.MonthlyReport, error) {
	var report []models.MonthlyReport
	rows, err := pg.DB.Query(context.Background(), `SELECT service.title, service.amount 
	FROM services_schema.service AS service
	JOIN orders_schema.order AS order ON order.service_id = service.id
	WHERE order.date BETWEEN $1 AND $2
	GROUP BY service.title`, beginDate, endDate)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r models.MonthlyReport
		err := rows.Scan(&r.ServiceName, &r.Margin)
		if err != nil {
			return nil, err
		}
		report = append(report, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return report, nil
}
