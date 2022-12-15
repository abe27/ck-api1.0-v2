package controllers

import (
	"fmt"

	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func UploadReceiveExcel(c *fiber.Ctx) error {
	var r models.Response
	// Upload GEDI File To Directory
	file, err := c.FormFile("file")
	if err != nil {
		r.Message = services.MessageUploadFileError(err.Error())
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	fName := fmt.Sprintf("./public/excels/%s", file.Filename)
	err = c.SaveFile(file, fName)
	if err != nil {
		r.Message = services.MessageSystemErrorNotSaveFile
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	//// Read Excel
	services.ImportReceiveCarton(fName)

	r.Message = services.MessageUploadFileCompleted(fName)
	return c.Status(fiber.StatusCreated).JSON(&r)
}
