package handlers

import (
	"x-men/app/adapter"
	"x-men/app/modules/dna"

	"github.com/gofiber/fiber/v2"
)

type dnaHandler struct {
	controller dna.Controller
}

func NewDnaHandlers(controller dna.Controller) dnaHandler {
	return dnaHandler{controller: controller}
}

func (s dnaHandler) Mutant(c *fiber.Ctx) (err error) {
	fiberAdapter := adapter.NewFiberAdapter(c)
	s.controller.IsMutant(fiberAdapter)
	return
}
