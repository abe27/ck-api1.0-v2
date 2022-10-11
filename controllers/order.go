package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrder(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Order
	// Fetch All Data
	err := configs.Store.
		Limit(100).
		Order("etd_date").
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
		Preload("Shipment").
		Preload("Pc").
		Preload("Commercial").
		Preload("SampleFlg").
		Preload("OrderTitle").
		Preload("Pallet.PalletType").
		Preload("Pallet.PalletDetail").
		Preload("OrderDetail.Ledger.Part").
		Preload("OrderDetail.Ledger.PartType").
		Preload("OrderDetail.Ledger.Unit").
		Preload("OrderDetail.OrderPlan.FileEdi.Factory").
		Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.ReviseOrder").
		Preload("OrderDetail.OrderPlan.OrderType").
		Find(&obj, "is_sync=?", false).
		Error
	if err != nil {
		r.Message = services.MessageNotFound("Order Ent")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("Order Ent")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Order
	id := c.Params("id")
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var data models.Order
	db := configs.Store
	err = db.First(&data, "id=?", &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if data.ID == "" {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = nil
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	data.IsSync = frm.IsSync
	data.IsActive = frm.IsActive
	db.Save(&data)
	// Update Order Status
	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GenerateOrder(c *fiber.Ctx) error {
	var r models.Response
	factory := c.Query("factory")
	if factory == "" {
		factory = "INJ"
	}
	db := configs.Store
	var fac models.Factory
	err := db.First(&fac, "title=?", factory).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var autoGen models.AutoGenerateInvoice
	err = db.First(&autoGen, "factory_id=?", &fac.ID).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if !autoGen.IsGenerate {
		r.Message = services.MessageShowNotAllow(fac.ID)
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	go services.CreateOrder(factory)
	r.Message = "Auto Generate Order"
	r.Data = nil
	return c.Status(fiber.StatusCreated).JSON(&r)
}
