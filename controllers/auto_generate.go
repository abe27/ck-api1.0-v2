package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllAutoGenerateInvoice(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.AutoGenerateInvoice
	// Fetch All Data
	err := configs.Store.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("AutoGenerateInvoice")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("AutoGenerateInvoice")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateAutoGenerateInvoice(c *fiber.Ctx) error {
	var r models.Response
	var frm models.AutoGenerateInvoice
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var factory models.Factory
	db.First(&factory, "title=?", frm.FactoryID)

	var obj models.AutoGenerateInvoice
	obj.FactoryID = &factory.ID
	obj.IsGenerate = frm.IsGenerate
	obj.IsActive = frm.IsActive
	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageDuplicateData(&factory.ID)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&factory.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowAutoGenerateInvoiceByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.AutoGenerateInvoice
	err := configs.Store.Preload("Factory").First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusFound).JSON(&r)
}

func UpdateAutoGenerateInvoiceByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.AutoGenerateInvoice
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.AutoGenerateInvoice
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Title = obj.Title
	data.IsGenerate = obj.IsGenerate
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

func DeleteAutoGenerateInvoiceByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.AutoGenerateInvoice
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
