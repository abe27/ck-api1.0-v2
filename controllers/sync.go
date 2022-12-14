package controllers

import (
	"fmt"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetSyncOrderPlan(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.OrderPlan
	if err := configs.Store.
		Limit(10000).
		Order("upddte,updtime").
		Preload("FileEdi").
		Preload("Whs").
		Preload("Consignee").
		Preload("ReviseOrder").
		Preload("Ledger.Whs").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.PartType").
		Preload("Ledger.Unit").
		Preload("Pc").
		Preload("Commercial").
		Preload("OrderType").
		Preload("Shipment").
		Preload("OrderZone").
		Preload("SampleFlg").
		Where("is_sync=?", false).
		Find(&obj).
		Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.Message = "Show Order Plan"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateSyncOrderPlan(c *fiber.Ctx) error {
	var r models.Response
	var frm models.OrderPlan
	if err := c.BodyParser(&frm); err != nil {
		panic(err)
	}

	if err := configs.Store.Model(&models.OrderPlan{}).Where("id=?", c.Params("id")).Update("is_sync", &frm.IsSync).Error; err != nil {
		panic(err)
	}
	r.Message = fmt.Sprintf("Update %s is Success!", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GetSyncOrderList(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Order
	if err := configs.Store.
		Limit(50).
		Order("etd_date,updated_at").
		Where("is_checked=?", true).
		Where("is_sync=?", false).
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
		Preload("Consignee.OrderGroup.User").
		Preload("Shipment").
		Preload("Pc").
		Preload("Commercial").
		Preload("SampleFlg").
		Preload("OrderTitle").
		Preload("Pallet.PalletType").
		Preload("Pallet.PalletDetail.OrderDetail.Ledger.Factory").
		Preload("Pallet.PalletDetail.OrderDetail.Ledger.Part").
		Preload("OrderDetail.Ledger.Whs").
		Preload("OrderDetail.Ledger.Factory").
		Preload("OrderDetail.Ledger.Part").
		Preload("OrderDetail.Ledger.PartType").
		Preload("OrderDetail.Ledger.Unit").
		Preload("OrderDetail.OrderPlan.FileEdi.Factory").
		Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.ReviseOrder").
		Preload("OrderDetail.OrderPlan.OrderType").
		Preload("OrderDetail.OrderPlan.FileEdi.Factory").
		Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.Whs").
		Preload("OrderDetail.OrderPlan.OrderZone").
		Preload("OrderDetail.OrderPlan.SampleFlg").
		Find(&obj).
		Error; err != nil {
		panic(err)
	}
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderSyncByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.OrderSyncForm
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := configs.Store.Model(models.Order{}).Where("id=?", c.Params("id")).Update("is_sync", frm.IsSync).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Update `%s` is Success", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
