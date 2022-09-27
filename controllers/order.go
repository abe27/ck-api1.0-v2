package controllers

import (
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GenerateOrder(c *fiber.Ctx) error {
	var r models.Response
	services.CreateOrder()
	return c.Status(fiber.StatusCreated).JSON(&r)
}
