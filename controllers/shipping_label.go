package controllers

import (
	"fmt"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllShippingLabel(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateShippingLabel(c *fiber.Ctx) error {
	var r models.Response
	var Shipping models.PostShippingLabel
	err := c.BodyParser(&Shipping)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	etd := (time.Now()).Format("2006-01-02")
	fmt.Println(etd)

	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreatePrintLabel(c *fiber.Ctx) error {
	var r models.Response
	var frm models.PrintShippingLabel
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store

	obj := models.PrintShippingLabel{
		InvoiceNo:    frm.InvoiceNo,
		OrderNo:      frm.OrderNo,
		PartNo:       frm.PartNo,
		CustCode:     frm.CustCode,
		CustName:     frm.CustName,
		PalletNo:     frm.PalletNo,
		PrintDate:    frm.PrintDate,
		QrCode:       fmt.Sprintf("06P%s;17Q%d;30T%s;32T%s;", frm.PartNo, frm.Qty, frm.OrderNo, frm.BarCode),
		BarCode:      frm.BarCode,
		LabelBarCode: fmt.Sprintf("*%s*", frm.BarCode),
		IsPrint:      frm.IsPrint,
	}

	err = db.FirstOrCreate(&obj, &models.PrintShippingLabel{BarCode: frm.BarCode}).Error
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	obj.QrCode = fmt.Sprintf("06P%s;17Q%d;30T%s;32T%s;", frm.PartNo, frm.Qty, frm.OrderNo, frm.BarCode)
	obj.IsPrint = 0
	err = db.Save(&obj).Error
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&frm.BarCode)
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}
