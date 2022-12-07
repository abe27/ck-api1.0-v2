package controllers

import (
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllReceiveEnt(c *fiber.Ctx) error {
	etd := time.Now().Format("2006-01-02")
	if c.Query("etd") != "" {
		etd = c.Query("etd")
	}

	isSync := false
	if c.Query("is_sync") != "" {
		isSync = true
	}

	var r models.Response
	var obj []models.Receive
	db := configs.Store
	err := db.
		Limit(25).
		Order("receive_date desc,transfer_out_no").
		Preload("FileEdi.Factory").
		Preload("FileEdi.Mailbox.Area").
		Preload("FileEdi.FileType").
		Preload("ReceiveType.Whs").
		Preload("ReceiveDetail.Ledger.Part").
		Preload("ReceiveDetail.Ledger.PartType").
		Preload("ReceiveDetail.Ledger.Unit").
		Preload("ReceiveDetail.CartonNotReceive").
		Where("receive_date=?", etd).
		Where("is_sync=?", isSync).
		Find(&obj).
		Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &r
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowAllData("receive")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowReceiveEntByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateReceiveEntByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var frm models.ReceiveEntForm
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = err.Error()
		r.Data = &r
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store
	var receEnt models.Receive
	err = db.First(&receEnt, "id=?", &id).Error
	if err != nil {
		r.Message = err.Error()
		r.Data = &r
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	receEnt.IsSync = frm.IsSync
	receEnt.IsActive = frm.IsActive
	db.Save(&receEnt)
	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &receEnt
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteReceiveEntByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
