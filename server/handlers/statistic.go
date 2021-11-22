package handlers

import (
	"x-men/app/adapter"
	"x-men/app/modules/statistic"

	"github.com/gofiber/fiber/v2"
)

type statisticHandler struct {
	controller statistic.Controller
}

func NewStatisticHandlers(controller statistic.Controller) statisticHandler {
	return statisticHandler{controller: controller}
}

func (s statisticHandler) Stats(c *fiber.Ctx) (err error) {
	fiberAdapter := adapter.NewFiberAdapter(c)
	s.controller.GetStatistic(fiberAdapter)
	return
}
