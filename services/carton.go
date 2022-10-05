package services

import (
	"fmt"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
)

func CreateCartonHistoryData(obj *models.CartonHistory) bool {
	db := configs.Store
	var sysLog models.SyncLogger
	sysLog.Title = "Create Carton history"
	sysLog.Description = fmt.Sprintf("%s created successfully", obj.SerialNo)
	sysLog.IsSuccess = true
	err := db.Create(&obj).Error
	if err != nil {
		sysLog.Title = fmt.Sprintf("Create %s is error", obj.SerialNo)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
	}
	db.Create(&sysLog)
	go CreateCarton(obj)
	return true
}

func CreateCarton(obj *models.CartonHistory) {
	db := configs.Store
	var sysLog models.SyncLogger
	sysLog.Title = "Create Carton"
	sysLog.Description = fmt.Sprintf("%s created successfully", obj.SerialNo)
	sysLog.IsSuccess = true
	// Get Master Data
	if obj.Whs == "C" {
		obj.Whs = "COM"
	}

	switch obj.Whs {
	case "C":
		obj.Whs = "COM"
	case "D":
		obj.Whs = "DOM"
	case "N":
		obj.Whs = "NESC"
	case "I":
		obj.Whs = "ICAM"
	default:
		obj.Whs = "COM"
	}

	var whs models.Whs
	db.First(&whs, "title=?", obj.Whs)
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

	var cartonData models.Carton
	cartonData.RowID = obj.RowID
	cartonData.LedgerID = &ledger.ID
	cartonData.LocationID = &shelve.ID
	// cartonData.ReceiveDetailID
	cartonData.LotNo = obj.LotNo
	cartonData.SerialNo = obj.SerialNo
	cartonData.LineNo = obj.DieNo
	cartonData.RevisionNo = obj.RevisionNo
	cartonData.Qty = float64(obj.Qty)
	cartonData.PalletNo = obj.RefNo
	cartonData.IsActive = true
	err := db.Create(&cartonData).Error
	if err != nil {
		sysLog.Title = fmt.Sprintf("Create %s is error", obj.SerialNo)
		sysLog.Description = err.Error()
		sysLog.IsSuccess = false
	}
	db.Create(&obj)
}
