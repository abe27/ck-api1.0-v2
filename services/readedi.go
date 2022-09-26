package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
)

func ConvertFloat(i string) float64 {
	n, err := strconv.ParseFloat(i, 64)
	if err != nil {
		db := configs.Store
		logData := models.SyncLogger{
			Title:       "edi convert float64",
			Description: fmt.Sprintf("%s", err),
			IsSuccess:   true,
		}
		db.Create(&logData)
		n = 0
	}
	return n
}

func ConvertInt(i string) int64 {
	n, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		db := configs.Store
		logData := models.SyncLogger{
			Title:       "edi convert int64",
			Description: fmt.Sprintf("%s", err),
			IsSuccess:   true,
		}
		db.Create(&logData)
		n = 0
	}
	return n
}

func ReadGediFile(fileEdi *models.FileEdi) {
	file, err := os.Open(fmt.Sprintf("./public/%s", fileEdi.BatchPath))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initailize the DB
	db := configs.Store

	scanner := bufio.NewScanner(file)
	var unitData models.Unit
	db.Select("id,title").First(&unitData, "title=?", fileEdi.Factory.PartUnit)

	var typeData models.PartType
	db.Select("id,title").First(&typeData, "title=?", fileEdi.Factory.PartType)

	if fileEdi.FileType.Title == "R" {
		for scanner.Scan() {
			txt := scanner.Text()
			receiveKey := strings.ReplaceAll((txt[4:(4 + 12)]), " ", "")
			// Check Whs
			var receiveTypeData models.ReceiveType
			db.Preload("Whs").First(&receiveTypeData, "title=?", receiveKey[:3])

			// Initailize Part Data
			PartNo := strings.ReplaceAll(txt[76:(76+25)], " ", "")
			PartName := strings.ReplaceAll(txt[101:(101+25)], " ", "")
			SlugPartNo := strings.ReplaceAll(PartNo, "-", "")

			part := models.Part{
				Slug:        SlugPartNo,
				Title:       PartNo,
				Description: PartName,
				IsActive:    true,
			}

			err := db.FirstOrCreate(&part, &models.Part{Slug: SlugPartNo}).Error
			if err != nil {
				logData := models.SyncLogger{
					Title:       fmt.Sprintf("create master part %s", SlugPartNo),
					Description: fmt.Sprintf("%s", err),
					IsSuccess:   true,
				}
				db.Create(&logData)
			}

			// Initailize Ledger
			ledger := models.Ledger{
				WhsID:      &receiveTypeData.Whs.ID,
				PartID:     &part.ID,
				PartTypeID: &typeData.ID,
				UnitID:     &unitData.ID,
			}
			db.FirstOrCreate(&ledger, &models.Ledger{WhsID: &receiveTypeData.Whs.ID, PartID: &part.ID})

			obj := models.GEDIReceive{
				Factory:          fileEdi.Factory.Title,
				FacZone:          txt[4:(4 + 12)],
				ReceivingKey:     txt[4:(4 + 12)],
				PartNo:           PartNo,
				PartName:         PartName,
				Vendor:           fileEdi.Factory.Title,
				Cd:               fileEdi.Factory.CdCode,
				Unit:             unitData.Title,
				Whs:              fileEdi.Factory.Title,
				Tagrp:            "C",
				RecType:          receiveTypeData.Whs.Value,
				PlanType:         receiveTypeData.Whs.Value,
				RecID:            txt[0:4],
				Aetono:           txt[4:(4 + 12)],
				Aetodt:           txt[16:(16 + 10)],
				Aetctn:           ConvertInt(txt[26:(26 + 9)]),
				Aetfob:           ConvertInt(txt[35:(35 + 9)]),
				Aenewt:           ConvertInt(txt[44:(44 + 11)]),
				Aentun:           txt[55:(55 + 5)],
				Aegrwt:           ConvertInt(txt[60:(60 + 11)]),
				Aegwun:           txt[71:(71 + 5)],
				Aeypat:           txt[76:(76 + 25)],
				Aeedes:           txt[101:(101 + 25)],
				Aetdes:           txt[101:(101 + 25)],
				Aetarf:           0, //ConvertInt(txt[151:(151 + 10)]),
				Aestat:           0, //ConvertInt(txt[161:(161 + 10)]),
				Aebrnd:           0, //ConvertInt(txt[171:(171 + 10)]),
				Aertnt:           0, //ConvertInt(txt[181:(181 + 5)]),
				Aetrty:           0, //ConvertInt(txt[186:(186 + 5)]),
				Aesppm:           0, //ConvertInt(txt[191:(191 + 5)]),
				AeQty1:           0, //ConvertInt(txt[196:(196 + 9)]),
				AeQty2:           0, //ConvertInt(txt[205:(205 + 9)]),
				Aeuntp:           0, //ConvertInt(txt[214:(214 + 9)]),
				Aeamot:           0, //ConvertInt(txt[223:(223 + 11)]),
				Plnctn:           ConvertInt(txt[26:(26 + 9)]),
				PlnQty:           ConvertInt(txt[196:(196 + 9)]),
				Minimum:          0,
				Maximum:          0,
				Picshelfbin:      "PNON",
				Stkshelfbin:      "SNON",
				Ovsshelfbin:      "ONON",
				PicshelfbasicQty: 0,
				Outerpcs:         0,
				AllocateQty:      0,
			}

			// Initailize ReceiveEnt
			dte, _ := time.Parse("02/01/2006", obj.Aetodt)
			receiveEnt := models.Receive{
				FileEdiID:     &fileEdi.ID,
				ReceiveTypeID: &receiveTypeData.ID,
				ReceiveDate:   dte,
				TransferOutNo: receiveKey,
				TexNo:         "-",
				IsSync:        true,
				IsActive:      true,
			}
			db.FirstOrCreate(&receiveEnt, models.Receive{
				TransferOutNo: obj.ReceivingKey,
			})

			receiveDetail := models.ReceiveDetail{
				ReceiveID: &receiveEnt.ID,
				LedgerID:  &ledger.ID,
				PlanQty:   obj.PlnQty,
				PlanCtn:   obj.Plnctn,
				IsActive:  true,
			}

			db.FirstOrCreate(&receiveDetail, models.ReceiveDetail{
				ReceiveID: &receiveEnt.ID,
				LedgerID:  &ledger.ID,
			})

			receiveDetail.PlanQty = obj.PlnQty
			receiveDetail.PlanCtn = obj.Plnctn
			db.Save(&receiveDetail)

			var countReceive []models.ReceiveDetail
			db.Where("receive_id=?", receiveEnt.ID).Find(&countReceive)
			ctn := 0
			for _, x := range countReceive {
				ctn += int(x.PlanCtn)
			}
			receiveEnt.FileEdiID = &fileEdi.ID
			receiveEnt.Item = int64(len(countReceive))
			receiveEnt.PlanCtn = int64(ctn)

			// Disable Sync AW Data
			receiveEnt.IsSync = fileEdi.Factory.IsActive
			db.Save(&receiveEnt)

			fmt.Printf("%s ==> %s unit:%d\n", SlugPartNo, obj.PartName, obj.PlnQty)
		}
	} else {
		// plantype := "ORDERPLAN"
		rnd := 1
		for scanner.Scan() {
			line := scanner.Text()
			Upddte, _ := time.Parse("20060102150405", strings.ReplaceAll(line[141:(141+14)], " ", ""))  //, "%Y%m%d%H%M%S"),
			Updtime, _ := time.Parse("20060102150405", strings.ReplaceAll(line[141:(141+14)], " ", "")) // "%Y%m%d%H%M%S"),
			EtdDte, _ := time.Parse("20060102", strings.ReplaceAll(line[28:(28+8)], " ", ""))
			OrderMonth, _ := time.Parse("20060102", strings.ReplaceAll(line[118:(118+8)], " ", ""))

			obj := models.OrderPlan{
				Seq:              int64(rnd),
				Vendor:           fileEdi.Factory.Title,
				Cd:               fileEdi.Factory.CdCode,
				Sortg1:           fileEdi.Factory.Sortg1,
				Sortg2:           fileEdi.Factory.Sortg2,
				Sortg3:           fileEdi.Factory.Sortg3,
				PlanType:         "ORDERPLAN",
				Pono:             strings.ReplaceAll(line[13:(13+15)], " ", ""),
				RecId:            strings.ReplaceAll(line[0:4], " ", ""),
				Biac:             strings.ReplaceAll(line[5:(5+8)], " ", ""),
				EtdTap:           EtdDte, //), "%Y%m%d"),
				PartNo:           strings.ReplaceAll(line[36:(36+25)], " ", ""),
				PartName:         strings.TrimRight(line[61:(61+25)], " "),
				SampFlg:          strings.ReplaceAll(line[88:(88+1)], " ", ""),
				Orderorgi:        ConvertFloat(line[89:(89 + 9)]),
				Orderround:       ConvertFloat(line[98:(98 + 9)]),
				FirmFlg:          strings.ReplaceAll(line[107:(107+1)], " ", ""),
				ShippedFlg:       strings.ReplaceAll(line[108:(108+1)], " ", ""),
				ShippedQty:       ConvertFloat(line[109:(109 + 9)]),
				Ordermonth:       OrderMonth, //, "%Y%m%d" ),
				BalQty:           ConvertFloat(line[126:(126 + 9)]),
				Bidrfl:           strings.ReplaceAll(line[135:(135+1)], " ", ""),
				DeleteFlg:        strings.ReplaceAll(line[136:(136+1)], " ", ""),
				Reasoncd:         strings.ReplaceAll(line[138:(138+3)], " ", ""),
				Upddte:           Upddte,                                         //, "%Y%m%d%H%M%S"),
				Updtime:          Updtime,                                        // "%Y%m%d%H%M%S"),
				CarrierCode:      strings.ReplaceAll(line[155:(155+4)], " ", ""), //
				Bioabt:           ConvertInt(line[159:(159 + 1)]),
				Bicomd:           strings.ReplaceAll(line[160:(160+1)], " ", ""), //
				Bistdp:           ConvertFloat(line[165:(165 + 9)]),
				Binewt:           ConvertFloat(line[174:(174 + 9)]),
				Bigrwt:           ConvertFloat(line[183:(183 + 9)]),
				Bishpc:           strings.ReplaceAll(line[192:(192+8)], " ", ""), //
				Biivpx:           strings.ReplaceAll(line[200:(200+2)], " ", ""), //
				Bisafn:           strings.ReplaceAll(line[202:(202+6)], " ", ""), //
				Biwidt:           ConvertFloat(line[212:(212 + 4)]),
				Bihigh:           ConvertFloat(line[216:(216 + 4)]),
				Bileng:           ConvertFloat(line[208:(208 + 4)]),
				LotNo:            strings.ReplaceAll(line[220:], " ", ""),
				Minimum:          0,
				Maximum:          0,
				Picshelfbin:      "PNON",
				Stkshelfbin:      "SNON",
				Ovsshelfbin:      "ONON",
				PicshelfbasicQty: 0,
				OuterPcs:         0,
				AllocateQty:      0,
				CreatedAt:        Updtime,
			}

			// Check Whs
			obj.Whs = "COM"
			if obj.Pono[:1] == "#" {
				obj.Whs = "NESC"
			} else if obj.Pono[:1] == "@" {
				obj.Whs = "ICAM"
			}

			var whs models.Whs
			db.Select("id,title").First(&whs, "title=?", obj.Whs)
			var orderZone models.OrderZone
			db.Select("id").
				Where("value=?", obj.Bioabt).
				Where("factory_id=?", fileEdi.Factory.ID).
				Where("whs_id=?", whs.ID).
				First(&orderZone)

			obj.FileEdiID = &fileEdi.ID
			obj.OrderZoneID = &orderZone.ID

			// For Consignee
			affcode := models.Affcode{
				Title:       obj.Biac,
				Description: "-",
			}
			db.FirstOrCreate(&affcode, &models.Affcode{Title: obj.Biac})
			customer := models.Customer{
				Title:       obj.Bishpc,
				Description: obj.Bisafn,
			}
			db.FirstOrCreate(&customer, &models.Customer{Title: obj.Bishpc})

			consigneeData := models.Consignee{
				WhsID:      &whs.ID,             // whs_id
				FactoryID:  &fileEdi.Factory.ID, // factory_id
				AffcodeID:  &affcode.ID,         // affcode_id
				CustomerID: &customer.ID,        // customer_id
				Prefix:     obj.Biivpx,          // prefix
			}
			db.FirstOrCreate(&consigneeData, models.Consignee{
				WhsID:      &whs.ID,             // whs_id
				FactoryID:  &fileEdi.Factory.ID, // factory_id
				AffcodeID:  &affcode.ID,         // affcode_id
				CustomerID: &customer.ID,        // customer_id
				Prefix:     obj.Biivpx,          // prefix
			})
			obj.ConsigneeID = &consigneeData.ID

			/// Revise Type
			obj.ReviseOrderID = nil
			obj.LedgerID = nil
			obj.PcID = nil
			obj.CommercialID = nil
			obj.OrderTypeID = nil
			obj.ShipmentID = nil
			obj.SampleFlgID = nil
			rnd++
		}
	}
}
