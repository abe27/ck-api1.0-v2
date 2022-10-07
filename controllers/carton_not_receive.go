package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllCartonNotReceive(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.CartonNotReceive
	err := configs.Store.Limit(100).Where("is_sync=?", false).Find(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageShowAllData("Carton not receive")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}
