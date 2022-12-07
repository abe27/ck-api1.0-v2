package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrderPlanToSync(c *fiber.Ctx) error {
	var r models.Response
	var orderPlan []models.OrderPlan
	if err := configs.Store.
		Limit(2000).
		Order("upddte,updtime").
		Where("is_sync=?", false).
		Preload("FileEdi.Factory").
		Preload("FileEdi.Mailbox.Area").
		Preload("FileEdi.FileType").
		Preload("Whs").
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
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
		Preload("OrderZone.Whs").
		Preload("OrderZone.Factory").
		Preload("SampleFlg").
		Find(&orderPlan).Error; err != nil {
		r.Message = services.MessageSystemErrorWith(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageSystemWithMessage("History")
	r.Data = &orderPlan
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GetAllOrderPlan(c *fiber.Ctx) error {
	var r models.Response
	vendor := c.Query("vendor")
	if vendor == "" {
		r.Message = services.MessageSystemErrorWith("vendor")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	tagrp := c.Query("tagrp")
	if tagrp == "" {
		r.Message = services.MessageSystemErrorWith("tagrp")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	pono := c.Query("pono")
	if pono == "" {
		r.Message = services.MessageSystemErrorWith("pono")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	part_no := c.Query("part_no")
	if part_no == "" {
		r.Message = services.MessageSystemErrorWith("part_no")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	bishpc := c.Query("bishpc")
	if bishpc == "" {
		r.Message = services.MessageSystemErrorWith("bishpc")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	biac := c.Query("biac")
	if biac == "" {
		r.Message = services.MessageSystemErrorWith("biac")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var orderPlan []models.OrderPlan
	if err := configs.Store.
		Order("upddte desc,updtime desc").
		Where("vendor=?", vendor).
		Where("tagrp=?", tagrp).
		Where("pono=?", pono).
		Where("biac=?", biac).
		Where("part_no=?", part_no).
		Where("bishpc=?", bishpc).
		Preload("FileEdi.Factory").
		Preload("FileEdi.Mailbox.Area").
		Preload("FileEdi.FileType").
		Preload("Whs").
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
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
		Preload("OrderZone.Whs").
		Preload("OrderZone.Factory").
		Preload("SampleFlg").
		Find(&orderPlan).Error; err != nil {
		r.Message = services.MessageSystemErrorWith(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageSystemWithMessage("History")
	r.Data = &orderPlan
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderPlan(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowOrderPlanByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderPlanByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.FrmOrderPlan
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := configs.Store.Where("id=?", c.Params("id")).Updates(&models.OrderPlan{IsSync: frm.IsSync, IsActive: frm.IsActive, RowID: frm.RowID}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = "Update Order Plan ID: " + c.Params("id")
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteOrderPlanByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
