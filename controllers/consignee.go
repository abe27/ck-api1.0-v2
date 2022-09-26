package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllConsignee(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Consignee
	// Fetch All Data
	err := configs.Store.
		Preload("Whs").
		Preload("Factory").
		Preload("Affcode").
		Preload("Customer").
		Preload("CustomerAddress").
		Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("Consignee")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("Consignee")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateConsignee(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Consignee
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	// Check Duplicate
	var whs models.Whs
	db.Select("id,title").First(&whs, "title=?", frm.WhsID)
	var factory models.Factory
	db.Select("id,title").First(&factory, "title=?", frm.FactoryID)
	var affcode models.Affcode
	db.Select("id,title").First(&affcode, "title=?", frm.AffcodeID)
	var customer models.Customer
	db.Select("id,title").First(&customer, "title=?", frm.CustomerID)
	var customerAddress models.CustomerAddress
	db.Select("id,title").First(&customerAddress, "title=?", frm.CustomerAddressID)

	var consigneeData models.Consignee
	db.Select("id").
		Where("whs_id=?", &whs.ID).
		Where("factory_id=?", &factory.ID).
		Where("affcode_id=?", &affcode.ID).
		Where("customer_id=?", &customer.ID).
		First(&consigneeData)

	if consigneeData.ID != "" {
		r.Message = services.MessageDuplicateData(&consigneeData.ID)
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	// not duplicate
	var obj models.Consignee
	obj.WhsID = &whs.ID
	obj.FactoryID = &factory.ID
	obj.AffcodeID = &affcode.ID
	obj.CustomerID = &customer.ID
	obj.CustomerAddressID = &customerAddress.ID
	obj.Prefix = frm.Prefix
	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&obj.Prefix)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowConsigneeByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Consignee
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

func UpdateConsigneeByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Consignee
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	// Check Duplicate
	var whs models.Whs
	db.Select("id,title").First(&whs, "title=?", obj.WhsID)
	var factory models.Factory
	db.Select("id,title").First(&factory, "title=?", obj.FactoryID)
	var affcode models.Affcode
	db.Select("id,title").First(&affcode, "title=?", obj.AffcodeID)
	var customer models.Customer
	db.Select("id,title").First(&customer, "title=?", obj.CustomerID)
	var customerAddress models.CustomerAddress
	db.Select("id,title").First(&customerAddress, "title=?", obj.CustomerAddressID)

	var data models.Consignee
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	obj.WhsID = &whs.ID
	obj.FactoryID = &factory.ID
	obj.AffcodeID = &affcode.ID
	obj.CustomerID = &customer.ID
	obj.CustomerAddressID = &customerAddress.ID
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

func DeleteConsigneeByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.Consignee
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
