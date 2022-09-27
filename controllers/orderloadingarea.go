package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func ShowAllOrderLoadingArea(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.OrderLoadingArea
	// Fetch All OrderLoadingArea
	db := configs.Store
	err := db.Preload("OrderZone").Find(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowAllData("OrderLoadingArea Data")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderLoadingArea(c *fiber.Ctx) error {
	var r models.Response
	var obj models.OrderLoadingAreaForm
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create OrderLoadingArea
	db := configs.Store
	// Fetch Factory
	var fac models.Factory
	db.Where("title=?", obj.Factory).First(&fac)

	// Fetch Order Zone
	var orderZone models.OrderZone
	db.Where("value=?", obj.Bioat).Where("factory_id=?", fac.ID).First(&orderZone)

	var loadingArea models.OrderLoadingArea
	loadingArea.OrderZoneID = &orderZone.ID
	loadingArea.Prefix = obj.Prefix
	loadingArea.LoadingArea = obj.LoadingArea
	loadingArea.Privilege = obj.Privilege
	loadingArea.IsActive = true

	var checkLoadingDuplicate models.OrderLoadingArea
	db.Select("id").Where("order_zone_id=?", &orderZone.ID).Where("prefix=?", obj.Prefix).First(&checkLoadingDuplicate)
	if checkLoadingDuplicate.ID != "" {
		r.Message = services.MessageDuplicateData(&checkLoadingDuplicate.ID)
		r.Data = nil
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	err = db.Create(&loadingArea).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&obj.ID)
	r.Data = &loadingArea
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowOrderLoadingAreaByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.OrderLoadingArea
	// Fetch By ID
	db := configs.Store
	err := db.Where("value=?", &id).First(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderLoadingAreaByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.OrderLoadingArea
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Update Data
	db := configs.Store
	err = db.Where("id=?", &id).Updates(&models.OrderLoadingArea{
		IsActive: obj.IsActive,
	}).Error

	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	err = db.Where("id=?", &id).First(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteOrderLoadingAreaByID(c *fiber.Ctx) error {
	var r models.Response
	// Update Data
	id := c.Params("id")
	db := configs.Store
	err := db.Where("id=?", &id).First(&models.OrderLoadingArea{}).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	err = db.Where("id=?", &id).Delete(&models.OrderLoadingArea{}).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageDeleteData(&id)
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}
