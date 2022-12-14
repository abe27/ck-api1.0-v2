package services

import (
	"fmt"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
)

func CreateCartonHistoryData(obj *models.CartonHistory) bool {
	db := configs.Store
	// var sysLog models.SyncLogger
	// sysLog.Title = "Create Carton history"
	// sysLog.Description = fmt.Sprintf("%s created successfully", obj.SerialNo)
	// sysLog.IsSuccess = true
	err := db.Create(&obj).Error
	if err != nil {
		var sysLog models.SyncLogger
		sysLog.Title = fmt.Sprintf("Create %s is error", obj.SerialNo)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
		db.Create(&sysLog)
	}
	return true
}

func CreateCarton(obj *models.CartonHistory) {
	db := configs.Store
	// var sysLog models.SyncLogger
	// sysLog.Title = "Create Carton"
	// sysLog.Description = fmt.Sprintf("%s created successfully", obj.SerialNo)
	// sysLog.IsSuccess = true
	// Get Master Data
	whsTitle := obj.Whs
	switch whsTitle {
	case "C":
		whsTitle = "COM"
	case "D":
		whsTitle = "DOM"
	case "N":
		whsTitle = "NESC"
	case "I":
		whsTitle = "ICAM"
	default:
		whsTitle = "COM"
	}

	var whs models.Whs
	db.First(&whs, "title=?", whsTitle)
	part := models.Part{
		Slug:        strings.ReplaceAll(obj.PartNo, "-", ""),
		Title:       obj.PartNo,
		Description: obj.PartNo,
		IsActive:    true,
	}
	db.FirstOrCreate(&part, &models.Part{
		Slug: strings.ReplaceAll(obj.PartNo, "-", ""),
	})

	facTitle := "INJ"
	unitTitle := "BOX"
	partType := "PART"
	if obj.PartNo[:1] == "1" {
		facTitle = "AW"
		unitTitle = "COIL"
		partType = "WIRE"
	}

	var partTypeData models.PartType
	db.First(&partTypeData, "title=?", partType)

	var unitData models.Unit
	db.First(&unitData, "title=?", unitTitle)

	var fac models.Factory
	db.First(&fac, "title=?", facTitle)
	ledger := models.Ledger{
		WhsID:       &whs.ID,
		PartID:      &part.ID,
		FactoryID:   &fac.ID,
		PartTypeID:  &partTypeData.ID,
		UnitID:      &unitData.ID,
		DimWidth:    0,
		DimLength:   0,
		DimHeight:   0,
		GrossWeight: 0,
		NetWeight:   0,
		Qty:         0,
		Ctn:         0,
	}

	err := db.FirstOrCreate(&ledger, &models.Ledger{
		WhsID:     &whs.ID,
		PartID:    &part.ID,
		FactoryID: &fac.ID,
	}).Error
	if err != nil {
		var sysLog models.SyncLogger
		sysLog.Title = fmt.Sprintf("Create %s on %s ledger is error", obj.PartNo, whs.Title)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
		db.Create(&sysLog)
		return
	}

	shelveTitle := strings.ReplaceAll(obj.Shelve, " ", "")
	if len(shelveTitle) == 0 {
		shelveTitle = "-"
	}
	shelve := models.Location{
		Title:       shelveTitle,
		Description: shelveTitle,
		IsActive:    true,
	}
	db.FirstOrCreate(&shelve, &models.Location{
		Title: shelveTitle,
	})

	cartonData := models.Carton{
		RowID:      obj.RowID,
		LedgerID:   &ledger.ID,
		LocationID: &shelve.ID,
		LotNo:      obj.LotNo,
		SerialNo:   obj.SerialNo,
		LineNo:     obj.DieNo,
		RevisionNo: obj.RevisionNo,
		Qty:        float64(obj.Qty),
		PalletNo:   obj.RefNo,
		IsActive:   true,
	}
	err = db.FirstOrCreate(&cartonData, &models.Carton{SerialNo: obj.SerialNo}).Error
	if err != nil {
		var sysLog models.SyncLogger
		sysLog.Title = fmt.Sprintf("Create %s on %s is error", obj.SerialNo, whs.Title)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
		db.Create(&sysLog)

		olderID := &cartonData.LedgerID
		// Save Carton New Ledger
		cartonData.RowID = obj.RowID
		cartonData.LedgerID = &ledger.ID
		cartonData.LocationID = &shelve.ID
		cartonData.Qty = float64(obj.Qty)
		cartonData.PalletNo = obj.RefNo
		cartonData.IsActive = true
		db.Save(&cartonData)

		// Update Older Stock
		var ctn int64
		db.Select("id").Where("ledger_id=?", &olderID).Where("qty > ?", "0").Find(&models.Carton{}).Count(&ctn)
		var olderLedger *models.Ledger
		db.Select("id,ctn").First(&olderLedger, "id=?", &olderID)
		olderLedger.Ctn = float64(ctn)
		db.Save(&olderLedger)
	}

	var ctn int64
	db.Raw(fmt.Sprintf("SELECT count(id) FROM tbt_cartons WHERE ledger_id='%s' AND qty > '0'", ledger.ID)).Scan(&ctn)
	ledger.Qty = 0
	if obj.Qty > 0 {
		ledger.Qty = float64(obj.Qty)
	}
	ledger.Ctn = float64(ctn)
	db.Save(&ledger)
}
