package services

import (
	"fmt"

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
				// bhodpo, _ := r.GetCol(1)
				bhivdt, _ := r.GetCol(2)
				// bhconn, _ := r.GetCol(3)
				// bhconns, _ := r.GetCol(4)
				// bhsven, _ := r.GetCol(5)
				// bhshpf, _ := r.GetCol(6)
				bhsafn, _ := r.GetCol(7)
				// bhshpt, _ := r.GetCol(8)
				// bhfrtn, _ := r.GetCol(9)
				// bhcon, _ := r.GetCol(10)
				// bhpaln, _ := r.GetCol(11)
				// bhpnam, _ := r.GetCol(12)
				bhypat, _ := r.GetCol(13)
				// bhctn, _ := r.GetCol(14)
				// bhwidt, _ := r.GetCol(15)
				// bhleng, _ := r.GetCol(16)
				// bhhigh, _ := r.GetCol(17)
				// bhgrwt, _ := r.GetCol(18)
				// bhcbmt, _ := r.GetCol(19)
				if bhivno.GetString() != "" {
					fmt.Printf("bhivno: %s\n", bhivno.GetString())
					d := bhivdt.GetString()
					if len(d) <= 5 {
						d = fmt.Sprintf("0%s", bhivdt.GetString())
					}
					etd := fmt.Sprintf("20%s-%s-%s", d[4:6], d[2:4], d[:2])
					fmt.Printf("%d ==> ETD: %s\n", line, etd)
					var orderPlan models.OrderPlan
					db.Order("created_at,seq").Where("bisafn=?", bhsafn.GetString()).Where("etd_tap=?", etd).Where("part_no=?", bhypat.GetString()).Last(&orderPlan)
					fmt.Printf("==> %d\n", orderPlan.Seq)
					// fmt.Printf(" bhivno: %s bhodpo: %s bhivdt: %s bhconn: %s bhconns: %s bhsven: %s bhshpf: %s bhafn: %s bhshpt: %s bhfrtn: %s bhcon: %s bhpaln: %s bhpnam: %s bhypat: %s bhctn: %s bhwidt: %s bhleng: %s bhhigh: %s bhgrwt: %s bhcbmt: %s\n", bhivno.GetString(), bhodpo.GetString(), d, bhconn.GetString(), bhconns.GetString(), bhsven.GetString(), bhshpf.GetString(), bhafn.GetString(), bhshpt.GetString(), bhfrtn.GetString(), bhcon.GetString(), bhpaln.GetString(), bhpnam.GetString(), bhypat.GetString(), bhctn.GetString(), bhwidt.GetString(), bhleng.GetString(), bhhigh.GetString(), bhgrwt.GetString(), bhcbmt.GetString())
				}
			}
		}
		line++
	}
}
