package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateOrder(factory, end_etd string) {
	var orderTitle models.OrderTitle
	if err := configs.Store.Select("id,title").Where("title=?", "000").First(&orderTitle).Error; err != nil {
		panic(err)
	}
	CreateOrderWithOutRevise(factory, end_etd, &orderTitle)
	/// After Generate Order Get Tap Data
	GenerateImportInvoiceTap()
	// After Get Tap
	CreateOrderWithRevise(factory, end_etd, &orderTitle)
	/// After Generate Order Get Tap Data
	ClearOrder()
	GenerateImportInvoiceTap()
}

func CreateOrderWithOutRevise(factory, endDate string, orderTitle *models.OrderTitle) {
	// Group Order to create order ent
	db := configs.Store
	var ord []models.OrderPlan
	if err := db.
		Order("etd_tap").
		Select("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,sample_flg_id,carrier_code").
		Where("length(reasoncd) = ?", 0).
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", &factory).
		Where("etd_tap <= ?", endDate).
		Group("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,sample_flg_id,carrier_code").
		Find(&ord).Error; err != nil {
		sysLogger := models.SyncLogger{
			Title:       "generate order ent",
			Description: fmt.Sprintf("Error fetch order: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	for _, v := range ord {
		fmt.Printf("Generate Order Etd: %s\n", v.EtdTap)
		etd := v.EtdTap.Format("20060102")
		var ship models.Shipment
		if err := db.Select("id,title").First(&ship, "id=?", v.ShipmentID).Error; err != nil {
			panic(err)
		}

		fmt.Printf("Shipment ID: %s\n", ship.ID)
		/// Generate ZoneCode
		var ordCount int64
		if err := db.Select("id").Where("etd_date=?", v.EtdTap).Find(&models.Order{}).Count(&ordCount).Error; err != nil {
			panic(err)
		}

		fmt.Printf("Count Order: %d\n", ordCount)
		keyCode := fmt.Sprintf("%s%s", etd[2:], ship.Title)
		var countSeq int64
		db.Raw("SELECT count(id) + 1 FROM tbt_orders WHERE zone_code like ?", keyCode+"%").Scan(&countSeq)
		keyCode = fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, countSeq)
		// fmt.Printf("ZoneCode: %s Seq: %d\n", keyCode, countSeq)

		// Check LoadingArea
		// fmt.Printf("Check LoadingArea %s\n", v.OrderGroup[:1])
		prefixOrder := "-"
		if v.OrderGroup[:1] == "@" {
			prefixOrder = "@"
		}
		var loadingData models.OrderLoadingArea
		db.Select("prefix,loading_area,privilege").Where("order_zone_id=?", v.OrderZoneID).Where("prefix=?", prefixOrder).First(&loadingData)

		var factoryEnt models.Factory
		db.Select("id,title").Where("title=?", v.Vendor).First(&factoryEnt)
		var affcodeData models.Affcode
		db.Select("id,title,description").Where("title=?", v.Biac).First(&affcodeData)

		// Get LastInvoiceNo
		invSeq := models.LastInvoice{
			FactoryID:   &factoryEnt.ID,
			AffcodeID:   &affcodeData.ID,
			OnYear:      ConvertInt((etd)[:4]),
			LastRunning: 0,
		}

		db.FirstOrCreate(&invSeq, &models.LastInvoice{
			FactoryID: &factoryEnt.ID,
			AffcodeID: &affcodeData.ID,
			OnYear:    ConvertInt((etd)[:4]),
		})

		// Fetch Order
		var order models.Order
		order.ConsigneeID = v.ConsigneeID
		order.ShipmentID = v.ShipmentID
		order.EtdDate = &v.EtdTap
		order.PcID = v.PcID
		order.CommercialID = v.CommercialID
		order.SampleFlgID = v.SampleFlgID
		order.OrderTitleID = &orderTitle.ID
		order.Bioabt = v.Bioabt
		order.ZoneCode = keyCode
		order.LoadingArea = loadingData.LoadingArea
		order.Privilege = loadingData.Privilege
		order.ShipForm = v.Bishpc         // bishpc
		order.ShipTo = v.Bisafn           // bisafn
		order.SampleFlg = v.SampleFlg     // sample_flg
		order.CarrierCode = v.CarrierCode // carriercod
		order.RunningSeq = (invSeq.LastRunning + 1)
		order.IsActive = false
		order.IsSync = true

		if err := db.First(&order, &models.Order{
			ConsigneeID:  v.ConsigneeID,
			ShipmentID:   v.ShipmentID,
			EtdDate:      &v.EtdTap,
			PcID:         v.PcID,
			CommercialID: v.CommercialID,
			SampleFlgID:  v.SampleFlgID,
			Bioabt:       v.Bioabt,
		}).Error; err != nil {
			if err := db.Create(&order).Error; err != nil {
				panic(err)
			}
			// update lastinvoice no
			invSeq.LastRunning += 1
			db.Save(&invSeq)
		}

		fmt.Printf("Order ID: %s\n", order.ID)
		// Fetch Order Plan
		if order.ID != "" {
			var orderPlan []models.OrderPlan
			if err := db.Raw("select * from tbt_order_plans where order_zone_id=? and consignee_id=? and pc_id=? and commercial_id=? and shipment_id=? and order_group=? and etd_tap=? and bioabt=? order by upddte,updtime,seq", v.OrderZoneID, v.ConsigneeID, v.PcID, v.CommercialID, v.ShipmentID, v.OrderGroup, v.EtdTap.Format("2006-01-02"), v.Bioabt).Scan(&orderPlan).Error; err != nil {
				panic(err)
			}

			rnd := 0
			for _, r := range orderPlan {
				ctn := 0
				if r.BalQty > 0 {
					ctn = int(r.BalQty) / int(r.Bistdp)
				}
				fmt.Printf("%d ::: %s %s %d Revise: %s %s\n", rnd, r.Pono, r.PartNo, ctn, r.Reasoncd, r.Updtime)
				var ordDetail models.OrderDetail
				ordDetail.OrderID = &order.ID
				ordDetail.Pono = &r.Pono
				ordDetail.LedgerID = r.LedgerID
				ordDetail.OrderPlanID = &r.ID
				ordDetail.OrderCtn = int64(ctn)
				ordDetail.TotalOnPallet = 0

				db.FirstOrCreate(&ordDetail, &models.OrderDetail{
					OrderID:  &order.ID,
					Pono:     &r.Pono,
					LedgerID: r.LedgerID,
				})

				// Confirm Data After Create
				ordDetail.OrderPlanID = &r.ID
				ordDetail.OrderCtn = int64(ctn)
				ordDetail.IsSync = true
				if err := db.Save(&ordDetail).Error; err == nil {
					// Update Order Plan Set Status Generated
					db.Model(&models.OrderPlan{}).Where("id=?", r.ID).Update("is_generate", true)
				}
				rnd++
			}
		}
	}
}

func CreateOrderWithRevise(factory, endDate string, orderTitle *models.OrderTitle) {
	db := configs.Store
	var ordDetail []models.OrderPlan
	if err := db.Raw("select * from tbt_order_plans where length(reasoncd) > 0 and is_generate=false and is_revise_error=false and vendor=? order by upddte,updtime,seq", factory).Scan(&ordDetail).Error; err != nil {
		sysLogger := models.SyncLogger{
			Title:       "generate order ent revises",
			Description: fmt.Sprintf("%v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}
	/// parse time
	parseEndDate, _ := time.Parse("2006-01-02", endDate)
	for _, ord := range ordDetail {
		// GenerateOrderDetailWithRevise(end_etd, &ord[i], orderTitle)
		if !(ord.EtdTap.After(parseEndDate)) {
			fmt.Printf("ID: %s ETD: %s Revise: %s IsAfter: %v\n", ord.ID, ord.EtdTap.Format("2006-01-02"), ord.Reasoncd[:1], (ord.EtdTap.After(parseEndDate)))
			var order models.Order
			db.Order("created_at desc").Select("id").Last(&order, &models.Order{
				ConsigneeID:  ord.ConsigneeID,
				EtdDate:      &ord.EtdTap,
				ShipmentID:   ord.ShipmentID,
				PcID:         ord.PcID,
				CommercialID: ord.CommercialID,
				SampleFlgID:  ord.SampleFlgID,
				Bioabt:       ord.Bioabt,
			})

			if order.ID == "" {
				etd := ord.EtdTap.Format("20060102")
				var ship models.Shipment
				if err := db.Select("id,title").First(&ship, "id=?", ord.ShipmentID).Error; err != nil {
					panic(err)
				}

				fmt.Printf("Shipment ID: %s\n", ship.ID)
				/// Generate ZoneCode
				var ordCount int64
				if err := db.Select("id").Where("etd_date=?", ord.EtdTap).Find(&models.Order{}).Count(&ordCount).Error; err != nil {
					panic(err)
				}

				fmt.Printf("Count Order: %d\n", ordCount)
				keyCode := fmt.Sprintf("%s%s", etd[2:], ship.Title)
				var countSeq int64
				db.Raw("SELECT count(id) + 1 FROM tbt_orders WHERE zone_code like ?", keyCode+"%").Scan(&countSeq)
				keyCode = fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, countSeq)

				// Check LoadingArea
				// fmt.Printf("Check LoadingArea %s\n", ord[x].OrderGroup[:1])
				prefixOrder := "-"
				if ord.OrderGroup[:1] == "@" {
					prefixOrder = "@"
				}
				var loadingData models.OrderLoadingArea
				db.Select("prefix,loading_area,privilege").Where("order_zone_id=?", ord.OrderZoneID).Where("prefix=?", prefixOrder).First(&loadingData)

				var factoryEnt models.Factory
				db.Select("id,title").Where("title=?", ord.Vendor).First(&factoryEnt)
				var affcodeData models.Affcode
				db.Select("id,title,description").Where("title=?", ord.Biac).First(&affcodeData)

				// Get LastInvoiceNo
				invSeq := models.LastInvoice{
					FactoryID:   &factoryEnt.ID,
					AffcodeID:   &affcodeData.ID,
					OnYear:      ConvertInt((etd)[:4]),
					LastRunning: 0,
				}

				db.FirstOrCreate(&invSeq, &models.LastInvoice{
					FactoryID: &factoryEnt.ID,
					AffcodeID: &affcodeData.ID,
					OnYear:    ConvertInt((etd)[:4]),
				})

				// Fetch Order Entries
				order.ConsigneeID = ord.ConsigneeID
				order.ShipmentID = ord.ShipmentID
				order.EtdDate = &ord.EtdTap
				order.PcID = ord.PcID
				order.CommercialID = ord.CommercialID
				order.SampleFlgID = ord.SampleFlgID
				order.OrderTitleID = &orderTitle.ID
				order.Bioabt = ord.Bioabt
				order.ZoneCode = keyCode
				order.LoadingArea = loadingData.LoadingArea
				order.Privilege = loadingData.Privilege
				order.ShipForm = ord.Bishpc         // bishpc
				order.ShipTo = ord.Bisafn           // bisafn
				order.SampleFlg = ord.SampleFlg     // sample_flg
				order.CarrierCode = ord.CarrierCode // carriercod
				order.RunningSeq = (invSeq.LastRunning + 1)
				order.IsActive = false
				order.IsSync = true

				if err := db.Save(&order).Error; err == nil {
					invSeq.LastRunning = order.RunningSeq
					db.Save(&invSeq)
				}
			}

			if order.ID != "" {
				fmt.Printf("ID: %s\n", order.ID)
				CreateOrderDetail(&order, &ord)
			}
		} else {
			var orderDetail models.OrderDetail
			if err := db.Where("pono=?", &ord.Pono).Where("ledger_id=?", ord.LedgerID).Where("total_on_pallet=?", 0).First(&orderDetail).Error; err == nil {
				fmt.Printf("Order %s found ID: %s  OrderID: %s ETD: %s Revise: %s\n", ord.ID, orderDetail.ID, *orderDetail.OrderID, ord.EtdTap, ord.Reasoncd)
				if err := db.Delete(&models.OrderDetail{}, "id", orderDetail.ID).Error; err == nil {
					var countOrdDetail int64
					db.Select("id").Where("order_id=?", orderDetail.OrderID).Find(&models.OrderDetail{}).Count(&countOrdDetail)
					if countOrdDetail == 0 {
						if err := db.Delete(&models.Order{}, "id", orderDetail.OrderID).Error; err != nil {
							panic(err)
						}
					}
				}
			}
		}
		DeleteOrder(&ord)
	}
}

func GenerateOrderDetailWithRevise(endDate string, ord *models.OrderPlan, orderTitle *models.OrderTitle) {
	db := configs.Store
	parseEndDate, _ := time.Parse("2006-01-02", endDate)
	if !(ord.EtdTap.After(parseEndDate)) {
		etd := ord.EtdTap.Format("20060102")
		var ship models.Shipment
		db.Select("id,title").First(&ship, "id=?", ord.ShipmentID)
		/// Generate ZoneCode
		var ordCount int64
		db.Select("id").Where("etd_date=?", ord.EtdTap).Find(&models.Order{}).Count(&ordCount)
		sum := ordCount + 1
		keyCode := fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, sum)
		var sumOrder models.Order
		db.Select("id").First(&sumOrder, "zone_code=?", keyCode)
		for !(len(sumOrder.ID) == 0) {
			keyCode = fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, sum)
		}

		// Check LoadingArea
		// fmt.Printf("Check LoadingArea %s\n", ord[x].OrderGroup[:1])
		prefixOrder := "-"
		if ord.OrderGroup[:1] == "@" {
			prefixOrder = "@"
		}
		var loadingData models.OrderLoadingArea
		db.Select("prefix,loading_area,privilege").Where("order_zone_id=?", ord.OrderZoneID).Where("prefix=?", prefixOrder).First(&loadingData)

		var factoryEnt models.Factory
		db.Select("id,title").Where("title=?", ord.Vendor).First(&factoryEnt)
		var affcodeData models.Affcode
		db.Select("id,title,description").Where("title=?", ord.Biac).First(&affcodeData)

		// Get LastInvoiceNo
		invSeq := models.LastInvoice{
			FactoryID:   &factoryEnt.ID,
			AffcodeID:   &affcodeData.ID,
			OnYear:      ConvertInt((etd)[:4]),
			LastRunning: 0,
		}

		db.FirstOrCreate(&invSeq, &models.LastInvoice{
			FactoryID: &factoryEnt.ID,
			AffcodeID: &affcodeData.ID,
			OnYear:    ConvertInt((etd)[:4]),
		})

		var order models.Order
		order.ConsigneeID = ord.ConsigneeID
		order.ShipmentID = ord.ShipmentID
		order.EtdDate = &ord.EtdTap
		order.PcID = ord.PcID
		order.CommercialID = ord.CommercialID
		order.SampleFlgID = ord.SampleFlgID
		order.OrderTitleID = &orderTitle.ID
		order.Bioabt = ord.Bioabt
		order.ZoneCode = keyCode
		order.LoadingArea = loadingData.LoadingArea
		order.Privilege = loadingData.Privilege
		order.ShipForm = ord.Bishpc         // bishpc
		order.ShipTo = ord.Bisafn           // bisafn
		order.SampleFlg = ord.SampleFlg     // sample_flg
		order.CarrierCode = ord.CarrierCode // carriercod
		order.RunningSeq = (invSeq.LastRunning + 1)
		order.IsActive = false
		order.IsSync = true

		if err := db.Last(&order, &models.Order{
			ConsigneeID:  ord.ConsigneeID,
			ShipmentID:   ord.ShipmentID,
			EtdDate:      &ord.EtdTap,
			PcID:         ord.PcID,
			CommercialID: ord.CommercialID,
			SampleFlgID:  ord.SampleFlgID,
			Bioabt:       ord.Bioabt,
		}).Error; err != nil {
			if err := db.Save(&order).Error; err == nil {
				invSeq.LastRunning = order.RunningSeq
				db.Save(&invSeq)
			}
		}
		// CreateOrderDetail(&order, &ord)
	}
}

func CreateOrderWithReviseChangeMode(factory, start_etd, end_etd string) {
	db := configs.Store
	var ord []models.OrderPlan
	if err := db.
		Order("upddte,updtime,seq").
		Preload("OrderDetail.Order").
		Where("length(reasoncd) > ?", 0).
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", &factory).
		Where("upddte <= ?", (time.Now()).Format("2006-01-02")).
		Where("substring(reasoncd, 1, 1) not in ?", []string{"Q", "P"}).
		Find(&ord).Error; err != nil {
		sysLogger := models.SyncLogger{
			Title:       "generate order ent revises",
			Description: fmt.Sprintf("%v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	var orderTitle models.OrderTitle
	if err := db.Select("id,title").Where("title=?", "000").First(&orderTitle).Error; err != nil {
		sysLogger := models.SyncLogger{
			Title:       "get order title",
			Description: fmt.Sprintf("Error fetch order title: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	i := 0
	for i < len(ord) {
		GenerateOrderDetailWithReviseChangeMode(end_etd, ord[i], orderTitle)
		i++
	}
}

func GenerateOrderDetailWithReviseChangeMode(endDate string, ord models.OrderPlan, orderTitle models.OrderTitle) {
	db := configs.Store
	/// parse time
	parseEndDate, _ := time.Parse("2006-01-02", endDate)
	// fmt.Printf("End Date: %v ETD: %v > %v is: %v\n", parseEndDate, ord.EtdTap, (ord.EtdTap.After(parseEndDate)), !(ord.EtdTap.After(parseEndDate)))
	if !(ord.EtdTap.After(parseEndDate)) {
		etd := ord.EtdTap.Format("20060102")
		var ship models.Shipment
		db.Select("id,title").First(&ship, "id=?", ord.ShipmentID)
		/// Generate ZoneCode
		var ordCount int64
		db.Select("id").Where("etd_date=?", ord.EtdTap).Find(&models.Order{}).Count(&ordCount)
		sum := ordCount + 1
		keyCode := fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, sum)
		var sumOrder models.Order
		db.Select("id").First(&sumOrder, "zone_code=?", keyCode)
		for !(len(sumOrder.ID) == 0) {
			keyCode = fmt.Sprintf("%s%s%03d", etd[2:], ship.Title, sum)
		}

		// Check LoadingArea
		// fmt.Printf("Check LoadingArea %s\n", ord[x].OrderGroup[:1])
		prefixOrder := "-"
		if ord.OrderGroup[:1] == "@" {
			prefixOrder = "@"
		}
		var loadingData models.OrderLoadingArea
		db.Select("prefix,loading_area,privilege").Where("order_zone_id=?", ord.OrderZoneID).Where("prefix=?", prefixOrder).First(&loadingData)

		var factoryEnt models.Factory
		db.Select("id,title").Where("title=?", ord.Vendor).First(&factoryEnt)
		var affcodeData models.Affcode
		db.Select("id,title,description").Where("title=?", ord.Biac).First(&affcodeData)

		// Get LastInvoiceNo
		invSeq := models.LastInvoice{
			FactoryID:   &factoryEnt.ID,
			AffcodeID:   &affcodeData.ID,
			OnYear:      ConvertInt((etd)[:4]),
			LastRunning: 0,
		}

		db.FirstOrCreate(&invSeq, &models.LastInvoice{
			FactoryID: &factoryEnt.ID,
			AffcodeID: &affcodeData.ID,
			OnYear:    ConvertInt((etd)[:4]),
		})

		// Fetch Order Entries
		var order models.Order
		order.ConsigneeID = ord.ConsigneeID
		order.ShipmentID = ord.ShipmentID
		order.EtdDate = &ord.EtdTap
		order.PcID = ord.PcID
		order.CommercialID = ord.CommercialID
		order.SampleFlgID = ord.SampleFlgID
		order.OrderTitleID = &orderTitle.ID
		order.Bioabt = ord.Bioabt
		order.ZoneCode = keyCode
		order.LoadingArea = loadingData.LoadingArea
		order.Privilege = loadingData.Privilege
		order.ShipForm = ord.Bishpc         // bishpc
		order.ShipTo = ord.Bisafn           // bisafn
		order.SampleFlg = ord.SampleFlg     // sample_flg
		order.CarrierCode = ord.CarrierCode // carriercod
		order.RunningSeq = (invSeq.LastRunning + 1)
		order.IsActive = false
		order.IsSync = true

		if ord.Reasoncd[:1] == "D" { /// แก้ไขวันที่ ETD
			if err := db.Last(&order, &models.Order{
				ConsigneeID:  ord.ConsigneeID,
				ShipmentID:   ord.ShipmentID,
				PcID:         ord.PcID,
				CommercialID: ord.CommercialID,
				SampleFlgID:  ord.SampleFlgID,
				Bioabt:       ord.Bioabt,
			}).Error; err != nil {
				if err := db.Save(&order).Error; err == nil {
					invSeq.LastRunning = order.RunningSeq
					db.Save(&invSeq)
				}
			}
		} else if ord.Reasoncd[:1] == "M" { /// แก้ไขการขนส่ง
			if err := db.First(&order, &models.Order{
				ConsigneeID:  ord.ConsigneeID,
				EtdDate:      &ord.EtdTap,
				PcID:         ord.PcID,
				CommercialID: ord.CommercialID,
				SampleFlgID:  ord.SampleFlgID,
				Bioabt:       ord.Bioabt,
			}).Error; err != nil {
				if err := db.Save(&order).Error; err == nil {
					invSeq.LastRunning = order.RunningSeq
					db.Save(&invSeq)
				}
			}
		} else { /// แก้ไขกรณีอื่นๆเช่น 0,H
			if err := db.First(&order, &models.Order{
				ConsigneeID:  ord.ConsigneeID,
				EtdDate:      &ord.EtdTap,
				ShipmentID:   ord.ShipmentID,
				PcID:         ord.PcID,
				CommercialID: ord.CommercialID,
				SampleFlgID:  ord.SampleFlgID,
				Bioabt:       ord.Bioabt,
			}).Error; err != nil {
				if err := db.Save(&order).Error; err == nil {
					invSeq.LastRunning = order.RunningSeq
					db.Save(&invSeq)
				}
			}
		}
		CreateOrderDetail(&order, &ord)
	} else {
		ctn := 0
		if ord.BalQty > 0 {
			ctn = int(ord.BalQty) / int(ord.Bistdp)
		}
		var orderDetail models.OrderDetail
		if err := db.Preload("Order").First(&orderDetail, &models.OrderDetail{
			Pono:     &ord.Pono,
			LedgerID: ord.LedgerID,
			OrderCtn: int64(ctn),
		}).Error; err == nil {
			if ord.EtdTap.After(parseEndDate) {
				if err := db.Delete(&models.OrderDetail{}, "id", orderDetail.ID).Error; err == nil {
					var countOrdDetail int64
					db.Where("order_id=?", orderDetail.Order.ID).Find(&models.OrderDetail{}).Count(&countOrdDetail)
					if countOrdDetail == 0 {
						if err := db.Delete(&models.Order{}, "id", orderDetail.Order.ID).Error; err != nil {
							panic(err)
						}
					}
					// After Save Check Order Detail
					db.Model(&models.OrderPlan{}).Where("id=?", ord.ID).Update("is_generate", true)
				}
			}
		}
	}
}

func CreateOrderDetail(order *models.Order, ord *models.OrderPlan) {
	db := configs.Store
	if order.ID != "" {
		ctn := 0
		if ord.BalQty > 0 {
			ctn = int(ord.BalQty) / int(ord.Bistdp)
		}

		var orderDetail models.OrderDetail
		orderDetail.OrderID = &order.ID
		orderDetail.Pono = &ord.Pono
		orderDetail.LedgerID = ord.LedgerID
		orderDetail.OrderPlanID = &ord.ID
		orderDetail.OrderCtn = int64(ctn)
		orderDetail.TotalOnPallet = 0
		if err := db.FirstOrCreate(&orderDetail, &models.OrderDetail{
			Pono:     &ord.Pono,
			LedgerID: ord.LedgerID,
			OrderCtn: int64(ctn),
		}).Error; err != nil {
			panic(err)
		}

		// After Save Check Order Detail
		db.Model(&models.OrderDetail{}).Where("id=?", &orderDetail.ID).Updates(&models.OrderDetail{
			OrderID:     &order.ID,
			Pono:        &ord.Pono,
			LedgerID:    ord.LedgerID,
			OrderPlanID: &ord.ID,
			OrderCtn:    int64(ctn),
		})
		if err := db.Model(&models.OrderPlan{}).Where("id=?", ord.ID).Update("is_generate", true).Error; err == nil {
			DeleteOrder(ord)
		}
	}
}

func DeleteOrder(ord *models.OrderPlan) {
	db := configs.Store
	var order []models.Order
	db.Order("created_at desc").Select("id").Find(&order, &models.Order{
		ConsigneeID:  ord.ConsigneeID,
		PcID:         ord.PcID,
		CommercialID: ord.CommercialID,
		SampleFlgID:  ord.SampleFlgID,
		Bioabt:       ord.Bioabt,
	})

	for _, v := range order {
		var countOrdDetail int64
		db.Select("id").Where("order_id=?", v.ID).Find(&models.OrderDetail{}).Count(&countOrdDetail)
		if countOrdDetail == 0 {
			if err := db.Delete(&models.Order{}, "id", v.ID).Error; err != nil {
				fmt.Printf("Err: %s\n", err.Error())
			}
		}
	}
}

func GetOrderGroup(c *fiber.Ctx) []string {
	db := configs.Store
	s := c.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")
	var userID string
	err := db.Select("user_id").First(&models.JwtToken{}, "id=?", token).Scan(&userID).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	var orderGroup []models.OrderGroup
	db.Find(&orderGroup, "user_id=?", &userID)
	conID := []string{}
	for _, v := range orderGroup {
		conID = append(conID, *v.ConsigneeID)
	}
	return conID
}

func ClearOrder() {
	db := configs.Store
	var orderList []string
	db.Raw("select t.id from tbt_orders t left join tbt_order_details d on t.id=d.order_id where d.id is null").Scan(&orderList)
	for _, id := range orderList {
		if err := db.Delete(&models.Order{}, "id=?", id).Error; err == nil {
			fmt.Printf("Delete Order ID: %s\n", id)
		}
	}
}
