package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/abe27/api/configs"
	"github.com/abe27/api/models"
	"github.com/shakinm/xlsReader/xls"
)

func ImportInvoiceTap(fileName *string) {
	workbook, err := xls.OpenFile(*fileName)
	if err != nil {
		return
	}

	sheet, err := workbook.GetSheet(0)
	if err != nil {
		return
	}

	// // sname := sheet.GetName()
	db := configs.Store
	line := 1
	var bRnd int64 = 1
	for i := 0; i <= sheet.GetNumberRows(); i++ {
		if i > 0 {
			if r, err := sheet.GetRow(i); err == nil {
				bhivno, _ := r.GetCol(0)
				bhodpo, _ := r.GetCol(1)
				bhivdt, _ := r.GetCol(2)
				bhconn, _ := r.GetCol(3)
				bhcons, _ := r.GetCol(4)
				bhsven, _ := r.GetCol(5)
				bhshpf, _ := r.GetCol(6)
				bhsafn, _ := r.GetCol(7)
				bhshpt, _ := r.GetCol(8)
				bhfrtn, _ := r.GetCol(9)
				bhcon, _ := r.GetCol(10)
				bhpaln, _ := r.GetCol(11)
				bhpnam, _ := r.GetCol(12)
				bhypat, _ := r.GetCol(13)
				bhctn, _ := r.GetCol(14)
				bhwidt, _ := r.GetCol(15)
				bhleng, _ := r.GetCol(16)
				bhhigh, _ := r.GetCol(17)
				bhgrwt, _ := r.GetCol(18)
				bhcbmt, _ := r.GetCol(19)
				if bhivno.GetString() != "" {
					inv := bhivno.GetString()
					inv_seq, _ := strconv.ParseInt(inv[5:len(inv)-1], 10, 64)
					var facData models.Factory
					db.Select("id,title,inv_prefix,label_prefix").First(&facData, "inv_prefix=?", inv[:2])
					// 	// fmt.Printf("bhivno: %s\n", inv)
					d := bhivdt.GetString()
					if len(d) <= 5 {
						d = fmt.Sprintf("0%s", bhivdt.GetString())
					}
					dd := fmt.Sprintf("20%s%s%s", d[4:6], d[2:4], d[:2])
					etd, _ := time.Parse("20060102", dd)
					// fmt.Printf("%d ==> ETD: %s\n", line, etd)
					var shipment models.Shipment
					db.First(&shipment, "title=?", inv[len(inv)-1:])
					var orderPlan models.OrderPlan
					db.Order("created_at desc,seq desc").Select("id").Where("bisafn=?", bhsafn.GetString()).Where("etd_tap=?", etd.Format("2006-01-02")).Where("part_no=?", bhypat.GetString()).Where("shipment_id=?", shipment.ID).Where("(bal_qty/bistdp)=?", bhctn.GetString()).First(&orderPlan)
					var intCountOrderPlan int64
					db.Order("created_at,seq").Select("id").Where("bisafn=?", bhsafn.GetString()).Where("etd_tap=?", etd.Format("2006-01-02")).Where("part_no=?", bhypat.GetString()).Where("shipment_id=?", shipment.ID).Where("(bal_qty/bistdp)=?", bhctn.GetString()).Find(&models.OrderPlan{}).Count(&intCountOrderPlan)
					Bhcon, _ := strconv.ParseInt(bhcon.GetString(), 10, 64)
					Bhctn, _ := strconv.ParseInt(bhctn.GetString(), 10, 64)
					Bhwidt, _ := strconv.ParseInt(bhwidt.GetString(), 10, 64)
					Bhleng, _ := strconv.ParseInt(bhleng.GetString(), 10, 64)
					Bhhigh, _ := strconv.ParseInt(bhhigh.GetString(), 10, 64)
					Bhgrwt, _ := strconv.ParseFloat(bhgrwt.GetString(), 64)
					Bhcbmt, _ := strconv.ParseFloat(bhcbmt.GetString(), 64)
					var invTap models.ImportInvoiceTap
					invTap.Bhivno = bhivno.GetString()
					invTap.Bhodpo = bhodpo.GetString()
					invTap.Bhivdt = etd
					invTap.Bhconn = bhconn.GetString()
					invTap.Bhcons = bhcons.GetString()
					invTap.Bhsven = bhsven.GetString()
					invTap.Bhshpf = bhshpf.GetString()
					invTap.Bhsafn = bhsafn.GetString()
					invTap.Bhshpt = bhshpt.GetString()
					invTap.Bhfrtn = bhfrtn.GetString()
					invTap.Bhcon = Bhcon
					invTap.Bhpaln = bhpaln.GetString()
					invTap.Bhpnam = bhpnam.GetString()
					invTap.Bhypat = bhypat.GetString()
					invTap.Bhctn = Bhctn
					invTap.Bhwidt = Bhwidt
					invTap.Bhleng = Bhleng
					invTap.Bhhigh = Bhhigh
					invTap.Bhgrwt = Bhgrwt
					invTap.Bhcbmt = Bhcbmt
					invTap.IsMatched = false
					db.FirstOrCreate(&invTap, &models.ImportInvoiceTap{
						Bhivno: bhivno.GetString(),
						Bhodpo: bhodpo.GetString(),
						Bhivdt: etd,
						Bhsafn: bhsafn.GetString(),
						Bhypat: bhypat.GetString(),
						Bhctn:  Bhctn,
					})

					// fmt.Printf("select * from tbt_order_plans where bisafn='%s' and etd_tap='%s' and part_no='%s' and shipment_id='%s' and (bal_qty/bistdp)=%s rows: %d\n", bhsafn.GetString(), etd.Format("2006-01-02"), bhypat.GetString(), shipment.ID, bhctn.GetString(), intCountOrderPlan)
					if orderPlan.ID != "" {
						invTap.OrderPlanID = &orderPlan.ID
						invTap.IsMatched = true
						var orderDetail models.OrderDetail
						db.Select("id,order_id").First(&orderDetail, "order_plan_id=?", orderPlan.ID)
						if orderDetail.ID != "" {
							var order models.Order
							if db.First(&order, "id=?", orderDetail.OrderID).Error == nil {
								// fmt.Printf("OrderID: %s SEQ: %d\n", order.ID, inv_seq)
								order.RunningSeq = inv_seq
								order.IsChecked = true
								db.Save(&order)
							}
							db.Save(&invTap)
							boxDimensions := fmt.Sprintf("%sx%sx%s", bhwidt.GetString(), bhleng.GetString(), bhhigh.GetString())
							var dimData models.PalletType
							err := db.Select("id,type,box_size,pallet_size").First(&dimData, "box_size=?", boxDimensions).Error
							if err != nil {
								db.Select("id,type,box_size,pallet_size").First(&dimData, "box_size=?", "0x0x0")
							}

							txtType := "C"
							if bhpaln.GetString() != "" {
								txtType = "P"
							}

							y := etd.Format("2006-01-02")
							ctnRnd, _ := strconv.ParseInt(bhctn.GetString(), 10, 64)
							for b := 0; b < int(ctnRnd); b++ {
								var lastFticket models.LastFticket
								db.Select("id,last_running").Where("factory_id=?", &facData.ID).Where("on_year=?", y[:4]).First(&lastFticket)
								lastFticket.LastRunning += 1
								lastFticket.OnYear, _ = strconv.ParseInt(y[:4], 10, 64)
								lastFticket.FactoryID = &facData.ID
								lastFticket.IsActive = true
								fticket_no := fmt.Sprintf("%s%s%08d\n", facData.LabelPrefix, y[3:4], (lastFticket.LastRunning))
								db.Save(&lastFticket)
								var pln int64 = 0
								switch txtType {
								case "C":
									pln = bRnd
									bRnd++
								default:
									pln, _ = strconv.ParseInt(bhpaln.GetString(), 10, 64)
								}

								// Create PalletNo/Box
								plData := models.Pallet{
									OrderID:      &order.ID,
									PalletTypeID: &dimData.ID,
									PalletPrefix: txtType,
									PalletNo:     pln,
								}
								db.FirstOrCreate(&plData, &models.Pallet{
									OrderID:      &order.ID,
									PalletTypeID: &dimData.ID,
									PalletPrefix: txtType,
									PalletNo:     pln,
								})
								fmt.Printf("PalletNo.: %s FTicketNo.:%s\n", fticket_no, fmt.Sprintf("%s%03d", txtType, pln))
							}
						}
						// fmt.Printf("Type: %s Pallet No.: %s DIM: %s(%s) PALLET SIZE: %s\n", txtType, bhpaln.GetString(), boxDimensions, dimData.Type, dimData.PalletSize)
					}
					// fmt.Printf("%d ==> bhivno: %s bhodpo: %s bhivdt: %s bhafn: %s bhshpt: %s bhypat: %s bhctn: %s\n", line, bhivno.GetString(), bhodpo.GetString(), d, bhsafn.GetString(), bhshpt.GetString(), bhypat.GetString(), bhctn.GetString())
				}
			}
		}
		line++
	}
}
