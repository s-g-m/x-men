package server

import (
	"x-mers/config"
	"x-mers/util/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func configFibre() fiber.Config {
	config := fiber.Config{
		ServerHeader: config.Name(),
	}
	return config
}

func configRecover() recover.Config {
	config := recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(e interface{}) {
			logs.Error("Panic error fiber recover", e)
		},
	}
	return config
}
