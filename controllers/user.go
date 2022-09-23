package controllers

import (
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var r models.Response
	var obj models.UserForm
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	password := services.HashingPassword(obj.Password)
	isMatch := services.CheckPasswordHashing(obj.Password, password)
	if !isMatch {
		r.Message = services.MessagePasswordNotMatched
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	obj.Password = password

	// Create New User
	userData := &models.User{
		UserName: obj.UserName,
		Email:    obj.Email,
		Password: obj.Password,
	}
	db := configs.Store
	err = db.Create(&userData).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create Profile
	profileData := models.Profile{
		UserID:    &userData.ID,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
	}

	db.FirstOrCreate(&profileData, &models.Profile{UserID: &userData.ID})

	r.Message = services.MessageRegister(obj.UserName)
	r.Data = &userData
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func Login(c *fiber.Ctx) error {
	var r models.Response
	var user models.UserLoginForm
	err := c.BodyParser(&user)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	// Check AuthorizationRequired
	db := configs.Store
	var userData models.User
	err = db.Where("username=?", user.UserName).First(&userData).Error
	if err != nil {
		r.Message = services.MessageNotFoundUser
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if !userData.IsActive {
		r.Message = services.MessageUserNotActive
		r.Data = nil
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	isMatched := services.CheckPasswordHashing(c.FormValue("password"), userData.Password)
	if !isMatched {
		r.Message = services.MessagePasswordNotMatch
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Create Token
	auth := services.CreateToken(userData)
	r.Message = services.MessageAuthentication
	r.Data = &auth
	return c.Status(fiber.StatusOK).JSON(&r)
}

func Verify(c *fiber.Ctx) error {
	var r models.Response
	// Delete Token
	db := configs.Store
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	err := db.Where("id=?", token).First(&models.JwtToken{}).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusUnauthorized).JSON(&r)
	}
	r.Message = services.MessageAuthentication
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}

func Profile(c *fiber.Ctx) error {
	var r models.Response
	// Delete Token
	db := configs.Store
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	var jwtToken models.JwtToken
	err := db.Select("user_id").Where("id=?", token).First(&jwtToken).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(token)
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}

func Logout(c *fiber.Ctx) error {
	var r models.Response
	// Delete Token
	db := configs.Store
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	err := db.Where("id=?", token).Delete(&models.JwtToken{}).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageUserLeave
	r.Data = nil
	return c.Status(fiber.StatusOK).JSON(&r)
}
