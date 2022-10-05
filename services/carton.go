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
	whsTite := obj.Whs
	switch whsTite {
	case "C":
		whsTite = "COM"
	case "D":
		whsTite = "DOM"
	case "N":
		whsTite = "NESC"
	case "I":
		whsTite = "ICAM"
	default:
		whsTite = "COM"
	}

	var whs models.Whs
	db.First(&whs, "title=?", whsTite)
	var part models.Part
	db.First(&part, "slug=?", strings.ReplaceAll(obj.PartNo, "-", ""))

	facTitle := "INJ"
	if obj.PartNo[:1] == "1" {
		facTitle = "AW"
	}

	var fac models.Factory
	db.First(&fac, "title=?", facTitle)
	var ledger models.Ledger
	db.Where("whs_id=?", whs.ID).Where("part_id=?", part.ID).Where("factory_id=?", fac.ID).First(&ledger)

	var shelve models.Location
	db.First(&shelve, "title=?", strings.ReplaceAll(obj.Shelve, " ", ""))

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
	err := db.FirstOrCreate(&cartonData, &models.Carton{SerialNo: obj.SerialNo}).Error
	if err != nil {
		var sysLog models.SyncLogger
		sysLog.Title = fmt.Sprintf("Create %s is error", obj.SerialNo)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
		db.Create(&sysLog)
	}

	var ctn int64
	db.Select("id").Where("ledger_id=?", ledger.ID).Where("qty > ?", "0").Find(&models.Carton{}).Count(&ctn)
	ledger.Ctn = float64(ctn)
	db.Save(&ledger)
}
