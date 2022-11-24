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
	var ord []models.OrderPlan
	err := db.
		Order("etd_tap,shipment_id").
		Select("order_zone_id,consignee_id,shipment_id,etd_tap,pc_id,commercial_id,bioabt,order_group,vendor,biac,bishpc,bisafn,sample_flg_id,carrier_code").
		// Where("substring(reasoncd, 1, 1) in (?)", []string{"", "0", "-", "H"}).
		Where("length(reasoncd) = ?", 0).
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

	// Revise Qty and Create
	CreateOrderWithRevise(factory, start_etd, end_etd)
	// Revise Change Shipment Mode
	CreateOrderWithReviseChangeMode(factory, start_etd, end_etd)
	// Revise Change Date Mode
	CreateOrderWithReviseChangeDateMode(factory, start_etd, end_etd)
}

func CreateOrderWithRevise(factory, start_etd, end_etd string) {
	db := configs.Store
	var ord []models.OrderPlan
	if err := db.
		Order("reasoncd,upddte,updtime,seq").
		Preload("OrderDetail.Order").
		Where("length(reasoncd) > ?", 0).
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", &factory).
		Where("etd_tap BETWEEN ? AND ?", start_etd, end_etd).
		Where("substring(reasoncd, 1, 1) not in ?", []string{"D", "M"}).
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
		obj := ord[i]
		// GenerateOrderDetailWithRevise(ord[i], orderTitle)
		if obj.Reasoncd[:1] != "D" && obj.Reasoncd[:1] != "M" {
			GenerateOrderDetailWithRevise(ord[i], orderTitle)
		}
		i++
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
		Where("etd_tap BETWEEN ? AND ?", start_etd, end_etd).
		Where("substring(reasoncd, 1, 1) in ?", []string{"M"}).
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
		obj := ord[i]
		if obj.Reasoncd[:1] == "M" {
			GenerateOrderDetailWithReviseChangeMode(ord[i], orderTitle)
		}
		i++
	}
}

func CreateOrderWithReviseChangeDateMode(factory, start_etd, end_etd string) {
	db := configs.Store
	var ord []models.OrderPlan
	if err := db.
		Order("upddte,updtime,seq").
		Preload("OrderDetail.Order").
		Where("length(reasoncd) > ?", 0).
		Where("is_generate=?", false).
		Where("is_revise_error=?", false).
		Where("vendor=?", &factory).
		Where("upddte BETWEEN ? AND ?", start_etd, end_etd).
		Where("substring(reasoncd, 1, 1) in ?", []string{"D"}).
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
		obj := ord[i]
		if obj.Reasoncd[:1] == "D" {
			// GenerateOrderDetailWithReviseChangeMode(ord[i], orderTitle)
			fmt.Printf("%s\n", obj.EtdTap)
		}
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

	var orderCountDuplicate int64
	db.Select("id").
		Where("consignee_id=?", ord.ConsigneeID).
		Where("shipment_id=?", ord.ShipmentID).
		Where("etd_date=?", &ord.EtdTap).
		Where("pc_id=?", ord.PcID).
		Where("commercial_id=?", ord.CommercialID).
		Where("sample_flg_id=?", ord.SampleFlgID).
		Where("bioabt=?", ord.Bioabt).
		Find(&models.Order{}).Count(&orderCountDuplicate)

	if orderCountDuplicate == 0 {
		if err := db.Create(&order).Error; err != nil {
			// Create log if Create Order is Error!
			sysLogger := models.SyncLogger{
				Title:       fmt.Sprintf("creating order %v", ord.ConsigneeID),
				Description: fmt.Sprintf("Error creating order: %v", err),
			}
			db.Create(&sysLogger)
			panic(err)
		}
	}

	if order.ID != "" {
		// Create Order Detail
		var orderPlan []models.OrderPlan
		if err := db.Order("etd_tap,created_at,seq").
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
			Where("length(reasoncd) = ?", 0).
			Find(&orderPlan).Error; err != nil {
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
			db.Save(&ordDetail)

			// Update Order Plan Set Status Generated
			ordPlan := &r
			ordPlan.IsGenerate = true
			db.Save(&ordPlan)
			j++
		}

		// update lastinvoice no
		invSeq.LastRunning += 1
		db.Save(&invSeq)
	}
}

