package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllLedger(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Ledger
	// Fetch All Data
	err := configs.Store.
		Preload("Whs").
		Preload("Part").
		Preload("PartType").
		Preload("Unit").
		Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("Ledger")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = services.MessageShowAll("Ledger")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateLedger(c *fiber.Ctx) error {
	var r models.Response
	var obj models.Ledger
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	// Fetch All Data
	db := configs.Store
	// Get Master Data
	var whs models.Whs
	db.First(&whs, "title=?", &obj.WhsID)
	var part models.Part
	db.First(&part, "slug=?", &obj.PartID)
	var partType models.PartType
	db.First(&partType, "title=?", &obj.PartTypeID)
	var unit models.Unit
	db.First(&unit, "title=?", &obj.UnitID)

	// Add Variable
	obj.WhsID = &whs.ID
	obj.PartID = &part.ID
	obj.PartTypeID = &partType.ID
	obj.UnitID = &unit.ID

	db.Where("whs_id=?", &whs.ID).Where("part_id=?", &part.ID).First(&obj)
	if obj.ID != "" {
		r.Message = services.MessageDuplicateData(&part.Slug)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&part.Slug)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowLedgerByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Ledger
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

func UpdateLedgerByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Ledger
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.Ledger
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

func DeleteLedgerByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.Ledger
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
