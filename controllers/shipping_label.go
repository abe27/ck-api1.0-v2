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
	var frm []models.PrintShippingLabel
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store
	var data []models.PrintShippingLabel
	for _, p := range frm {
		obj := models.PrintShippingLabel{
			InvoiceNo:    p.InvoiceNo,
			OrderNo:      p.OrderNo,
			PartNo:       p.PartNo,
			CustCode:     p.CustCode,
			CustName:     p.CustName,
			PalletNo:     p.PalletNo,
			PrintDate:    p.PrintDate,
			Qty:          p.Qty,
			QrCode:       fmt.Sprintf("06P%s;17Q%d;30T%s;32T%s;", p.PartNo, p.Qty, p.OrderNo, p.BarCode),
			BarCode:      p.BarCode,
			LabelBarCode: fmt.Sprintf("*%s*", p.BarCode),
			IsPrint:      p.IsPrint,
		}
		err = db.FirstOrCreate(&obj, &models.PrintShippingLabel{BarCode: p.BarCode}).Error
		if err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		obj.Qty = p.Qty
		obj.QrCode = fmt.Sprintf("06P%s;17Q%d;30T%s;32T%s;", p.PartNo, p.Qty, p.OrderNo, p.BarCode)
		obj.IsPrint = 0
		err = db.Save(&obj).Error
		if err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		data = append(data, obj)
	}

	txt := fmt.Sprintf("%d รายการ", len(data))
	if len(data) == 1 {
		txt = data[0].BarCode
	}
	r.Message = services.MessagePrintShippingLabel(txt)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}