func GenerateOrderDetailWithRevise(ord models.OrderPlan, orderTitle models.OrderTitle) {
	// fmt.Printf("%s ::: %d\n", ord.ID, len(ord.OrderDetail))
	ctn := 0
	if ord.BalQty > 0 {
		ctn = int(ord.BalQty) / int(ord.Bistdp)
	}
	db := configs.Store
	// Check LoadingArea
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

	// Start Process
	var order models.Order
	if len(ord.OrderDetail) == 0 {
		// fmt.Printf("Create OrderDetail\n")
		// ค้นหา order ent
		db.First(&order, &models.Order{
			ConsigneeID:  ord.ConsigneeID,
			ShipmentID:   ord.ShipmentID,
			EtdDate:      &ord.EtdTap,
			PcID:         ord.PcID,
			CommercialID: ord.CommercialID,
			SampleFlgID:  ord.SampleFlgID,
			Bioabt:       ord.Bioabt,
		})
		if order.ID != "" {
			SaveOrderDetail(&order, &ord, &ctn)
		} else {
			// บันทึกข้อมูล Inv
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

			// // Get LastInvoiceNo
			var invSeq models.LastInvoice
			invSeq.FactoryID = &factoryEnt.ID
			invSeq.AffcodeID = &affcodeData.ID
			invSeq.OnYear = ConvertInt((etd)[:4])
			invSeq.LastRunning = 0

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

			if err := db.FirstOrCreate(&order, &models.Order{
				ConsigneeID:  ord.ConsigneeID,
				ShipmentID:   ord.ShipmentID,
				EtdDate:      &ord.EtdTap,
				PcID:         ord.PcID,
				CommercialID: ord.CommercialID,
				SampleFlgID:  ord.SampleFlgID,
				Bioabt:       ord.Bioabt,
			}).Error; err == nil {
				invSeq.LastRunning += 1
				db.Save(&invSeq)
			}

			if order.ID != "" {
				SaveOrderDetail(&order, &ord, &ctn)
			}
			fmt.Printf("Create ZoneCode: %s\n", keyCode)
		}
	}
}

func GenerateOrderDetailWithReviseChangeMode(ord models.OrderPlan, orderTitle models.OrderTitle) {
	// fmt.Printf("%s ::: %d\n", ord.ID, len(ord.OrderDetail))
	ctn := 0
	if ord.BalQty > 0 {
		ctn = int(ord.BalQty) / int(ord.Bistdp)
	}
	db := configs.Store
	// Check LoadingArea
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

	// Start Process
	var order models.Order
	if len(ord.OrderDetail) == 0 {
		// ค้นหา order ent
		db.First(&order, &models.Order{
			ConsigneeID:  ord.ConsigneeID,
			ShipmentID:   ord.ShipmentID,
			EtdDate:      &ord.EtdTap,
			PcID:         ord.PcID,
			CommercialID: ord.CommercialID,
			SampleFlgID:  ord.SampleFlgID,
			Bioabt:       ord.Bioabt,
			// IsInvoice:    false,
		})
		if order.ID != "" {
			// fmt.Printf("Update Order ID: %s ==> %d\n", order.ID, ctn)
			SaveOrderChangeModeDetail(&order, &ord, &ctn)
		} else {
			// fmt.Printf("Create Order ID: %s\n", order.ID)
			// บันทึกข้อมูล Inv
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

			// // Get LastInvoiceNo
			var invSeq models.LastInvoice
			invSeq.FactoryID = &factoryEnt.ID
			invSeq.AffcodeID = &affcodeData.ID
			invSeq.OnYear = ConvertInt((etd)[:4])
			invSeq.LastRunning = 0

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

			if err := db.FirstOrCreate(&order, &models.Order{
				ConsigneeID:  ord.ConsigneeID,
				ShipmentID:   ord.ShipmentID,
				EtdDate:      &ord.EtdTap,
				PcID:         ord.PcID,
				CommercialID: ord.CommercialID,
				SampleFlgID:  ord.SampleFlgID,
				Bioabt:       ord.Bioabt,
			}).Error; err == nil {
				invSeq.LastRunning += 1
				db.Save(&invSeq)
			}

			if order.ID != "" {
				SaveOrderChangeModeDetail(&order, &ord, &ctn)
			}
			fmt.Printf("Create ZoneCode: %s\n", keyCode)
		}
	}
}

