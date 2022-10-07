package controllers

import (
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func UploadReceiveExcel(c *fiber.Ctx) error {
	var r models.Response
	// r.Message = services.MessageUploadFileCompleted()
	return c.Status(fiber.StatusCreated).JSON(&r)
}
