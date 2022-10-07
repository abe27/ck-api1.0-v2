package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/shakinm/xlsReader/xls"
)

func CheckSerialIsReceived(serial_no string) bool {

	url := fmt.Sprintf("http://127.0.0.1:4000/carton/search?serial_no=%s", serial_no)
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
	var receEnt models.Receive
	db.First(&receEnt, transferOutNo)
	receEnt.ReceiveCtn = ctn
	db.Save(&receEnt)
}

func ImportReceiveCarton(path string) (err error) {
	workbook, err := xls.OpenFile(path)
	if err != nil {
		return
	}

	sheet, err := workbook.GetSheet(0)
	if err != nil {
		return
	}

	// sname := sheet.GetName()
	var transferNo *string
	db := configs.Store
	for i := 0; i <= sheet.GetNumberRows(); i++ {
		if i > 0 {
			if r, err := sheet.GetRow(i); err == nil {
				rec_no, _ := r.GetCol(1)
				if rec_no.GetString() != "" {
					transfer_out_no, _ := r.GetCol(1)
					part_no, _ := r.GetCol(2)
					lot_no, _ := r.GetCol(3)
					serial_no, _ := r.GetCol(5)
					qty, _ := r.GetCol(6)

					isFound := CheckSerialIsReceived(serial_no.GetString())
					cartonNotReceive := models.CartonNotReceive{
						TransferOutNo: transfer_out_no.GetString(),
						PartNo:        part_no.GetString(),
						LotNo:         lot_no.GetString(),
						SerialNo:      serial_no.GetString(),
						Qty:           qty.GetInt64(),
						IsSync:        isFound,
					}
					db.FirstOrCreate(&cartonNotReceive, models.CartonNotReceive{SerialNo: serial_no.GetString()})
					if isFound {
						var recEnt models.Receive
						db.Preload("FileEdi.Factory").Preload("ReceiveType.Whs").First(&recEnt, "transfer_out_no=?", transfer_out_no.GetString())
						if recEnt.ID != "" {
							var part models.Part
							db.Select("id").First(&part, "slug=?", strings.ReplaceAll(part_no.GetString(), "-", ""))
							var ledger models.Ledger
							db.Select("id").Where("whs_id=?", &recEnt.ReceiveType.WhsID).Where("factory_id=?", &recEnt.FileEdi.FactoryID).Where("part_id=?", &part.ID).First(&ledger)

							// Search Receive Detail
							var receiveDetail models.ReceiveDetail
							db.Select("id").Where("receive_id=?", &recEnt.ID).Where("ledger_id=?", &ledger.ID).First(&receiveDetail)

							var carton models.Carton
							db.First(&carton, "serial_no=?", serial_no.GetString())
							if carton.ID != "" {
								carton.ReceiveDetailID = &receiveDetail.ID
								db.Save(&carton)
							}

							if transferNo != nil && transferNo != &recEnt.ID {
								go SummaryReceive(transferNo)
							}
							transferNo = &recEnt.ID
						}
					}
					fmt.Printf("TransferOutNo: %s PartNo: %s SerialNO: %s is: %v\n", transfer_out_no.GetString(), part_no.GetString(), serial_no.GetString(), isFound)
				}
			}
		}
	}
	db.Where("is_sync=?", true).Delete(&models.CartonNotReceive{})
	go SummaryReceive(transferNo)
	return
}
