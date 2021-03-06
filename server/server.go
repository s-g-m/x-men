package server

import (
	"x-men/config"
	"x-men/server/handlers"
	"x-men/server/routes"
	"x-men/server/subscriber"
	"x-men/util/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunServer() *fiber.App {
	logs.Info("starting")

	app := fiber.New(configFibre())
	app.Get("/health", handlers.Health)
	app.Use(recover.New(configRecover()))
	app.Use(validateAccess)
	app.Use(logger.New())
	routes.Load(app)

	subscriber.Start()

	app.Listen(config.Port())
	return nil
}