func SaveOrderChangeModeDetail(order *models.Order, ord *models.OrderPlan, ctn *int) {
	db := configs.Store
	var ordDetail models.OrderDetail
	ordDetail.OrderID = &order.ID
	ordDetail.Pono = &ord.Pono
	ordDetail.LedgerID = ord.LedgerID
	ordDetail.OrderPlanID = &ord.ID
	ordDetail.OrderCtn = int64(*ctn)
	ordDetail.TotalOnPallet = 0

	if err := db.FirstOrCreate(&ordDetail, &models.OrderDetail{
		OrderID:  &order.ID,
		Pono:     &ord.Pono,
		LedgerID: ord.LedgerID,
		OrderCtn: int64(*ctn),
	}).Error; err != nil {
		panic(err)
	}

	ordDetail.OrderID = &order.ID
	ordDetail.OrderPlanID = &ord.ID
	ordDetail.OrderCtn = int64(*ctn)
	ordDetail.IsSync = true
	if err := db.Save(&ordDetail).Error; err != nil {
		sysLogger := models.SyncLogger{
			Title:       fmt.Sprintf("update status order detail %v", ord.ID),
			Description: fmt.Sprintf("%v", err),
		}
		db.Save(sysLogger)
		panic(err)
	}

	// Update Order Plan Set Status Generated
	ord.IsGenerate = true
	if err := db.Save(&ord).Error; err != nil {
		panic(err)
	}

	var checkOrderDetail int64
	db.Select("id").Where("order_id=?", order.ID).Find(&models.OrderDetail{}).Count(&checkOrderDetail)
	if checkOrderDetail == 0 {
		if err := db.Delete(&models.Order{}, order.ID).Error; err != nil {
			panic(err)
		}
	}
}

func SaveOrderDetail(order *models.Order, ord *models.OrderPlan, ctn *int) {
	db := configs.Store
	var ordDetail models.OrderDetail
	ordDetail.OrderID = &order.ID
	ordDetail.Pono = &ord.Pono
	ordDetail.LedgerID = ord.LedgerID
	ordDetail.OrderPlanID = &ord.ID
	ordDetail.OrderCtn = int64(*ctn)
	ordDetail.TotalOnPallet = 0

	if err := db.FirstOrCreate(&ordDetail, &models.OrderDetail{OrderPlanID: &ord.ID}).Error; err == nil {
		// Confirm Data After Create
		ordDetail.OrderID = &order.ID
		ordDetail.OrderPlanID = &ord.ID
		ordDetail.OrderCtn = int64(*ctn)
		ordDetail.IsSync = true
		if err := db.Save(&ordDetail).Error; err != nil {
			sysLogger := models.SyncLogger{
				Title:       fmt.Sprintf("update status order detail %v", ord.ID),
				Description: fmt.Sprintf("%v", err),
			}
			db.Save(sysLogger)
			panic(err)
		}

		// Update Order Plan Set Status Generated
		ord.IsGenerate = true
		if err := db.Save(&ord).Error; err != nil {
			panic(err)
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
