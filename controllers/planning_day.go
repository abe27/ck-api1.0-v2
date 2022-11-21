package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllPlanningDay(c *fiber.Ctx) error {
	var r models.Response
	var data []models.PlanningDay
	err := configs.Store.Where("is_active=?", true).Find(&data).Error
	if err != nil {
		r.Message = services.MessageShowAll(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowAll("Planning Day")
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreatePlanningDay(c *fiber.Ctx) error {
	var r models.Response
	var data models.PlanningDay
	err := c.BodyParser(&data)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Error = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	err = configs.Store.Create(&data).Error
	if err != nil {
		r.Message = services.MessageCreatedData(&data.Title)
		r.Error = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&data.Title)
	r.Data = &data
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowPlanningDayByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	if id == "" {
		r.Message = services.MessageRequireField(id)
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	var data models.PlanningDay
	err := configs.Store.Where("id=?", id).First(&data).Error
	if err != nil {
		r.Message = services.MessageShowAll(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdatePlanningDayByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeletePlanningDayByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}
