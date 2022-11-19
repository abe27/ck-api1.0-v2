package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func ShowAllOrderShort(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderShort(c *fiber.Ctx) error {
	var r models.Response
	var frm models.OrderShort
	if err := c.BodyParser(&frm); err != nil {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	db := configs.Store
	var orderPlan models.OrderPlan
	if err := db.First(&orderPlan, &frm.OrderPlanID).Error; err != nil {
		r.Message = services.MessageNotFound(frm.OrderPlanID)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	/// ค้นหาตารางงานว่าลูกค้านี้ออกวันไหน schedule plan

	r.Message = services.MessageCreatedData(&frm.OrderPlanID)
	r.Data = &orderPlan
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowOrderShortByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderShortByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteOrderShortByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
