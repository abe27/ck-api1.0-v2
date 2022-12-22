package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func ImportInvoiceTap(c *fiber.Ctx) error {
	var r models.Response
	// Upload GEDI File To Directory
	file, err := c.FormFile("file")
	if err != nil {
		r.Message = services.MessageUploadFileError(err.Error())
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	fName := fmt.Sprintf("./public/invoices/%s", file.Filename)
	err = c.SaveFile(file, fName)
	if err != nil {
		r.Message = services.MessageSystemErrorNotSaveFile
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	//// Read Excel
	services.ImportInvoiceTap(&fName)
	r.Message = services.MessageUploadFileCompleted(fName)
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func GenerateFTicketInvoiceTap(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	//// Read Excel
	var invTap []models.ImportInvoiceTap
	if err := db.Limit(100).Where("is_matched=?", false).Find(&invTap).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	for _, v := range invTap {
		inv := v.Bhivno
		inv_seq, _ := strconv.ParseInt(inv[5:len(inv)-1], 10, 64)
		var facData models.Factory
		db.Select("id,title,inv_prefix,label_prefix").First(&facData, "inv_prefix=?", inv[:2])
		etd := v.Bhivdt.Format("2006-01-02")
		// fmt.Printf("%d ==> ETD: %s\n", line, etd)
		var shipment models.Shipment
		db.First(&shipment, "title=?", inv[len(inv)-1:])
		var orderPlan models.OrderPlan
		if err := db.Raw(fmt.Sprintf("select id,bal_qty,bistdp from tbt_order_plans where etd_tap='%s' and part_no='%s' and shipment_id='%s' and pono in ('%s','%s') order by created_at desc,seq desc limit 1", etd, v.Bhypat, shipment.ID, strings.Trim(v.Bhodpo, ""), strings.Trim(strings.ReplaceAll(v.Bhodpo, " ", ""), ""))).Scan(&orderPlan).Error; err == nil {
			services.CreateOrderPallet(&v, &orderPlan, inv_seq, fmt.Sprintf("%d", v.Bhwidt), fmt.Sprintf("%d", v.Bhleng), fmt.Sprintf("%d", v.Bhhigh), v.Bhpaln, fmt.Sprintf("%d", v.Bhctn), v.Bhivdt, &facData)
		}
	}
	r.Message = fmt.Sprintf("Generate FTicket No %d.", len(invTap))
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func CheckInvoiceTap(c *fiber.Ctx) error {
	var r models.Response
	services.CheckInvoiceTap()
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ClientImportInvoiceTap(c *fiber.Ctx) error {
	var r models.Response
	// Upload GEDI File To Directory
	file, err := c.FormFile("file")
	if err != nil {
		r.Message = services.MessageUploadFileError(err.Error())
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	fName := fmt.Sprintf("./public/invoices/%s", file.Filename)
	err = c.SaveFile(file, fName)
	if err != nil {
		r.Message = services.MessageSystemErrorNotSaveFile
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	//// Read Excel
	// services.ImportInvoiceTap(&fName)
	go services.ImportInvoiceTap(&fName)

	r.Message = services.MessageUploadFileCompleted(fName)
	return c.Status(fiber.StatusCreated).JSON(&r)
}
