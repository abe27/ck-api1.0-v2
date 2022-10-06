package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllLineNotifyToken(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.LineNotifyToken
	// Fetch All Data
	err := configs.Store.Preload("Whs").Preload("Factory").Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("LineNotifyToken")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("LineNotifyToken")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateLineNotifyToken(c *fiber.Ctx) error {
	var r models.Response
	var obj models.LineNotifyToken
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

	var factory models.Factory
	db.First(&factory, "title=?", obj.FactoryID)

	var frm models.LineNotifyToken
	frm.WhsID = &whs.ID
	frm.FactoryID = &factory.ID
	frm.Token = obj.Token
	frm.IsActive = true
	err = db.Create(&frm).Error
	if err != nil {
		r.Message = services.MessageDuplicateData(&frm.Token)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&frm.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowLineNotifyTokenByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.LineNotifyToken
	err := configs.Store.Preload("Whs").Preload("Factory").First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusFound).JSON(&r)
}

func UpdateLineNotifyTokenByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.LineNotifyToken
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.LineNotifyToken
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

func DeleteLineNotifyTokenByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.LineNotifyToken
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
