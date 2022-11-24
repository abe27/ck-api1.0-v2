package services

import (
	"fmt"
	"strconv"
	"strings"
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

	for _, sheet := range workbook.GetSheets() {
		// // sname := sheet.GetName()
		db := configs.Store
		line := 1
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
					// fmt.Printf("Start inv: %s\n", bhivno.GetString())
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
						if err := db.Order("created_at desc,seq desc").Select("id,bal_qty,bistdp").Where("bisafn like ?", "%"+bhsafn.GetString()+"%").Where("etd_tap=?", etd.Format("2006-01-02")).Where("part_no=?", bhypat.GetString()).Where("shipment_id=?", shipment.ID).Where("pono in ?", []string{strings.Trim(bhodpo.GetString(), ""), strings.Trim(strings.ReplaceAll(bhodpo.GetString(), " ", ""), "")}).First(&orderPlan).Error; err != nil {
							print(err.Error())
						}

						Bhcon, _ := strconv.ParseInt(bhcon.GetString(), 10, 64)
						Bhctn, _ := strconv.ParseInt(bhctn.GetString(), 10, 64)
						Bhwidt, _ := strconv.ParseInt(bhwidt.GetString(), 10, 64)
						Bhleng, _ := strconv.ParseInt(bhleng.GetString(), 10, 64)
						Bhhigh, _ := strconv.ParseInt(bhhigh.GetString(), 10, 64)
						Bhgrwt, _ := strconv.ParseFloat(bhgrwt.GetString(), 64)
						Bhcbmt, _ := strconv.ParseFloat(bhcbmt.GetString(), 64)
						var invTap models.ImportInvoiceTap
						invTap.Biseq = int64(line)
						invTap.Bhivno = bhivno.GetString()
						invTap.Bhodpo = bhodpo.GetString()
						invTap.BhodpoTrim = strings.Trim(strings.ReplaceAll(bhodpo.GetString(), " ", ""), "")
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
							Bhivno:     bhivno.GetString(),
							Bhodpo:     bhodpo.GetString(),
							BhodpoTrim: strings.Trim(strings.ReplaceAll(bhodpo.GetString(), " ", ""), ""),
							Bhivdt:     etd,
							Bhsafn:     bhsafn.GetString(),
							Bhypat:     bhypat.GetString(),
							Bhctn:      Bhctn,
						})

						// fmt.Printf("select * from tbt_order_plans where bisafn='%s' and etd_tap='%s' and part_no='%s' and shipment_id='%s' and (bal_qty/bistdp)=%s rows: %d\n", bhsafn.GetString(), etd.Format("2006-01-02"), bhypat.GetString(), shipment.ID, bhctn.GetString(), intCountOrderPlan)
						if orderPlan.ID != "" {
							invTap.OrderPlanID = &orderPlan.ID
							invTap.IsMatched = true
							var orderDetail models.OrderDetail
							db.Select("id,order_id,order_ctn,total_on_pallet").First(&orderDetail, "order_plan_id=?", orderPlan.ID)
							if orderDetail.ID != "" {
								var order models.Order
								if err := db.First(&order, "id=?", orderDetail.OrderID).Error; err == nil {
									// fmt.Printf("OrderID: %s SEQ: %d\n", order.ID, inv_seq)
									order.RunningSeq = inv_seq
									order.IsChecked = true
									order.IsInvoice = false
									order.IsSync = true
									db.Save(&order)
								}
								db.Save(&invTap)
								var dimData models.PalletType
								err := db.
									Select("id,type,box_size_width,box_size_length,box_size_hight,pallet_size_width,pallet_size_length,pallet_size_hight").
									Where("box_size_width=?", bhwidt.GetString()).
									Where("box_size_length=?", bhleng.GetString()).
									Where("box_size_hight=?", bhhigh.GetString()).
									First(&dimData).Error
								if err != nil {
									db.
										Select("id,type,box_size_width,box_size_length,box_size_hight,pallet_size_width,pallet_size_length,pallet_size_hight").
										Where("box_size_width=?", "0").
										Where("box_size_length=?", "0").
										Where("box_size_hight=?", "0").
										First(&dimData)
								}

								txtType := "C"
								if bhpaln.GetString() != "" {
									txtType = "P"
								}

								y := etd.Format("2006-01-02")
								ctnRnd, _ := strconv.ParseInt(bhctn.GetString(), 10, 64)
								// ctnRnd := orderPlan.BalQty / orderPlan.Bistdp
								for b := 0; b < int(ctnRnd); b++ {
									var pln int64 = 0
									switch txtType {
									case "C":
										var bRnd int64
										db.Where("order_id=?", &order.ID).Where("pallet_prefix=?", "C").Find(&models.Pallet{}).Count(&bRnd)
										pln = bRnd + 1
									default:
										pln, _ = strconv.ParseInt(bhpaln.GetString(), 10, 64)
									}

									// Create PalletNo/Box
									plData := models.Pallet{
										OrderID:      &order.ID,
										PalletTypeID: &dimData.ID,
										PalletPrefix: txtType,
										PalletNo:     pln,
										IsActive:     true,
									}
									err := db.FirstOrCreate(&plData, &models.Pallet{
										OrderID:      &order.ID,
										PalletPrefix: txtType,
										PalletNo:     pln,
									}).Error
									if err != nil {
										var sysLog models.SyncLogger
										sysLog.Title = fmt.Sprintf("Can not create pln: %d", pln)
										sysLog.Description = fmt.Sprintf("Error %s", err.Error())
										sysLog.IsSuccess = false
										db.Create(&sysLog)
									}
									var checkPlDuplicate int64
									db.Select("id").Where("pallet_id=?", &plData.ID).Where("order_detail_id=?", &orderDetail.ID).Find(&models.PalletDetail{}).Count(&checkPlDuplicate)
									if checkPlDuplicate < orderDetail.OrderCtn {
										var lastFticket models.LastFticket
										db.Select("id,last_running").Where("factory_id=?", &facData.ID).Where("on_year=?", y[:4]).First(&lastFticket)
										seqNo := (lastFticket.LastRunning + 1)
										// fmt.Printf("%s:  %d != %d SEQ: %d PLID: %s\n", orderDetail.ID, ctnRnd, checkPlDuplicate, seqNo, plData.ID)
										//Create PlletDetails
										plDetailData := models.PalletDetail{
											PalletID:      &plData.ID,
											OrderDetailID: &orderDetail.ID,
											SeqNo:         seqNo,
											IsActive:      true,
										}

										if db.Create(&plDetailData).Error == nil {
											lastFticket.LastRunning = (seqNo + 1)
											lastFticket.OnYear, _ = strconv.ParseInt(y[:4], 10, 64)
											lastFticket.FactoryID = &facData.ID
											lastFticket.IsActive = true
											db.Save(&lastFticket)
										}
									}
								}
								// Update Status OrderDetail
								db.Model(&orderDetail).Select("total_on_pallet", "order_ctn", "is_matched", "is_checked", "is_sync").Updates(models.OrderDetail{TotalOnPallet: orderDetail.TotalOnPallet + int64(ctnRnd), OrderCtn: int64(orderPlan.BalQty) / int64(orderPlan.Bistdp), IsMatched: true, IsChecked: true, IsSync: true})
							}
						}
					}
				}
			}
			line++
		}
	}
}
