package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrder(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Order
	// Fetch All Data
	err := configs.Store.
		Limit(2).
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
		Preload("Shipment").
		Preload("Pc").
		Preload("Commercial").
		Preload("SampleFlg").
		Preload("OrderTitle").
		Preload("OrderDetail").
		Find(&obj).
		Error
	if err != nil {
		r.Message = services.MessageNotFound("Order Ent")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("Order Ent")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GenerateOrder(c *fiber.Ctx) error {
	var r models.Response
	factory := c.Query("factory")
	if factory == "" {
		factory = "INJ"
	}
	go services.CreateOrder(factory)
	r.Message = "Auto Generate Order"
	r.Data = nil
	return c.Status(fiber.StatusCreated).JSON(&r)
}
