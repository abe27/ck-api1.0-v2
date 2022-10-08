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

	// sname := sheet.GetName()
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
				if bhivno.GetString() != "" {
					inv := bhivno.GetString()
					fmt.Printf("bhivno: %s\n", inv)
					d := bhivdt.GetString()
					if len(d) <= 5 {
						d = fmt.Sprintf("0%s", bhivdt.GetString())
					}
					dd := fmt.Sprintf("20%s%s%s", d[4:6], d[2:4], d[:2])
					etd, _ := time.Parse("20060102", dd)
					fmt.Printf("%d ==> ETD: %s\n", line, etd)
					var shipment models.Shipment
					db.First(&shipment, "title=?", inv[len(inv)-1:])
					var orderPlan models.OrderPlan
					db.Order("created_at,seq").Select("id").Where("bisafn=?", bhsafn.GetString()).Where("etd_tap=?", etd).Where("part_no=?", bhypat.GetString()).Where("shipment_id=?", shipment.ID).Where("(bal_qty/bistdp)=?", bhctn.GetString()).Last(&orderPlan)
					if orderPlan.ID != "" {
						fmt.Printf("ID: %s\n", orderPlan.ID)
					}

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
					if orderPlan.ID != "" {
						invTap.OrderPlanID = &orderPlan.ID
						invTap.IsMatched = true
					}

					db.FirstOrCreate(&invTap, &models.ImportInvoiceTap{
						Bhivno: bhivno.GetString(),
						Bhodpo: bhodpo.GetString(),
						Bhivdt: etd,
						Bhsafn: bhsafn.GetString(),
						Bhypat: bhypat.GetString(),
						Bhctn:  Bhctn,
					})
					// fmt.Printf(" bhivno: %s bhodpo: %s bhivdt: %s bhconn: %s bhconns: %s bhsven: %s bhshpf: %s bhafn: %s bhshpt: %s bhfrtn: %s bhcon: %s bhpaln: %s bhpnam: %s bhypat: %s bhctn: %s bhwidt: %s bhleng: %s bhhigh: %s bhgrwt: %s bhcbmt: %s\n", bhivno.GetString(), bhodpo.GetString(), d, bhconn.GetString(), bhconns.GetString(), bhsven.GetString(), bhshpf.GetString(), bhafn.GetString(), bhshpt.GetString(), bhfrtn.GetString(), bhcon.GetString(), bhpaln.GetString(), bhpnam.GetString(), bhypat.GetString(), bhctn.GetString(), bhwidt.GetString(), bhleng.GetString(), bhhigh.GetString(), bhgrwt.GetString(), bhcbmt.GetString())
				}
			}
		}
		line++
	}
}
