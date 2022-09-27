package controllers

import (
	"fmt"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllFileEdi(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.FileEdi
	// Fetch All Data
	err := configs.Store.Find(&obj).Error
	if err != nil {
		r.Message = services.MessageNotFound("FileEdi")
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowAll("FileEdi")
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateFileEdi(c *fiber.Ctx) error {
	var r models.Response
	// Create FileEdi
	db := configs.Store
	var obj models.FileEdi
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var objDup models.FileEdi
	db.Select("batch_no").First(&objDup, "batch_no=?", obj.BatchNo)

	// Upload GEDI File To Directory
	file, err := c.FormFile("file_edi")
	if err != nil {
		r.Message = services.MessageUploadFileError(err.Error())
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	fName := fmt.Sprintf("%s.%s", obj.BatchNo, file.Filename)
	obj.BatchPath = fmt.Sprintf("edi/%s", fName)
	err = c.SaveFile(file, fmt.Sprintf("./public/edi/%s", fName))
	if err != nil {
		r.Message = services.MessageSystemErrorNotSaveFile
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	/// Check File Type
	objName := strings.ReplaceAll(fmt.Sprint(file.Filename[0:13]), " ", "")
	FileTypeID := "R"
	if fmt.Sprint(objName[0:5]) == "OES.V" {
		FileTypeID = "O"
	}

	FactoryID := "AW"
	if objName[12:13] == "5" {
		FactoryID = "INJ"
	}
	/// End Check File Type
	var factory models.Factory
	err = db.Where("title=?", FactoryID).First(&factory).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&FactoryID)
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var mailbox models.Mailbox
	err = db.Preload("Area").Where("mailbox=?", obj.MailboxID).First(&mailbox).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(obj.MailboxID)
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var filetype models.FileType
	err = db.Where("title=?", FileTypeID).First(&filetype).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&FileTypeID)
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	// Create GEDI file
	obj.FactoryID = &factory.ID
	obj.MailboxID = &mailbox.ID
	obj.FileTypeID = &filetype.ID
	obj.IsDownload = true
	obj.Size = file.Size
	obj.BatchName = file.Filename

	err = db.FirstOrCreate(&obj, models.FileEdi{BatchNo: obj.BatchNo}).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Goroutines Read GEDI files
	obj.Factory = factory
	obj.Mailbox = mailbox
	obj.FileType = filetype
	// // Create upload log
	logData := models.SyncLogger{
		Title:       "upload gedi file",
		Description: fmt.Sprintf("Uploaded %s is completed", obj.BatchNo),
		IsSuccess:   true,
	}

	r.Message = services.MessageCreatedData(&obj.ID)
	r.Data = &obj
	response := c.Status(fiber.StatusCreated).JSON(&r)
	if objDup.BatchNo != "" {
		logData = models.SyncLogger{
			Title:       "upload gedi file",
			Description: fmt.Sprintf("Duplicate edi batch number: %s", obj.BatchNo),
			IsSuccess:   true,
		}
		response = c.Status(fiber.StatusOK).JSON(&r)
	}

	// Read Text files
	if objDup.BatchNo == "" {
		go services.ReadGediFile(&obj)
	}
	db.Create(&logData)
	return response
}

func ShowFileEdiByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.FileEdi
	err := configs.Store.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusFound).JSON(&r)
}

func UpdateFileEdiByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.FileEdi
	err := c.BodyParser(&obj)
	if err != nil {
		r.Message = services.MessageInputValidationError
		r.Data = &err
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}
	// Fetch All Data
	db := configs.Store
	var data models.FileEdi
	err = db.First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	/// Save Data
	// data.Title = obj.Title
	// data.Description = obj.Description
	data.IsActive = obj.IsActive
	////
	err = db.Save(&data).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusAccepted).JSON(&r)
}

func DeleteFileEdiByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	db := configs.Store
	var obj models.FileEdi
	err := db.First(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	err = db.Delete(&obj, &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageDeleteData(&id)
	r.Data = &obj
	return c.Status(fiber.StatusAccepted).JSON(&r)
}
