package controllers

import (
	"fmt"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/abe27/api/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllShippingLabel(c *fiber.Ctx) error {
	var r models.Response
	orderID := c.Query("order_id")
	if orderID == "" {
		r.Message = services.MessageRequireField("Order ID")
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	var palletData []models.Pallet
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
		Preload("PalletType").
		Preload("PalletDetail.OrderDetail.OrderPlan.FileEdi.Factory").
		Preload("PalletDetail.OrderDetail.OrderPlan.FileEdi.Mailbox.Area").
		Preload("PalletDetail.OrderDetail.OrderPlan.FileEdi.FileType").
		Where("order_id=?", &orderID).Find(&palletData).Error
	if err != nil {
		r.Message = services.MessageSystemErrorWith(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	r.Message = services.MessageShowDataByID(&orderID)
	r.Data = &palletData
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateShippingLabel(c *fiber.Ctx) error {
	var r models.Response
	var Shipping models.PostShippingLabel
	err := c.BodyParser(&Shipping)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(&r)
	}

	db := configs.Store
	var orderDetail models.OrderDetail
	err = db.Preload("OrderPlan.Consignee.Factory").First(&orderDetail, &Shipping.PartID).Error
	if err != nil {
		r.Message = services.MessageSystemErrorWith(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	// Create ShippingLabel
	etd := (time.Now()).Format("2006-01-02")
	for i := 0; i < Shipping.Seq; i++ {
		// Get LastFticket
		var lastNo int64
		db.Select("last_running").Where("on_year=?", etd[0:4]).Find(&models.LastFticket{}).Scan(&lastNo)
		// Count Ctn Set Pallet
		var countSetPallet int64
		db.Model(&models.PalletDetail{}).Where("order_detail_id=?", &Shipping.PartID).Count(&countSetPallet)
		// fmt.Printf("%d\n", countSetPallet)
		if countSetPallet < orderDetail.OrderCtn {
			// fmt.Printf("%d\n", lastNo)
			plData := models.PalletDetail{
				PalletID:      &Shipping.PalletID,
				OrderDetailID: &Shipping.PartID,
				SeqNo:         (lastNo + 1),
				IsPrintLabel:  false,
				IsActive:      true,
			}
			db.Save(&plData)
			err = db.Model(&models.OrderDetail{}).Where("id=?", &Shipping.PartID).Update("total_on_pallet", (countSetPallet + 1)).Error
			if err != nil {
				r.Message = services.MessageSystemErrorWith(err.Error())
				return c.Status(fiber.StatusInternalServerError).JSON(&r)
			}
		}

		labelNo := fmt.Sprintf("%s%s%08d", orderDetail.OrderPlan.Consignee.Factory.LabelPrefix, etd[3:4], (lastNo + 1))
		fmt.Printf("%s ==> %d\n", labelNo, i)
		db.Model(&models.LastFticket{}).Where("on_year=?", etd[0:4]).Update("last_running", (lastNo + 1))
	}

	r.Message = services.MessageCreatedData(&Shipping.PalletID)
	r.Data = &orderDetail
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func UpdateShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func DeleteShippingLabelByID(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreatePrintLabel(c *fiber.Ctx) error {
	var r models.Response
	var frm []models.PrintShippingLabel
	err := c.BodyParser(&frm)
	if err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	db := configs.Store
	var data []models.PrintShippingLabel
	for _, p := range frm {
		obj := models.PrintShippingLabel{
			InvoiceNo:    p.InvoiceNo,
			OrderNo:      p.OrderNo,
			PartNo:       p.PartNo,
			CustCode:     p.CustCode,
			CustName:     p.CustName,
			PalletNo:     p.PalletNo,
			PrintDate:    p.PrintDate,
			Qty:          p.Qty,
			QrCode:       fmt.Sprintf("06P%s;17Q%d;30T%s;32T%s;", p.PartNo, p.Qty, p.OrderNo, p.BarCode),
			BarCode:      p.BarCode,
			LabelBarCode: fmt.Sprintf("*%s*", p.BarCode),
			IsPrint:      p.IsPrint,
		}
		err = db.FirstOrCreate(&obj, &models.PrintShippingLabel{BarCode: p.BarCode}).Error
		if err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}

		obj.Qty = p.Qty
		obj.QrCode = fmt.Sprintf("06P%s;17Q%d;30T%s;32T%s;", p.PartNo, p.Qty, p.OrderNo, p.BarCode)
		obj.IsPrint = 0
		err = db.Save(&obj).Error
		if err != nil {
			r.Message = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(&r)
		}
		data = append(data, obj)
	}

	txt := fmt.Sprintf("%d รายการ", len(data))
	if len(data) == 1 {
		txt = data[0].BarCode
	}
	r.Message = services.MessagePrintShippingLabel(txt)
	r.Data = &data
	return c.Status(fiber.StatusOK).JSON(&r)
}
