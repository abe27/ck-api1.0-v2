package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllLastFTicket(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.LastFticket
	// Fetch All Data
	err := configs.Store.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("LastFTicket")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("LastFTicket")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateLastFTicket(c *fiber.Ctx) error {
	var r models.Response
	var obj models.LastFticket
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	err = configs.Store.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageDuplicateData(&obj.ID)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowLastFTicketByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.LastFticket
	err := configs.Store.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusFound).JSON(&r)
}

func UpdateLastFTicketByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.LastFticket
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.LastFticket
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Floors = obj.Floors
	// data.BoxSize = obj.BoxSize
	// data.PalletSize = obj.PalletSize
	// data.LimitTotal = obj.LimitTotal
	data.IsActive = obj.IsActive
	////
	err = db.Save(&data).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusAccepted).JSON(&r)
}

func DeleteLastFTicketByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.LastFticket
	err := db.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	err = db.Delete(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageDeleteData(&id)
	r.Data = &obj
	return c.Status(fiber.StatusAccepted).JSON(&r)
}
