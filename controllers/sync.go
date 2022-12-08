package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func GetSyncOrderList(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Order
	if err := configs.Store.
		Limit(10).
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
