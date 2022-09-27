package services

import (
	"fmt"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
)

func CreateOrder() {
	// Group Order to create order ent
	db := configs.Store
	// dt := time.Now()
	etd := (time.Now()).Format("2006-01-02")
	fmt.Printf("Create Order: %v\n", etd)
	var ord []models.OrderPlan
	db.Model(&models.OrderPlan{}).
		Order("etd_tap").
		Select("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,'sample_flg',carrier_code").
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", "INJ").
		// Where("etd_tap >=?", etd).
		Group("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,'sample_flg',carrier_code").
		Find(&ord)

	var orderTitle models.OrderTitle
	db.Where("value=?", "000").First(&orderTitle)

	x := 0
	for x < len(ord) {
		etd := ord[x].EtdTap.Format("20060102")
		var ship models.Shipment
		db.First(&ship, "id=?", ord[x].ShipmentID)
		/// Generate ZoneCode
		var od []models.Order
		db.Where("etd_date=?", ord[x].EtdTap).Find(&od)
		sum := len(od) + 1
		keyCode := fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, sum)
		var sumOrder models.Order
		db.First(&sumOrder, "zone_code=?", keyCode)
		for !(len(sumOrder.ID) == 0) {
			keyCode = fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, sum)
		}

		// Check LoadingArea
		prefixOrder := "-"
		if ord[x].OrderGroup[:1] == "@" {
			prefixOrder = "@"
		}
		var loadingData models.OrderLoadingArea
		db.Select("prefix,loading_area,privilege").Where("order_zone_id=?", ord[x].OrderZoneID).Where("prefix=?", prefixOrder).First(&loadingData)

		var factoryEnt models.Factory
		db.Where("title=?", ord[x].Vendor).First(&factoryEnt)
		var affcodeData models.Affcode
		db.Where("title=?", ord[x].Biac).First(&affcodeData)

		// Get LastInvoiceNo
		invSeq := models.LastInvoice{
			FactoryID:   &factoryEnt.ID,
			AffcodeID:   &affcodeData.ID,
			LastRunning: 1,
		}
		db.FirstOrCreate(&invSeq, &models.LastInvoice{
			FactoryID: &factoryEnt.ID,
			AffcodeID: &affcodeData.ID,
		})

		order := models.Order{
			ConsigneeID:  ord[x].ConsigneeID,
			ShipmentID:   ord[x].ShipmentID,
			EtdDate:      &ord[x].EtdTap,
			PcID:         ord[x].PcID,
			CommercialID: ord[x].CommercialID,
			OrderTitleID: &orderTitle.ID,
			Bioat:        ord[x].Bioabt,
			ZoneCode:     keyCode,
			LoadingArea:  loadingData.LoadingArea,
			Privilege:    loadingData.Privilege,
			ShipForm:     ord[x].Bishpc,      // bishpc,
			ShipTo:       ord[x].Bisafn,      // bisafn,
			SampleFlg:    ord[x].SampleFlg,   // sample_flg,
			CarrierCode:  ord[x].CarrierCode, // carriercode
			RunningSeq:   (invSeq.LastRunning + 1),
		}

		err := db.FirstOrCreate(&order, &models.Order{
			ConsigneeID:  ord[x].ConsigneeID,
			ShipmentID:   ord[x].ShipmentID,
			EtdDate:      &ord[x].EtdTap,
			PcID:         ord[x].PcID,
			CommercialID: ord[x].CommercialID,
			Bioat:        ord[x].Bioabt,
			IsInvoice:    false,
		}).Error

		if err != nil {
			// Create log if Create Order is Error!
			sysLogger := models.SyncLogger{
				Title:       fmt.Sprintf("creating order %v", ord[x].ConsigneeID),
				Description: fmt.Sprintf("Error creating order: %v", err),
			}
			db.Create(&sysLogger)
			// panic(err)
		}

		// update lastinvoice no
		invSeq.LastRunning += 1
		db.Save(&invSeq)

		if order.ID != "" {
			// Create Order Detail
			var orderPlan []models.OrderPlan
			db.Order("etd_tap,created_at,seq").
				Where("is_revise_error=?", false).
				Where("is_generate=?", false).
				Where("order_zone_id=?", ord[x].OrderZoneID).
				Where("consignee_id=?", ord[x].ConsigneeID).
				Where("shipment_id=?", ord[x].ShipmentID).
				Where("etd_tap=?", ord[x].EtdTap).
				Where("pc_id=?", ord[x].PcID).
				Where("commercial_id=?", ord[x].CommercialID).
				Where("bioabt=?", ord[x].Bioabt).
				Where("order_group=?", ord[x].OrderGroup).
				Find(&orderPlan)
			j := 0
			for j < len(orderPlan) {
				r := orderPlan[j]
				ctn := 0
				if r.BalQty > 0 {
					ctn = int(r.BalQty) / int(r.Bistdp)
				}
				ordDetail := models.OrderDetail{
					OrderID:       &order.ID,
					Pono:          &r.Pono,
					LedgerID:      r.LedgerID,
					OrderPlanID:   &r.ID,
					OrderCtn:      int64(ctn),
					TotalOnPallet: 0,
				}

				db.FirstOrCreate(&ordDetail, &models.OrderDetail{
					OrderID:  &order.ID,
					Pono:     &r.Pono,
					LedgerID: r.LedgerID,
				})

				// Confirm Data After Create
				ordDetail.OrderPlanID = &r.ID
				ordDetail.OrderCtn = int64(ctn)
				db.Save(&ordDetail)

				// Update Order Plan Set Status Generated
				ordPlan := &r
				ordPlan.IsGenerate = true
				db.Save(&ordPlan)
				j++
			}
		}
		x++
	}
}
