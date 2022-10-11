package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func CreateCartonHistory(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	var sysLog models.SyncLogger
	var frm models.CartonHistory
	if err := c.BodyParser(&frm); err != nil {
		sysLog.Title = "Carton history body parser not allow!"
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
		db.Create(&sysLog)
		// Return History
		r.Message = sysLog.Title
		r.Data = nil
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if frm.SerialNo != "" {
		sysLog.Title = "Carton history body parser not allow!"
		sysLog.Description = "Serail No is null!"
		sysLog.IsSuccess = false
		db.Create(&sysLog)
		// Return History
		r.Message = sysLog.Title
		r.Data = nil
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	go services.CreateCartonHistoryData(&frm)
	go services.CreateCarton(&frm)
	r.Message = "Create Carton History"
	r.Data = nil
	return c.Status(fiber.StatusCreated).JSON(&r)
}
