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

	IsSync := false
	if fileEdi.Factory.Title == "AW" {
		IsSync = true
	}

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
			receiveEnt.Item = int64(len(countReceive))
			receiveEnt.PlanCtn = int64(ctn)

			// Disable Sync AW Data
			receiveEnt.IsSync = IsSync
			db.Save(&receiveEnt)

			fmt.Printf("%s ==> %s unit:%d\n", SlugPartNo, obj.PartName, obj.PlnQty)
		}
	} else {
		// plantype := "ORDERPLAN"
		fmt.Printf("for order plan\n")
	}
}
