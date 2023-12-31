package main

import (
	"avito-intership-2022/config"
	"avito-intership-2022/routes"
	"avito-intership-2022/utils"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//	@title			Avito Intership 2022 Solution API
//	@version		1.0
//	@description	This is solution iv avito backend intership 2022 by prr133f

// @host		localhost:8080
// @BasePath	/v1
func main() {
	app := fiber.New(fiber.Config{
		AppName:        "prr133f-avito-intership-2022",
		RequestMethods: []string{"GET", "HEAD", "PUT"},
	})
	app.Use(logger.New(logger.ConfigDefault))
	utils.InitLogger()

	config := config.ParseConfig()

	routes.InitRouter(app, config)

	log.Fatal().Err(app.Listen(fmt.Sprintf(":%s", config.Port)))
}
