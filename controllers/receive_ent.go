package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllReceiveEnt(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Receive
	db := configs.Store
	err := db.
		Limit(2).
		Order("receive_date,transfer_out_no").
		Preload("FileEdi.Factory").
		Preload("FileEdi.Mailbox.Area").
		Preload("FileEdi.FileType").
		Preload("ReceiveType.Whs").
		Preload("ReceiveDetail.Ledger.Part").
		Preload("ReceiveDetail.Ledger.PartType").
		Preload("ReceiveDetail.Ledger.Unit").
		Find(&obj, "is_sync=?", false).
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
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteReceiveEntByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
