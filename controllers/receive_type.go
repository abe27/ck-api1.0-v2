package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllReceiveType(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.ReceiveType
	// Fetch All Data
	err := configs.Store.Preload("Whs").Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("ReceiveType")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("ReceiveType")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateReceiveType(c *fiber.Ctx) error {
	var r models.Response
	var obj models.ReceiveType
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var whs models.Whs
	db.First(&whs, "title=?", obj.WhsID)
	obj.WhsID = whs.ID
	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageDuplicateData(&obj.Title)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&obj.Title)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowReceiveTypeByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.ReceiveType
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

func UpdateReceiveTypeByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.ReceiveType
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.ReceiveType
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Value = obj.Value
	data.Description = obj.Description
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

func DeleteReceiveTypeByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.ReceiveType
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
