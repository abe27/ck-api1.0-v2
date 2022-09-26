package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllReviseOrder(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.ReviseOrder
	// Fetch All Data
	err := configs.Store.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("ReviseOrder")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("ReviseOrder")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateReviseOrder(c *fiber.Ctx) error {
	var r models.Response
	var obj models.ReviseOrder
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	err = configs.Store.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageDuplicateData(&obj.Title)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&obj.Title)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowReviseOrderByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.ReviseOrder
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

func UpdateReviseOrderByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.ReviseOrder
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.ReviseOrder
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Title = obj.Title
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

func DeleteReviseOrderByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.ReviseOrder
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
