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
	err := configs.Store.Limit(100).Order("transfer_out_no,lot_no,serial_no").Where("is_sync=?", false).Find(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageShowAllData("Carton not receive")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateCartonNotReceiveByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.CartonNotReceive
	id := c.Params("id")
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = err.Error()
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	db := configs.Store
	var obj models.CartonNotReceive
	err = db.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.IsSync = frm.IsSync
	db.Save(&obj)
	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &obj

	// After Save Delete All Data When IsSync is true
	db.Where("is_sync=?", true).Delete(&models.CartonNotReceive{})
	return c.Status(fiber.StatusOK).JSON(&r)
}
