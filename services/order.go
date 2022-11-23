package services

import (
	"fmt"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateOrder(factory, start_etd, end_etd string) {
	// Group Order to create order ent
	db := configs.Store
	// etd := (time.Now()).Format("2006-01-02")
	// fmt.Printf("Create Order: %v\n", etd)
	var ord []models.OrderPlan
	err := db.
		Order("etd_tap,shipment_id").
		Select("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,sample_flg_id,carrier_code").
		Where("substring(reasoncd, 1, 1) in (?)", []string{"", "0", "-", "H"}).
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", &factory).
		Where("etd_tap BETWEEN ? AND ?", start_etd, end_etd).
		Group("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,sample_flg_id,carrier_code").
		Find(&ord).Error
	if err != nil {
		sysLogger := models.SyncLogger{
			Title:       "generate order ent",
			Description: fmt.Sprintf("Error fetch order: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	// fmt.Printf("Fetch Order Title: 000\n")

	var orderTitle models.OrderTitle
	err = db.Select("id,title").Where("title=?", "000").First(&orderTitle).Error
	if err != nil {
		sysLogger := models.SyncLogger{
			Title:       "get order title",
			Description: fmt.Sprintf("Error fetch order title: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	// fmt.Printf("Fetch Order Title: %s\n", orderTitle.ID)

	x := 0
	for x < len(ord) {
		GenerateOrderDetail(ord[x], orderTitle)
		x++
	}

	CreateOrderWithRevise(factory, start_etd, end_etd)
}

func CreateOrderWithRevise(factory, start_etd, end_etd string) {
	db := configs.Store
	var ord []models.OrderPlan
	if err := db.
		Order("seq").
		Where("substring(reasoncd, 1, 1) not in (?)", []string{"", "0", "-", "H"}).
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", &factory).
		Where("etd_tap BETWEEN ? AND ?", start_etd, end_etd).
		Find(&ord).Error; err != nil {
		sysLogger := models.SyncLogger{
			Title:       "generate order ent",
			Description: fmt.Sprintf("Error fetch order: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	// fmt.Printf("Fetch Order Title: 000\n")

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
		obj := ord[i]
		fmt.Printf("%d::: %d ==> %s\n", i, obj.Seq, obj.Reasoncd[:1])
		GenerateOrderDetailWithRevise(ord[i], orderTitle)
		// if obj.Reasoncd[:1] != "D" {
		// 	GenerateOrderDetailWithRevise(ord[i], orderTitle)
		// }
		i++
	}
}

func GenerateOrderDetail(ord models.OrderPlan, orderTitle models.OrderTitle) {
	db := configs.Store
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

	order := models.Order{
		ConsigneeID:  ord.ConsigneeID,
		ShipmentID:   ord.ShipmentID,
		EtdDate:      &ord.EtdTap,
		PcID:         ord.PcID,
		CommercialID: ord.CommercialID,
		SampleFlgID:  ord.SampleFlgID,
		OrderTitleID: &orderTitle.ID,
		Bioat:        ord.Bioabt,
		ZoneCode:     keyCode,
		LoadingArea:  loadingData.LoadingArea,
		Privilege:    loadingData.Privilege,
		ShipForm:     ord.Bishpc,      // bishpc,
		ShipTo:       ord.Bisafn,      // bisafn,
		SampleFlg:    ord.SampleFlg,   // sample_flg,
		CarrierCode:  ord.CarrierCode, // carriercode
		RunningSeq:   (invSeq.LastRunning + 1),
		IsActive:     false,
		IsSync:       true,
	}

	err := db.FirstOrCreate(&order, &models.Order{
		ConsigneeID:  ord.ConsigneeID,
		ShipmentID:   ord.ShipmentID,
		EtdDate:      &ord.EtdTap,
		PcID:         ord.PcID,
		CommercialID: ord.CommercialID,
		SampleFlgID:  ord.SampleFlgID,
		Bioat:        ord.Bioabt,
	}).Error

	if err != nil {
		// Create log if Create Order is Error!
		sysLogger := models.SyncLogger{
			Title:       fmt.Sprintf("creating order %v", ord.ConsigneeID),
			Description: fmt.Sprintf("Error creating order: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	// update lastinvoice no
	invSeq.LastRunning += 1
	db.Save(&invSeq)

	if order.ID != "" {
		// Create Order Detail
		var orderPlan []models.OrderPlan
		err = db.Order("etd_tap,created_at,seq").
			Where("substring(reasoncd, 1, 1) in (?)", []string{"", "0", "-", "H"}).
			Where("is_revise_error=?", false).
			Where("is_generate=?", false).
			Where("order_zone_id=?", ord.OrderZoneID).
			Where("consignee_id=?", ord.ConsigneeID).
			Where("shipment_id=?", ord.ShipmentID).
			Where("etd_tap=?", ord.EtdTap).
			Where("pc_id=?", ord.PcID).
			Where("commercial_id=?", ord.CommercialID).
			Where("sample_flg_id=?", ord.SampleFlgID).
			Where("bioabt=?", ord.Bioabt).
			Where("order_group=?", ord.OrderGroup).
			Find(&orderPlan).Error
		if err != nil {
			// Create log if Create Order is Error!
			sysLogger := models.SyncLogger{
				Title:       "fetch order plan",
				Description: fmt.Sprintf("Error fetch order: %v", err),
			}
			db.Create(&sysLogger)
			panic(err)
		}

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
			ordDetail.IsSync = true
			db.Save(&ordDetail)

			// Update Order Plan Set Status Generated
			ordPlan := &r
			ordPlan.IsGenerate = true
			db.Save(&ordPlan)
			j++
		}
	}
}

func GenerateOrderDetailWithRevise(ord models.OrderPlan, orderTitle models.OrderTitle) {
	db := configs.Store
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

	order := models.Order{
		ConsigneeID:  ord.ConsigneeID,
		ShipmentID:   ord.ShipmentID,
		EtdDate:      &ord.EtdTap,
		PcID:         ord.PcID,
		CommercialID: ord.CommercialID,
		SampleFlgID:  ord.SampleFlgID,
		OrderTitleID: &orderTitle.ID,
		Bioat:        ord.Bioabt,
		ZoneCode:     keyCode,
		LoadingArea:  loadingData.LoadingArea,
		Privilege:    loadingData.Privilege,
		ShipForm:     ord.Bishpc,      // bishpc,
		ShipTo:       ord.Bisafn,      // bisafn,
		SampleFlg:    ord.SampleFlg,   // sample_flg,
		CarrierCode:  ord.CarrierCode, // carriercode
		RunningSeq:   (invSeq.LastRunning + 1),
		IsActive:     false,
		IsSync:       true,
	}

	err := db.FirstOrCreate(&order, &models.Order{
		ConsigneeID:  ord.ConsigneeID,
		ShipmentID:   ord.ShipmentID,
		EtdDate:      &ord.EtdTap,
		PcID:         ord.PcID,
		CommercialID: ord.CommercialID,
		SampleFlgID:  ord.SampleFlgID,
		Bioat:        ord.Bioabt,
	}).Error

	if err != nil {
		// Create log if Create Order is Error!
		sysLogger := models.SyncLogger{
			Title:       fmt.Sprintf("creating order %v", ord.ConsigneeID),
			Description: fmt.Sprintf("Error creating order: %v", err),
		}
		db.Create(&sysLogger)
		panic(err)
	}

	// update lastinvoice no
	invSeq.LastRunning += 1
	db.Save(&invSeq)

	if order.ID != "" {
		// Create Order Detail
		var orderPlan models.OrderPlan
		err = db.First(&orderPlan, orderPlan.ID).Error
		if err != nil {
			// Create log if Create Order is Error!
			sysLogger := models.SyncLogger{
				Title:       "fetch order plan",
				Description: fmt.Sprintf("Error fetch order: %v", err),
			}
			db.Create(&sysLogger)
			panic(err)
		}

		r := orderPlan
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
		ordDetail.IsSync = true
		db.Save(&ordDetail)

		// Update Order Plan Set Status Generated
		ordPlan := &r
		ordPlan.IsGenerate = true
		db.Save(&ordPlan)
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
