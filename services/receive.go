package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/shakinm/xlsReader/xls"
)

func CheckSerialIsReceived(serial_no string) bool {

	url := fmt.Sprintf("%s/carton/search?serial_no=%s", os.Getenv("API_TRIGGER_URL"), serial_no)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()
	return res.StatusCode != 404
}

func SummaryReceive(transferOutNo *string) {
	db := configs.Store
	var receiveDetail []models.ReceiveDetail
	db.Select("id").Where("receive_id=?", transferOutNo).Find(&receiveDetail)
	var ctn int64 = 0
	for _, v := range receiveDetail {
		var countReceive int64
		db.Select("id").Where("receive_detail_id=?", v.ID).Find(&models.Carton{}).Count(&countReceive)
		ctn += countReceive
	}
	// fmt.Printf("TransferOutNo: %s Total: %d\n", *transferOutNo, ctn)
	var recEnt models.Receive
	db.First(&recEnt, transferOutNo)
	recEnt.ReceiveCtn = ctn
	db.Save(&recEnt)
}

func ImportReceiveCarton(path string) (err error) {
	db := configs.Store
	workbook, err := xls.OpenFile(path)
	if err != nil {
		return
	}

	for _, sheet := range workbook.GetSheets() {
		i := 0
		// var transferNo *string
		for _, r := range sheet.GetRows() {
			if i > 0 {
				rec_no, err := r.GetCol(1)
				if err == nil && rec_no.GetString() != "" {
					if rec_no.GetString() != "" {
						transfer_out_no, _ := r.GetCol(1)
						var receiveEnt models.Receive
						if err := db.Preload("ReceiveType").Where("transfer_out_no=?", transfer_out_no.GetString()).First(&receiveEnt).Error; err != nil {
							panic(err)
						}
						part_no, _ := r.GetCol(2)
						lot_no, _ := r.GetCol(3)
						serial_no, _ := r.GetCol(5)
						q, _ := r.GetCol(6)
						qty, _ := strconv.ParseInt(q.GetString(), 10, 64)
						if qty > 0 {
							var part models.Part
							if err := db.Select("id").Where("title=?", part_no.GetString()).First(&part).Error; err != nil {
								panic(err)
							}

							var ledger models.Ledger
							if err := db.Select("id").Where("whs_id=?", &receiveEnt.ReceiveType.WhsID).Where("part_id=?", &part.ID).First(&ledger).Error; err != nil {
								panic(err)
							}

							fmt.Printf("%s ==> RecENT: %s PartID: %s LedgerID: %s\n", rec_no.GetString(), receiveEnt.ID, part.ID, ledger.ID)
							var recDetail models.ReceiveDetail
							if err := db.Select("id,plan_qty").Where("receive_id=?", &receiveEnt.ID).Where("ledger_id=?", &ledger.ID).First(&recDetail).Error; err != nil {
								fmt.Println(err.Error())
								recDetail.ReceiveID = &receiveEnt.ID
								recDetail.LedgerID = &ledger.ID
								recDetail.PlanQty = qty
								recDetail.PlanCtn = 1
								db.Save(&recDetail)
							}

							var receiveCarton models.CartonNotReceive
							receiveCarton.ReceiveDetailID = recDetail.ID
							receiveCarton.TransferOutNo = rec_no.GetString()
							receiveCarton.PartNo = part_no.GetString()
							receiveCarton.LotNo = lot_no.GetString()
							receiveCarton.SerialNo = serial_no.GetString()
							receiveCarton.Qty = qty
							if err := db.FirstOrCreate(&receiveCarton, &models.CartonNotReceive{SerialNo: serial_no.GetString()}).Error; err != nil {
								panic(err)
							}

							receiveCarton.IsReceived = CheckSerialIsReceived(serial_no.GetString())
							if err := db.Save(&receiveCarton).Error; err != nil {
								panic(err)
							}

							// var recCtn int64
							// if err := db.Raw("select count(id) from tbt_carton_not_receives where is_received=true and transfer_out_no=?", rec_no.GetString()).Scan(&recCtn).Error; err != nil {
							// 	panic(err)
							// }
							// fmt.Printf("%s ==> %d\n", rec_no.GetString(), recCtn)
						}
					}
				}
			}
			i++
		}

		// db.Where("is_sync=?", true).Delete(&models.CartonNotReceive{})
		// go SummaryReceive(transferNo)
	}
	return
}
