package controllers

import (
	"strconv"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrder(c *fiber.Ctx) error {
	db := configs.Store
	var r models.Response
	var obj []models.Order
	etd := c.Query("etd")
	if etd != "" {
		var facData models.Factory
		db.First(&facData, "title=?", c.Query("factory"))
		if services.IsAdmin(c) {
			err := db.
				Order("etd_date,updated_at").
				Where("etd_date=?", etd).
				Preload("Consignee.Whs").
				Preload("Consignee.Factory").
				Preload("Consignee.Affcode").
				Preload("Consignee.Customer").
				Preload("Consignee.CustomerAddress").
				Preload("Consignee.OrderGroup.User").
				Preload("Shipment").
				Preload("Pc").
				Preload("Commercial").
				Preload("SampleFlg").
				Preload("OrderTitle").
				Preload("Pallet.PalletType").
				Preload("Pallet.PalletDetail").
				Preload("OrderDetail.Ledger.Part").
				Preload("OrderDetail.Ledger.PartType").
				Preload("OrderDetail.Ledger.Unit").
				Preload("OrderDetail.OrderPlan.FileEdi.Factory").
				Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
				Preload("OrderDetail.OrderPlan.FileEdi.FileType").
				Preload("OrderDetail.OrderPlan.FileEdi.FileType").
				Preload("OrderDetail.OrderPlan.ReviseOrder").
				Preload("OrderDetail.OrderPlan.OrderType").
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

		conID := services.GetOrderGroup(c)
		if c.Query("custname") != "" {
			var custData []models.Customer
			db.Select("id").Where("description like ?", "%"+strings.ToUpper(c.Query("custname"))+"%").Find(&custData)
			custID := []string{}
			for _, v := range custData {
				custID = append(custID, v.ID)
			}
			var consigneeData []models.Consignee
			db.Where("factory_id=?", &facData.ID).Where("customer_id in ?", custID).Find(&consigneeData)
			conID = []string{}
			for _, v := range consigneeData {
				conID = append(conID, v.ID)
			}
		}

		err := db.
			Order("etd_date,updated_at").
			Where("etd_date=?", etd).
			Where("consignee_id in ?", conID).
			Preload("Consignee.Whs").
			Preload("Consignee.Factory").
			Preload("Consignee.Affcode").
			Preload("Consignee.Customer").
			Preload("Consignee.CustomerAddress").
			Preload("Consignee.OrderGroup.User").
			Preload("Shipment").
			Preload("Pc").
			Preload("Commercial").
			Preload("SampleFlg").
			Preload("OrderTitle").
			Preload("Pallet.PalletType").
			Preload("Pallet.PalletDetail").
			Preload("OrderDetail.Ledger.Part").
			Preload("OrderDetail.Ledger.PartType").
			Preload("OrderDetail.Ledger.Unit").
			Preload("OrderDetail.OrderPlan.FileEdi.Factory").
			Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
			Preload("OrderDetail.OrderPlan.FileEdi.FileType").
			Preload("OrderDetail.OrderPlan.FileEdi.FileType").
			Preload("OrderDetail.OrderPlan.ReviseOrder").
			Preload("OrderDetail.OrderPlan.OrderType").
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

	start_etd := c.Query("start_etd")
	to_etd := c.Query("to_etd")
	isDownload, err := strconv.ParseBool(c.Query("status"))
	if err != nil {
		isDownload = false
	}
	if start_etd != "" {
		err := db.
			Order("etd_date").
			Select("id").
			Where("etd_date between ? and ?", start_etd, to_etd).
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
	// Fetch All Data
	err = db.
		Limit(100).
		Order("etd_date,updated_at").
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
		Preload("Consignee.OrderGroup.User").
		Preload("Shipment").
		Preload("Pc").
		Preload("Commercial").
		Preload("SampleFlg").
		Preload("OrderTitle").
		Preload("Pallet.PalletType").
		Preload("Pallet.PalletDetail").
		Preload("OrderDetail.Ledger.Part").
		Preload("OrderDetail.Ledger.PartType").
		Preload("OrderDetail.Ledger.Unit").
		Preload("OrderDetail.OrderPlan.FileEdi.Factory").
		Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.ReviseOrder").
		Preload("OrderDetail.OrderPlan.OrderType").
		Find(&obj, "is_sync=?", isDownload).
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

func ShowOrderByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	var obj models.Order
	err := configs.Store.
		Order("etd_date,updated_at").
		Preload("Consignee.Whs").
		Preload("Consignee.Factory").
		Preload("Consignee.Affcode").
		Preload("Consignee.Customer").
		Preload("Consignee.CustomerAddress").
		Preload("Consignee.OrderGroup.User").
		Preload("Shipment").
		Preload("Pc").
		Preload("Commercial").
		Preload("SampleFlg").
		Preload("OrderTitle").
		Preload("Pallet.PalletType").
		Preload("Pallet.PalletDetail").
		Preload("OrderDetail.Ledger.Part").
		Preload("OrderDetail.Ledger.PartType").
		Preload("OrderDetail.Ledger.Unit").
		Preload("OrderDetail.OrderPlan.FileEdi.Factory").
		Preload("OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.FileEdi.FileType").
		Preload("OrderDetail.OrderPlan.ReviseOrder").
		Preload("OrderDetail.OrderPlan.OrderType").
		Find(&obj, &id).
		Error
	if err != nil {
		r.Message = services.MessageNotFound(id)
		r.Data = &err
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&id)
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Order
	id := c.Params("id")
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var data models.Order
	db := configs.Store
	err = db.First(&data, "id=?", &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if data.ID == "" {
		r.Message = services.MessageNotFoundData(&id)
		r.Data = nil
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	data.IsSync = frm.IsSync
	data.IsActive = frm.IsActive
	db.Save(&data)
	// Update Order Status
	r.Message = services.MessageUpdateDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func GenerateOrder(c *fiber.Ctx) error {
	var r models.Response
	factory := c.Query("factory")
	if factory == "" {
		factory = "INJ"
	}
	db := configs.Store
	var fac models.Factory
	err := db.First(&fac, "title=?", factory).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var autoGen models.AutoGenerateInvoice
	err = db.First(&autoGen, "factory_id=?", &fac.ID).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if !autoGen.IsGenerate {
		r.Message = services.MessageShowNotAllow(fac.ID)
		r.Data = &err
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	go services.CreateOrder(factory)
	r.Message = "Auto Generate Order"
	r.Data = nil
	return c.Status(fiber.StatusCreated).JSON(&r)
}
