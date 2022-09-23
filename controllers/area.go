package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllArea(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Area
	// Fetch All Data
	db := configs.Store
	err := db.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("Area")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("Area")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}
