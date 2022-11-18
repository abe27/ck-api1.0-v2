package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type PostShippingLabel struct {
	PalletID string `json:"pallet_id" form:"pallet_id"`
	PartID   string `json:"part_id" form:"part_id"`
	Seq      int    `json:"seq" form:"seq"`
}

type PrintShippingLabel struct {
	ID           string    `gorm:"primaryKey;size:21" json:"id"`
	InvoiceNo    string    `gorm:"not null;size:25" json:"invoice_no" form:"invoice_no"`
	OrderNo      string    `gorm:"not null;size:50" json:"order_no" form:"order_no"`
	PartNo       string    `gorm:"not null;size:25" json:"part_no" form:"part_no"`
	Qty          int64     `json:"qty" form:"qty"`
	CustCode     string    `gorm:"not null;size:25" json:"cust_code" form:"cust_code"`
	CustName     string    `gorm:"not null;size:25" json:"cust_name" form:"cust_name"`
	PalletNo     string    `gorm:"not null;size:5" json:"pallet_no" form:"pallet_no"`
	PrintDate    string    `gorm:"not null;size:25;" json:"print_date" form:"print_date"`
	QrCode       string    `json:"qr_code" form:"qr_code"`
	BarCode      string    `gorm:"unique;" json:"bar_code" form:"bar_code"`
	LabelBarCode string    `json:"label_bar_code" form:"label_bar_code"`
	IsPrint      int64     `json:"is_print" form:"is_print" binding:"required" default:"1"`
	CreatedAt    time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (u *PrintShippingLabel) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
