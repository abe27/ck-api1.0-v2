package controllers

import (
	"fmt"
	"time"

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

	var ReviseOrder models.ReviseOrder
	err := db.Where("title=?", "H").First(&ReviseOrder).Error
	if err != nil {
		r.Message = services.MessageNotFound("ReviseOrder")
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	/// ค้นหาตารางงานว่าลูกค้านี้ออกวันไหน schedule plan
	var orderShort models.OrderPlan
	orderShort.WhsID = orderPlan.WhsID
	orderShort.OrderZoneID = orderPlan.OrderZoneID
	orderShort.ConsigneeID = orderPlan.ConsigneeID
	orderShort.ReviseOrderID = &ReviseOrder.ID
	orderShort.LedgerID = orderPlan.LedgerID
	orderShort.PcID = orderPlan.PcID
	orderShort.CommercialID = orderPlan.CommercialID
	orderShort.OrderTypeID = orderPlan.OrderTypeID
	orderShort.ShipmentID = orderPlan.ShipmentID
	orderShort.SampleFlgID = orderPlan.SampleFlgID
	orderShort.Seq = 0
	orderShort.Vendor = orderPlan.Vendor
	orderShort.Cd = orderPlan.Cd
	orderShort.Tagrp = orderPlan.Tagrp
	orderShort.Sortg1 = orderPlan.Sortg1
	orderShort.Sortg2 = orderPlan.Sortg2
	orderShort.Sortg3 = orderPlan.Sortg3
	orderShort.PlanType = orderPlan.PlanType
	orderShort.OrderGroup = orderPlan.OrderGroup
	orderShort.Pono = orderPlan.Pono
	orderShort.RecId = "SI00" //orderPlan.RecId
	orderShort.Biac = orderPlan.Biac
	orderShort.EtdTap = frm.OrderEtd
	orderShort.PartNo = orderPlan.PartNo
	orderShort.PartName = orderPlan.PartName
	orderShort.SampFlg = orderPlan.SampFlg
	orderShort.Orderorgi = orderPlan.Orderorgi
	orderShort.Orderround = orderPlan.Orderround
	orderShort.FirmFlg = orderPlan.FirmFlg
	orderShort.ShippedFlg = orderPlan.ShippedFlg
	orderShort.ShippedQty = orderPlan.ShippedQty
	orderShort.Ordermonth = orderPlan.Ordermonth
	orderShort.BalQty = orderPlan.Bistdp //orderPlan.BalQty
	orderShort.Bidrfl = orderPlan.Bidrfl
	orderShort.DeleteFlg = orderPlan.DeleteFlg
	orderShort.Reasoncd = "H" //orderPlan.Reasoncd
	orderShort.Upddte = time.Now()
	orderShort.Updtime = time.Now()
	orderShort.CarrierCode = orderPlan.CarrierCode
	orderShort.Bioabt = orderPlan.Bioabt
	orderShort.Bicomd = orderPlan.Bicomd
	orderShort.Bistdp = orderPlan.Bistdp
	orderShort.Binewt = orderPlan.Binewt
	orderShort.Bigrwt = orderPlan.Bigrwt
	orderShort.Bishpc = orderPlan.Bishpc
	orderShort.Biivpx = orderPlan.Biivpx
	orderShort.Bisafn = orderPlan.Bisafn
	orderShort.Biwidt = orderPlan.Biwidt
	orderShort.Bihigh = orderPlan.Bihigh
	orderShort.Bileng = orderPlan.Bileng
	orderShort.LotNo = orderPlan.LotNo
	orderShort.Minimum = orderPlan.Minimum
	orderShort.Maximum = orderPlan.Maximum
	orderShort.Picshelfbin = orderPlan.Picshelfbin
	orderShort.Stkshelfbin = orderPlan.Stkshelfbin
	orderShort.Ovsshelfbin = orderPlan.Ovsshelfbin
	orderShort.PicshelfbasicQty = orderPlan.PicshelfbasicQty
	orderShort.OuterPcs = orderPlan.OuterPcs
	orderShort.AllocateQty = orderPlan.AllocateQty
	orderShort.Description = "Order Short"
	orderShort.IsReviseError = false
	orderShort.IsGenerate = false
	orderShort.ByManually = true
	orderShort.IsSync = false
	orderShort.IsActive = true

	// ลบข้อมูล Order Detail
	var ctnPallet int64
	db.Select("id").Where("order_detail_id=?", &frm.OrderDetailID).Find(&models.PalletDetail{}).Count(&ctnPallet)
	if err := db.Model(&models.OrderDetail{}).Where("id=?", &frm.OrderDetailID).Updates(&models.OrderDetail{
		OrderCtn:      ctnPallet,
		TotalOnPallet: ctnPallet,
	}).Error; err != nil {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// ลบข้อมูล Shipping Label
	if err := db.Where("id=?", &frm.OrderShippingID).Delete(&models.PalletDetail{}).Error; err != nil {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	if err := db.Create(&orderShort).Error; err != nil {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create System Log
	userID := services.GetUserID(c)
	var history models.SyncLogger
	history.Title = "Order Short"
	history.Description = fmt.Sprintf("%s ทำการตัด Short PO %s  Part %s Ctn %d ไปออกวันที่ %s", userID.UserName, orderPlan.Pono, orderPlan.PartNo, int64(float64(frm.OrderCtn)*orderPlan.Bistdp), frm.OrderEtd.Format("2006-01-02"))
	history.IsSuccess = true
	db.Create(&history)

	r.Message = services.MessageUpdateData(&frm.OrderPlanID)
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
