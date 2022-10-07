package controllers

import (
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllCartonNotReceive(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.CartonNotReceive
	err := configs.Store.Limit(500).Order("transfer_out_no,lot_no,serial_no").Where("is_sync=?", false).Find(&obj).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageShowAllData("Carton not receive")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateCartonNotReceiveByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.CartonNotReceive
	id := c.Params("id")
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = err.Error()
		r.Data = &err
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	db := configs.Store
	var obj models.CartonNotReceive

	err = db.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var recEnt models.Receive
	db.Preload("FileEdi.Factory").Preload("ReceiveType.Whs").First(&recEnt, "transfer_out_no=?", obj.TransferOutNo)
	if recEnt.ID != "" {
		var part models.Part
		db.Select("id").First(&part, "slug=?", strings.ReplaceAll(obj.PartNo, "-", ""))
		var ledger models.Ledger
		db.Select("id").Where("whs_id=?", &recEnt.ReceiveType.WhsID).Where("factory_id=?", &recEnt.FileEdi.FactoryID).Where("part_id=?", &part.ID).First(&ledger)

		// Search Receive Detail
		var receiveDetail models.ReceiveDetail
		db.Select("id").Where("receive_id=?", &recEnt.ID).Where("ledger_id=?", &ledger.ID).First(&receiveDetail)

		var carton models.Carton
		db.First(&carton, "serial_no=?", obj.SerialNo)
		if carton.ID != "" {
			carton.ReceiveDetailID = &receiveDetail.ID
			db.Save(&carton)
		}
	}

	// fmt.Printf("ID: %s serial: %s\n", recEnt.ID, obj.SerialNo)

	obj.IsSync = frm.IsSync
	db.Save(&obj)
	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &obj

	// After Save Delete All Data When IsSync is true
	err = db.Where("is_sync=?", true).Delete(&models.CartonNotReceive{}).Error
	if err != nil {
		r.Message = err.Error()
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	go services.SummaryReceive(&recEnt.ID)
	return c.Status(fiber.StatusOK).JSON(&r)
}
