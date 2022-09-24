package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllMailbox(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Mailbox
	// Fetch All Data
	err := configs.Store.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("Mailbox")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("Mailbox")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateMailbox(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Mailbox
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var area models.Area
	err = db.First(&area, "title=?", frm.AreaID).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(frm.AreaID)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var obj models.Mailbox
	obj.Mailbox = frm.Mailbox
	obj.Password = frm.Password
	obj.HostUrl = frm.HostUrl
	obj.AreaID = &area.ID
	obj.IsActive = frm.IsActive

	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageDuplicateData(&obj.Mailbox)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&obj.Mailbox)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowMailboxByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Mailbox
	err := configs.Store.Preload("Area").First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusFound).JSON(&r)
}

func UpdateMailboxByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Mailbox
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.Mailbox
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Title = obj.Title
	// data.Description = obj.Description
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

func DeleteMailboxByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.Mailbox
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
