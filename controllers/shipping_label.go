package controllers

import (
	"fmt"
	"time"

	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllShippingLabel(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateShippingLabel(c *fiber.Ctx) error {
	var r models.Response
	var Shipping models.PostShippingLabel
	err := c.BodyParser(&Shipping)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	etd := (time.Now()).Format("2006-01-02")
	fmt.Println(etd)

	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
