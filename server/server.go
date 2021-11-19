package server

import (
	"x-mers/config"
	"x-mers/server/routes"
	"x-mers/util/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer() *fiber.App {
	logs.Info("starting")

	app := fiber.New(configFibre())
	app.Get("/health", routes.Health)
	app.Use(recover.New(configRecover()))
	app.Use(validateAccess)
	app.Use(logger.New())
	routes.Load(app)

	app.Listen(config.Port())
	return nil
}
