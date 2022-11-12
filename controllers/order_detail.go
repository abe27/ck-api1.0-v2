package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrderDetail(c *fiber.Ctx) error {
	var r models.Response
	id := c.Query("order_id")
	if id == "" {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var data []models.OrderDetail
	err := configs.Store.
		Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("Ledger.Whs").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.PartType").
		Preload("Ledger.Unit").
		Preload("OrderPlan.FileEdi.Factory").
		Preload("OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderPlan.FileEdi.FileType").
		Preload("OrderPlan.Whs").
		Preload("OrderPlan.ReviseOrder").
		Preload("OrderPlan.OrderZone").
		Preload("OrderPlan.SampleFlg").
		Where("order_id=?", &id).Find(&data).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateOrderDetail(c *fiber.Ctx) error {
	var r models.Response
	id := c.Query("order_id")
	if id == "" {
		r.Message = services.MessageRequireField("order_id")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	db := configs.Store
	// Get OrderDetail
	var orderDetail models.OrderDetail
	err := db.
		Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("Ledger.Whs").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.PartType").
		Preload("Ledger.Unit").
		Preload("OrderPlan.FileEdi.Factory").
		Preload("OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderPlan.FileEdi.FileType").
		Preload("OrderPlan.Whs").
		Preload("OrderPlan.ReviseOrder").
		Preload("OrderPlan.OrderZone").
		Preload("OrderPlan.SampleFlg").
		Where("order_id=?", &id).First(&orderDetail).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var frm models.AddOrderDetailForm
	err = c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	// to Upper
	frm.Pono = strings.ToUpper(frm.Pono)
	// Find part no
	var part models.Part
	err = db.Select("id,title,description").Where("slug=?", strings.ToUpper(frm.PartNo)).First(&part).Error
	if err != nil {
		r.Message = services.MessageNotFoundData(&frm.PartNo)
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var reviseData models.ReviseOrder
	err = db.First(&reviseData, "title=?", "0").Error
	if err != nil {
		r.Message = services.MessageSystemError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Fetch Ledger
	ledger := models.Ledger{
		WhsID:       &orderDetail.Order.Consignee.Whs.ID,
		FactoryID:   &orderDetail.Order.Consignee.Factory.ID,
		PartID:      &part.ID,
		PartTypeID:  &orderDetail.Ledger.PartType.ID,
		UnitID:      &orderDetail.Ledger.Unit.ID,
		DimWidth:    0,
		DimLength:   0,
		DimHeight:   0,
		GrossWeight: 0,
		NetWeight:   0,
		Qty:         0,
		Ctn:         0,
	}

	db.FirstOrCreate(&ledger, &models.Ledger{
		WhsID:     &orderDetail.Order.Consignee.Whs.ID,
		FactoryID: &orderDetail.Order.Consignee.Factory.ID,
		PartID:    &part.ID,
	})

	p := orderDetail.OrderPlan
	var orderPlan models.OrderPlan
	orderPlan.WhsID = p.WhsID
	orderPlan.OrderZoneID = p.OrderZoneID
	orderPlan.ConsigneeID = p.ConsigneeID
	orderPlan.ReviseOrderID = p.ReviseOrderID
	orderPlan.LedgerID = p.LedgerID
	orderPlan.PcID = p.PcID
	orderPlan.CommercialID = p.CommercialID
	orderPlan.OrderTypeID = p.OrderTypeID
	orderPlan.ShipmentID = p.ShipmentID
	orderPlan.SampleFlgID = p.SampleFlgID
	orderPlan.Vendor = p.Vendor
	orderPlan.Cd = p.Cd
	orderPlan.Tagrp = p.Tagrp
	orderPlan.Sortg1 = p.Sortg1
	orderPlan.Sortg2 = p.Sortg2
	orderPlan.Sortg3 = p.Sortg3
	orderPlan.PlanType = p.PlanType
	orderPlan.OrderGroup = p.OrderGroup
	orderPlan.Pono = frm.Pono
	orderPlan.Biac = p.Biac
	orderPlan.EtdTap = p.EtdTap
	orderPlan.PartNo = p.PartNo
	orderPlan.PartName = part.Description
	orderPlan.SampFlg = p.SampFlg
	orderPlan.Orderorgi = float64(frm.OrderBalQty)
	orderPlan.Orderround = float64(frm.OrderBalQty)
	orderPlan.FirmFlg = p.FirmFlg
	orderPlan.ShippedFlg = p.ShippedFlg
	orderPlan.ShippedQty = 0
	orderPlan.Ordermonth = time.Now()
	orderPlan.BalQty = float64(frm.OrderBalQty)
	orderPlan.Bidrfl = p.Bidrfl
	orderPlan.DeleteFlg = p.DeleteFlg
	orderPlan.Reasoncd = reviseData.Title
	orderPlan.Upddte = time.Now()
	orderPlan.Updtime = time.Now()
	orderPlan.CarrierCode = p.CarrierCode
	orderPlan.Bioabt = p.Bioabt
	orderPlan.Bicomd = p.Bicomd
	orderPlan.Bistdp = float64(frm.OrderBalQty) / float64(frm.OrderCtn)
	orderPlan.Binewt = ledger.NetWeight
	orderPlan.Bigrwt = ledger.GrossWeight
	orderPlan.Bishpc = p.Bishpc
	orderPlan.Biivpx = p.Biivpx
	orderPlan.Bisafn = p.Bisafn
	orderPlan.Biwidt = ledger.DimWidth
	orderPlan.Bihigh = ledger.DimHeight
	orderPlan.Bileng = ledger.DimLength
	orderPlan.LotNo = ""
	orderPlan.Minimum = p.Minimum
	orderPlan.Maximum = p.Maximum
	orderPlan.Picshelfbin = p.Picshelfbin
	orderPlan.Stkshelfbin = p.Stkshelfbin
	orderPlan.Ovsshelfbin = p.Ovsshelfbin
	orderPlan.PicshelfbasicQty = p.PicshelfbasicQty
	orderPlan.OuterPcs = 0
	orderPlan.AllocateQty = 0
	orderPlan.Description = "-"
	orderPlan.IsReviseError = true
	orderPlan.IsGenerate = false
	orderPlan.ByManually = false
	orderPlan.IsSync = false
	orderPlan.IsActive = true
	err = db.Create(&orderPlan).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusCreated).JSON(&r)
	}

	// Create History OrderPlan
	userID := services.GetUserID(c)
	var HistoryOrderPlan models.HistoryOrderPlan
	HistoryOrderPlan.UserID = userID.ID
	HistoryOrderPlan.WhsID = orderPlan.WhsID
	HistoryOrderPlan.OrderZoneID = orderPlan.OrderZoneID
	HistoryOrderPlan.ConsigneeID = orderPlan.ConsigneeID
	HistoryOrderPlan.ReviseOrderID = orderPlan.ReviseOrderID
	HistoryOrderPlan.LedgerID = orderPlan.LedgerID
	HistoryOrderPlan.PcID = orderPlan.PcID
	HistoryOrderPlan.CommercialID = orderPlan.CommercialID
	HistoryOrderPlan.OrderTypeID = orderPlan.OrderTypeID
	HistoryOrderPlan.ShipmentID = orderPlan.ShipmentID
	HistoryOrderPlan.SampleFlgID = orderPlan.SampleFlgID
	HistoryOrderPlan.Vendor = orderPlan.Vendor
	HistoryOrderPlan.Cd = orderPlan.Cd
	HistoryOrderPlan.Tagrp = orderPlan.Tagrp
	HistoryOrderPlan.Sortg1 = orderPlan.Sortg1
	HistoryOrderPlan.Sortg2 = orderPlan.Sortg2
	HistoryOrderPlan.Sortg3 = orderPlan.Sortg3
	HistoryOrderPlan.PlanType = orderPlan.PlanType
	HistoryOrderPlan.OrderGroup = orderPlan.OrderGroup
	HistoryOrderPlan.Pono = orderPlan.Pono
	HistoryOrderPlan.RecId = orderPlan.RecId
	HistoryOrderPlan.Biac = orderPlan.Biac
	HistoryOrderPlan.EtdTap = orderPlan.EtdTap
	HistoryOrderPlan.PartNo = orderPlan.PartNo
	HistoryOrderPlan.PartName = orderPlan.PartName
	HistoryOrderPlan.SampFlg = orderPlan.SampFlg
	HistoryOrderPlan.Orderorgi = orderPlan.Orderorgi
	HistoryOrderPlan.Orderround = orderPlan.Orderround
	HistoryOrderPlan.FirmFlg = orderPlan.FirmFlg
	HistoryOrderPlan.ShippedFlg = orderPlan.ShippedFlg
	HistoryOrderPlan.ShippedQty = orderPlan.ShippedQty
	HistoryOrderPlan.Ordermonth = orderPlan.Ordermonth
	HistoryOrderPlan.BalQty = orderPlan.BalQty
	HistoryOrderPlan.Bidrfl = orderPlan.Bidrfl
	HistoryOrderPlan.DeleteFlg = orderPlan.DeleteFlg
	HistoryOrderPlan.Reasoncd = orderPlan.Reasoncd
	HistoryOrderPlan.Upddte = orderPlan.Upddte
	HistoryOrderPlan.Updtime = orderPlan.Updtime
	HistoryOrderPlan.CarrierCode = orderPlan.CarrierCode
	HistoryOrderPlan.Bioabt = orderPlan.Bioabt
	HistoryOrderPlan.Bicomd = orderPlan.Bicomd
	HistoryOrderPlan.Bistdp = orderPlan.Bistdp
	HistoryOrderPlan.Binewt = orderPlan.Binewt
	HistoryOrderPlan.Bigrwt = orderPlan.Bigrwt
	HistoryOrderPlan.Bishpc = orderPlan.Bishpc
	HistoryOrderPlan.Biivpx = orderPlan.Biivpx
	HistoryOrderPlan.Bisafn = orderPlan.Bisafn
	HistoryOrderPlan.Biwidt = orderPlan.Biwidt
	HistoryOrderPlan.Bihigh = orderPlan.Bihigh
	HistoryOrderPlan.Bileng = orderPlan.Bileng
	HistoryOrderPlan.LotNo = orderPlan.LotNo
	HistoryOrderPlan.Minimum = orderPlan.Minimum
	HistoryOrderPlan.Maximum = orderPlan.Maximum
	HistoryOrderPlan.Picshelfbin = orderPlan.Picshelfbin
	HistoryOrderPlan.Stkshelfbin = orderPlan.Stkshelfbin
	HistoryOrderPlan.Ovsshelfbin = orderPlan.Ovsshelfbin
	HistoryOrderPlan.PicshelfbasicQty = orderPlan.PicshelfbasicQty
	HistoryOrderPlan.OuterPcs = orderPlan.OuterPcs
	HistoryOrderPlan.AllocateQty = orderPlan.AllocateQty
	HistoryOrderPlan.Description = orderPlan.Description
	HistoryOrderPlan.IsSync = false

	err = db.Create(&HistoryOrderPlan).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create Order Detail
	ordDetail := models.OrderDetail{
		OrderID:       &id,
		Pono:          &frm.Pono,
		LedgerID:      &ledger.ID,
		OrderPlanID:   &orderPlan.ID,
		OrderCtn:      frm.OrderCtn,
		TotalOnPallet: 0,
		IsMatched:     false,
		IsSync:        false,
		IsActive:      true,
	}

	var orderDetailID string
	db.Select("id").Where("order_id=?", &id).Where("pono=?", &frm.Pono).Where("ledger_id=?", &ledger.ID).First(&models.OrderDetail{}).Scan(&orderDetailID)
	if orderDetailID != "" {
		r.Message = services.MessageDuplicateData(&orderDetailID)
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	err = db.Create(&ordDetail).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create System Log
	var history models.SyncLogger
	history.Title = "Add Order"
	history.Description = fmt.Sprintf("%s เพิ่มรายการ %s %s %d", userID.UserName, orderPlan.Pono, orderPlan.PartNo, int64(float64(frm.OrderCtn)*orderPlan.Bistdp))
	history.IsSuccess = true
	db.Create(&history)

	r.Message = services.MessageCreatedData(&id)
	r.Data = &ledger
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func ShowOrderDetailByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	if id == "" {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var data models.OrderDetail
	err := configs.Store.
		Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("Ledger.Whs").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.PartType").
		Preload("Ledger.Unit").
		Preload("OrderPlan.FileEdi.Factory").
		Preload("OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderPlan.FileEdi.FileType").
		Preload("OrderPlan.Whs").
		Preload("OrderPlan.ReviseOrder").
		Preload("OrderPlan.OrderZone").
		Preload("OrderPlan.SampleFlg").
		Where("id=?", &id).First(&data).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = services.MessageShowDataByID(&id)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateOrderDetailByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.OrderDetailForm
	// Get UserID
	userID := services.GetUserID(c)
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	id := c.Params("id")
	if id == "" {
		r.Message = services.MessageRequireField("id")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store
	var data models.OrderDetail
	err = db.Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("Ledger.Whs").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.PartType").
		Preload("Ledger.Unit").
		Preload("OrderPlan.FileEdi.Factory").
		Preload("OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderPlan.FileEdi.FileType").
		Preload("OrderPlan.Whs").
		Preload("OrderPlan.ReviseOrder").
		Preload("OrderPlan.OrderZone").
		Preload("OrderPlan.SampleFlg").
		First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// orderPlan.BalQty = float64(frm.OrderCtn) * orderPlan.Bistdp
	// orderPlan.Reasoncd = frm.ReviseID
	// orderPlan.ReviseOrderID = &reviseData.ID
	// err = db.Save(&orderPlan).Error
	// if err != nil {
	// 	r.Message = services.MessageSystemError
	// 	r.Data = err.Error()
	// 	return c.Status(fiber.StatusInternalServerError).JSON(&r)
	// }

	/// Save OrderDetail
	var orderDetail models.OrderDetail
	err = db.First(&orderDetail, &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	orderDetail.OrderCtn = frm.OrderCtn
	orderDetail.TotalOnPallet = frm.TotalOnPallet
	err = db.Save(&orderDetail).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create New OrderPlan
	var orderPlan models.OrderPlan
	err = db.First(&orderPlan, data.OrderPlanID).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	/// Save OrderPlan
	var reviseData models.ReviseOrder
	err = db.First(&reviseData, "title=?", frm.ReviseID).Error
	if err != nil {
		r.Message = services.MessageSystemError
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	/// Create Log
	var HistoryOrderPlan models.HistoryOrderPlan
	HistoryOrderPlan.UserID = userID.ID
	HistoryOrderPlan.WhsID = orderPlan.WhsID
	HistoryOrderPlan.OrderZoneID = orderPlan.OrderZoneID
	HistoryOrderPlan.ConsigneeID = orderPlan.ConsigneeID
	HistoryOrderPlan.ReviseOrderID = orderPlan.ReviseOrderID
	HistoryOrderPlan.LedgerID = orderPlan.LedgerID
	HistoryOrderPlan.PcID = orderPlan.PcID
	HistoryOrderPlan.CommercialID = orderPlan.CommercialID
	HistoryOrderPlan.OrderTypeID = orderPlan.OrderTypeID
	HistoryOrderPlan.ShipmentID = orderPlan.ShipmentID
	HistoryOrderPlan.SampleFlgID = orderPlan.SampleFlgID
	HistoryOrderPlan.Vendor = orderPlan.Vendor
	HistoryOrderPlan.Cd = orderPlan.Cd
	HistoryOrderPlan.Tagrp = orderPlan.Tagrp
	HistoryOrderPlan.Sortg1 = orderPlan.Sortg1
	HistoryOrderPlan.Sortg2 = orderPlan.Sortg2
	HistoryOrderPlan.Sortg3 = orderPlan.Sortg3
	HistoryOrderPlan.PlanType = orderPlan.PlanType
	HistoryOrderPlan.OrderGroup = orderPlan.OrderGroup
	HistoryOrderPlan.Pono = orderPlan.Pono
	HistoryOrderPlan.RecId = orderPlan.RecId
	HistoryOrderPlan.Biac = orderPlan.Biac
	HistoryOrderPlan.EtdTap = orderPlan.EtdTap
	HistoryOrderPlan.PartNo = orderPlan.PartNo
	HistoryOrderPlan.PartName = orderPlan.PartName
	HistoryOrderPlan.SampFlg = orderPlan.SampFlg
	HistoryOrderPlan.Orderorgi = orderPlan.Orderorgi
	HistoryOrderPlan.Orderround = orderPlan.Orderround
	HistoryOrderPlan.FirmFlg = orderPlan.FirmFlg
	HistoryOrderPlan.ShippedFlg = orderPlan.ShippedFlg
	HistoryOrderPlan.ShippedQty = orderPlan.ShippedQty
	HistoryOrderPlan.Ordermonth = orderPlan.Ordermonth
	HistoryOrderPlan.BalQty = orderPlan.BalQty
	HistoryOrderPlan.Bidrfl = orderPlan.Bidrfl
	HistoryOrderPlan.DeleteFlg = orderPlan.DeleteFlg
	HistoryOrderPlan.Reasoncd = orderPlan.Reasoncd
	HistoryOrderPlan.Upddte = orderPlan.Upddte
	HistoryOrderPlan.Updtime = orderPlan.Updtime
	HistoryOrderPlan.CarrierCode = orderPlan.CarrierCode
	HistoryOrderPlan.Bioabt = orderPlan.Bioabt
	HistoryOrderPlan.Bicomd = orderPlan.Bicomd
	HistoryOrderPlan.Bistdp = orderPlan.Bistdp
	HistoryOrderPlan.Binewt = orderPlan.Binewt
	HistoryOrderPlan.Bigrwt = orderPlan.Bigrwt
	HistoryOrderPlan.Bishpc = orderPlan.Bishpc
	HistoryOrderPlan.Biivpx = orderPlan.Biivpx
	HistoryOrderPlan.Bisafn = orderPlan.Bisafn
	HistoryOrderPlan.Biwidt = orderPlan.Biwidt
	HistoryOrderPlan.Bihigh = orderPlan.Bihigh
	HistoryOrderPlan.Bileng = orderPlan.Bileng
	HistoryOrderPlan.LotNo = orderPlan.LotNo
	HistoryOrderPlan.Minimum = orderPlan.Minimum
	HistoryOrderPlan.Maximum = orderPlan.Maximum
	HistoryOrderPlan.Picshelfbin = orderPlan.Picshelfbin
	HistoryOrderPlan.Stkshelfbin = orderPlan.Stkshelfbin
	HistoryOrderPlan.Ovsshelfbin = orderPlan.Ovsshelfbin
	HistoryOrderPlan.PicshelfbasicQty = orderPlan.PicshelfbasicQty
	HistoryOrderPlan.OuterPcs = orderPlan.OuterPcs
	HistoryOrderPlan.AllocateQty = orderPlan.AllocateQty
	HistoryOrderPlan.Description = orderPlan.Description
	HistoryOrderPlan.IsSync = false

	err = db.Create(&HistoryOrderPlan).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Old Qty
	olderQty := orderPlan.BalQty
	// Save OderPlan
	orderPlan.BalQty = float64(frm.OrderCtn) * orderPlan.Bistdp
	orderPlan.Reasoncd = frm.ReviseID
	orderPlan.ReviseOrderID = &reviseData.ID
	orderPlan.ByManually = true
	orderPlan.IsSync = false
	err = db.Save(&orderPlan).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Data = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create System Log
	var history models.SyncLogger
	history.Title = "Revise Order"
	history.Description = fmt.Sprintf("%s แก้ไข %s %s ข้อมูล qty จาก %d เป็น %d", userID.UserName, orderPlan.Pono, orderPlan.PartNo, int64(olderQty), int64(float64(frm.OrderCtn)*orderPlan.Bistdp))
	history.IsSuccess = true
	db.Create(&history)
	r.Message = services.MessageUpdateDataByID(&orderDetail.ID)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteOrderDetailByID(c *fiber.Ctx) error {
	var r models.Response
	id := c.Params("id")
	if id == "" {
		r.Message = services.MessageInputValidationError
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	db := configs.Store
	var data models.OrderDetail
	err := db.Preload("Order.Consignee.Whs").
		Preload("Order.Consignee.Factory").
		Preload("Order.Consignee.Affcode").
		Preload("Order.Consignee.Customer").
		Preload("Order.Consignee.CustomerAddress").
		Preload("Order.Shipment").
		Preload("Order.Pc").
		Preload("Order.Commercial").
		Preload("Order.SampleFlg").
		Preload("Order.OrderTitle").
		Preload("Ledger.Whs").
		Preload("Ledger.Factory").
		Preload("Ledger.Part").
		Preload("Ledger.PartType").
		Preload("Ledger.Unit").
		Preload("OrderPlan.FileEdi.Factory").
		Preload("OrderPlan.FileEdi.Mailbox.Area").
		Preload("OrderPlan.FileEdi.FileType").
		Preload("OrderPlan.Whs").
		Preload("OrderPlan.ReviseOrder").
		Preload("OrderPlan.OrderZone").
		Preload("OrderPlan.SampleFlg").First(&data, &id).Error
	if err != nil {
		r.Message = services.MessageNotFound(id)
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	err = configs.Store.Delete(&models.OrderDetail{}, &id).Error
	if err != nil {
		r.Message = services.MessageSystemError
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create System Log
	userID := services.GetUserID(c)
	var history models.SyncLogger
	history.Title = "Delete Order"
	history.Description = fmt.Sprintf("%s ลบรายการ %s %s %d", userID.UserName, *data.Pono, data.Ledger.Part.Title, data.OrderCtn)
	history.IsSuccess = true
	db.Create(&history)

	r.Message = services.MessageDeleteData(&id)
	return c.Status(fiber.StatusOK).JSON(&r)
}
