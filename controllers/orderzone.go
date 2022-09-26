package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrderZone(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.OrderZone
	// Fetch All Data
	err := configs.Store.Preload("Factory").Preload("Whs").Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("OrderZone")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = services.MessageShowAll("OrderZone")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderZone(c *fiber.Ctx) error {
	var r models.Response
	var obj models.OrderZone
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	// Fetch All Data
	db := configs.Store
	var factoryData models.Factory
	db.First(&factoryData, "title=?", obj.FactoryID)
	var whsData models.Whs
	db.First(&whsData, "title=?", obj.WhsID)

	/// Insert
	obj.FactoryID = &factoryData.ID
	obj.WhsID = &whsData.ID
	obj.Description = whsData.Title

	// Check Duplicate Data
	var orderZone models.OrderZone
	db.
		Where("value=?", obj.Value).
		Where("factory_id=?", &factoryData.ID).
		Where("whs_id=?", &whsData.ID).First(&orderZone)

	if orderZone.ID != "" {
		r.Message = services.MessageDuplicateData(&obj.ID)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowOrderZoneByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.OrderZone
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

func UpdateOrderZoneByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.OrderZone
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.OrderZone
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Title = obj.Title
	var factoryData models.Factory
	db.First(&factoryData, "title=?", obj.FactoryID)
	var whsData models.Whs
	db.First(&whsData, "title=?", obj.WhsID)

	/// Insert
	data.FactoryID = &factoryData.ID
	data.WhsID = &whsData.ID
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

func DeleteOrderZoneByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.OrderZone
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
