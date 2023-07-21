package routes

import (
	"avito-intership-2022/config"
	"avito-intership-2022/database"
	"avito-intership-2022/domain"
	v1 "avito-intership-2022/routes/v1"
	"avito-intership-2022/view"
	"context"

	"github.com/rs/zerolog/log"

	_ "avito-intership-2022/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, conf *config.Config) {
	pg, err := database.NewPG(context.Background(), conf.DSN)
	if err != nil {
		log.Fatal().Err(err)
	}
	view := view.View{Logic: &domain.Logic{Pg: pg}}
	router := v1.Router{App: app, View: &view}
	router.Routes()
	app.Get("/ping", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })
	app.Get("/swagger/*", swagger.HandlerDefault)
}
