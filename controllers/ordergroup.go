package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.OrderGroup
	// Fetch All Data
	err := configs.Store.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("OrderGroup")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("OrderGroup")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderGroup(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FormGroupConsignee
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	// Fetch Data Master
	var user models.User
	db.First(&user, &models.User{UserName: frm.UserID})
	var whs models.Whs
	db.First(&whs, &models.Whs{Title: frm.WhsID})
	var factory models.Factory
	db.First(&factory, &models.Factory{Title: frm.FactoryID})
	var orderType models.OrderGroupType
	db.First(&orderType, &models.OrderGroupType{Title: frm.OrderGroupTypeID})

	// Fetch Affcode
	var affcode models.Affcode
	db.First(&affcode, &models.Affcode{Title: frm.AffcodeID})
	var customer models.Customer
	db.First(&customer, &models.Customer{Title: frm.CustcodeID})

	// Fetch Consignee Data
	var consignee models.Consignee
	db.First(&consignee, &models.Consignee{
		WhsID:      &whs.ID,
		FactoryID:  &factory.ID,
		AffcodeID:  &affcode.ID,
		CustomerID: &customer.ID,
	})

	var obj models.OrderGroup
	obj.UserID = &user.ID
	obj.ConsigneeID = &consignee.ID
	obj.OrderGroupTypeID = &orderType.ID
	obj.SubOrder = frm.SubOrder
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive

	// Check duplicate
	var orderGroup models.OrderGroup
	db.Select("id,title").First(&orderGroup, &models.OrderGroup{
		UserID:           &user.ID,
		ConsigneeID:      &consignee.ID,
		OrderGroupTypeID: &orderType.ID,
		SubOrder:         frm.SubOrder,
	})

	if orderGroup.ID != "" {
		r.Message = services.MessageDuplicateData(&orderGroup.ID)
		r.Data = nil
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	err = db.Create(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageCreatedData(&obj.SubOrder)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowOrderGroupByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.OrderGroup
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

func UpdateOrderGroupByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.OrderGroup
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.OrderGroup
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

func DeleteOrderGroupByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.OrderGroup
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
