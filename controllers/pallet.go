package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrderPallet(c *fiber.Ctx) error {
	var r models.Response
	id := c.Query("order_id")
	if id == "" {
		r.Message = services.MessageNotFoundData(&id)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var data []models.Pallet
	err := configs.Store.Where("order_id=?", &id).
		Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("PalletType").
		Preload("PalletDetail.OrderDetail.Order").
		Preload("PalletDetail.OrderDetail.Ledger.Whs").
		Preload("PalletDetail.OrderDetail.Ledger.Factory").
		Preload("PalletDetail.OrderDetail.Ledger.Part").
		Preload("PalletDetail.OrderDetail.Ledger.PartType").
		Preload("PalletDetail.OrderDetail.Ledger.Unit").
		Preload("PalletDetail.OrderDetail.OrderPlan").
		Find(&data).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderPallet(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Pallet
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	err = configs.Store.Create(&frm).Error
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&frm.ID)
	r.Data = &frm
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowOrderPalletByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	if id == "" {
		r.Message = services.MessageNotFoundData(&id)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var data models.Pallet
	err := configs.Store.Where("id=?", &id).
		Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("PalletType").
		Preload("PalletDetail.OrderDetail.Order").
		Preload("PalletDetail.OrderDetail.Ledger.Whs").
		Preload("PalletDetail.OrderDetail.Ledger.Factory").
		Preload("PalletDetail.OrderDetail.Ledger.Part").
		Preload("PalletDetail.OrderDetail.Ledger.PartType").
		Preload("PalletDetail.OrderDetail.Ledger.Unit").
		Preload("PalletDetail.OrderDetail.OrderPlan").
		First(&data).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderPalletByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteOrderPalletByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	if id == "" {
		r.Message = services.MessageNotFoundData(&id)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	db := configs.Store
	err := db.Where("pallet_id=?", id).Delete(&models.PalletDetail{}).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	err = db.Where("id=?", id).Delete(&models.Pallet{}).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	return c.Status(fiber.StatusOK).JSON(&r)
}
