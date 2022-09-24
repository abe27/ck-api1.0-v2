package controllers

import (
	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllAdministrator(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Administrator
	err := configs.Store.Preload("User").Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("Administrator")
		r.Data = nil
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAllData("Administrator")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateAdministrator(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Administrator
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store
	// FetchData UserID
	var user models.User
	err = db.First(&user, "username=?", &frm.UserID).Error
	if err != nil {
		r.Message = services.MessageNotFound("UserID")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// When Found User
	admin := models.Administrator{
		UserID:   &user.ID,
		IsActive: frm.IsActive,
	}
	err = db.Create(&admin).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageCreatedData(&user.ID)
	r.Data = &admin
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowAdministratorByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Administrator
	err := configs.Store.Preload("User").First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusFound).JSON(&r)
}

func UpdateAdministratorByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Administrator
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store
	id := c.Params("id")
	var obj models.Administrator
	err = db.Preload("User").First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.IsActive = frm.IsActive

	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusAccepted).JSON(&r)
}

func DeleteAdministratorByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Administrator
	err := configs.Store.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	err = configs.Store.Delete(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageDeleteData(&id)
	return c.Status(fiber.StatusAccepted).JSON(&r)
}
