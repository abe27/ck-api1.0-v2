package controllers

import (
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func HandlerHello(c *fiber.Ctx) error {
	var r models.Response
	r.Message = services.HelloWorld
	return c.Status(fiber.StatusOK).JSON(&r)
}
