package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type OrderPlan struct {
	ID               string      `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	OrderId          string      `gorm:"not null;unique;size:21;index;,omitempty" form:"order_id" json:"order_id"`
	FileEdiID        *string     `gorm:"not null;" json:"file_edi_id,omitempty" form:"file_edi_id"`
	WhsID            *string     `gorm:"not null;" json:"whs_id,omitempty" form:"whs_id"`
	OrderZoneID      *string     `gorm:"not null;" json:"order_zone_id,omitempty" form:"order_type_id" binding:"required"`
	ConsigneeID      *string     `gorm:"not null;" json:"consignee_id,omitempty" form:"consignee_id"`
	ReviseOrderID    *string     `gorm:"null;" json:"revise_order_id,omitempty" form:"revise_order_id" binding:"required"`
	LedgerID         *string     `gorm:"not null;" json:"ledger_id,omitempty" form:"ledger_id" binding:"required"`
	PcID             *string     `gorm:"not null;" json:"pc_id,omitempty" form:"pc_id" binding:"required"`
	CommercialID     *string     `gorm:"not null;" json:"commercial_id,omitempty" form:"commercial_id" binding:"required"`
	OrderTypeID      *string     `gorm:"not null;" json:"order_type_id,omitempty" form:"order_type_id" binding:"required"`
	ShipmentID       *string     `gorm:"not null;" json:"shipment_id,omitempty" form:"shipment_id" binding:"required"`
	SampleFlgID      *string     `gorm:"not null;" json:"sample_flg_id,omitempty" form:"sample_flg_id" binding:"required"`
	Seq              int64       `form:"seq" json:"seq,omitempty"`
	Vendor           string      `gorm:"size:5;" form:"vendor" json:"vendor,omitempty"`
	Cd               string      `gorm:"size:5;" form:"cd" json:"cd,omitempty"`
	Tagrp            string      `gorm:"size:5;" form:"tagrp" json:"tagrp,omitempty"`
	Sortg1           string      `gorm:"size:25" form:"sortg1" json:"sortg1,omitempty"`
	Sortg2           string      `gorm:"size:25" form:"sortg2" json:"sortg2,omitempty"`
	Sortg3           string      `gorm:"size:25" form:"sortg3" json:"sortg3,omitempty"`
	PlanType         string      `gorm:"size:25" form:"plan_type" json:"plan_type,omitempty"`
	OrderGroup       string      `gorm:"size:25" form:"order_group" json:"order_groups,omitempty"`
	Pono             string      `gorm:"size:25" form:"pono" json:"pono,omitempty"`
	RecId            string      `gorm:"size:25" form:"rec_id" json:"rec_id,omitempty"`
	Biac             string      `gorm:"size:25" form:"biac" json:"biac,omitempty"`
	EtdTap           time.Time   `gorm:"type:date;" form:"etd_tap" json:"etd_tap"`
	PartNo           string      `gorm:"size:25" form:"part_no" json:"part_no,omitempty"`
	PartName         string      `gorm:"size:50" form:"part_name" json:"part_name"`
	SampFlg          string      `gorm:"column:sample_flg;size:2,omitempty" form:"sample_flg" json:"sample_flg"`
	Orderorgi        float64     `form:"orderorgi" json:"orderorgi,omitempty"`
	Orderround       float64     `form:"orderround" json:"orderround,omitempty"`
	FirmFlg          string      `gorm:"size:2" form:"firm_flg" json:"firm_flg,omitempty"`
	ShippedFlg       string      `gorm:"size:2" form:"shipped_flg" json:"shipped_flg,omitempty"`
	ShippedQty       float64     `form:"shipped_qty" json:"shipped_qty,omitempty"`
	Ordermonth       time.Time   `gorm:"type:date;" form:"ordermonth" json:"ordermonth,omitempty"`
	BalQty           float64     `form:"balqty" json:"balqty,omitempty"`
	Bidrfl           string      `gorm:"size:2" form:"bidrfl" json:"bidrfl,omitempty"`
	DeleteFlg        string      `gorm:"size:2" form:"delete_flg" json:"delete_flg,omitempty"`
	Reasoncd         string      `gorm:"size:5" orm:"reasoncd" json:"reasoncd,omitempty"`
	Upddte           time.Time   `gorm:"type:date;" form:"upddte" json:"upddte,omitempty"`
	Updtime          time.Time   `gorm:"type:Time;" form:"updtime" json:"updtime,omitempty"`
	CarrierCode      string      `gorm:"size:5" form:"carrier_code" json:"carrier_code,omitempty"`
	Bioabt           int64       `form:"bioat" json:"bioat,omitempty"`
	Bicomd           string      `gorm:"size:2" form:"bicomd" json:"bicomd,omitempty"`
	Bistdp           float64     `form:"bistdp" json:"bistdp,omitempty"`
	Binewt           float64     `form:"binewt" json:"binewt,omitempty"`
	Bigrwt           float64     `form:"bigrwt" json:"bigrwt,omitempty"`
	Bishpc           string      `gorm:"size:25" form:"bishpc" json:"bishpc,omitempty"`
	Biivpx           string      `gorm:"size:5" form:"biivpx" json:"biivpx,omitempty"`
	Bisafn           string      `gorm:"size:25" form:"bisafn" json:"bisafn,omitempty"`
	Biwidt           float64     `form:"biwidt" json:"biwidt,omitempty"`
	Bihigh           float64     `form:"bihigh" json:"bihigh,omitempty"`
	Bileng           float64     `form:"bileng" json:"bileng,omitempty"`
	LotNo            string      `gorm:"size:25" form:"lotno" json:"lotno,omitempty"`
	Minimum          int64       `form:"minimum" json:"minimum,omitempty"`
	Maximum          int64       `form:"maximum" json:"maximum,omitempty"`
	Picshelfbin      string      `gorm:"size:25" form:"picshelfbin" json:"picshelfbin,omitempty"`
	Stkshelfbin      string      `gorm:"size:25" form:"stkshelfbin" json:"stkshelfbin,omitempty"`
	Ovsshelfbin      string      `gorm:"size:25" form:"ovsshelfbin" json:"ovsshelfbin,omitempty"`
	PicshelfbasicQty float64     `form:"picshelfbasicqty" json:"picshelfbasicqty,omitempty"`
	OuterPcs         float64     `form:"outerpcs" json:"outerpcs,omitempty"`
	AllocateQty      float64     `json:"allocate_qty,omitempty" form:"allocate_qty"`
	Description      string      `json:"description,omitempty" form:"description"`
	IsReviseError    bool        `json:"is_revise_error,omitempty" form:"is_revise_error" default:"false"`
	IsGenerate       bool        `json:"is_generate,omitempty" form:"is_generate" default:"false"`
	ByManually       bool        `json:"by_manually,omitempty" form:"by_manually" default:"false"`
	IsSync           bool        `json:"is_sync,omitempty" form:"is_sync" default:"false"`
	IsActive         bool        `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt        time.Time   `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt        time.Time   `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	FileEdi          FileEdi     `gorm:"foreignKey:FileEdiID;references:ID;" json:"file_gedi,omitempty"`
	Whs              Whs         `gorm:"foreignKey:WhsID;references:ID;" json:"whs,omitempty"`
	Consignee        Consignee   `gorm:"foreignKey:ConsigneeID;references:ID" json:"consignee,omitempty"`
	ReviseOrder      ReviseOrder `gorm:"foreignKey:ReviseOrderID;references:ID" json:"reviseOrder,omitempty"`
	Ledger           Ledger      `gorm:"foreignKey:LedgerID;references:ID" json:"ledger,omitempty"`
	Pc               Pc          `gorm:"foreignKey:PcID;references:ID" json:"pc,omitempty"`
	Commercial       Commercial  `gorm:"foreignKey:CommercialID;references:ID" json:"commercial,omitempty"`
	OrderType        OrderType   `gorm:"foreignKey:OrderTypeID;references:ID" json:"orderType,omitempty"`
	Shipment         Shipment    `gorm:"foreignKey:ShipmentID;references:ID" json:"shipment,omitempty"`
	OrderZone        OrderZone   `gorm:"foreignKey:OrderZoneID;references:ID" json:"orderzone,omitempty"`
	SampleFlg        SampleFlg   `gorm:"foreignKey:SampleFlgID;references:ID" json:"sampleflg,omitempty"`
}

func (u *OrderPlan) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	OrderId, _ := g.New()
	u.OrderId = OrderId
	return nil
}

// type OrderPlanManually struct {
// 	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
// 	OrderPlanID *string   `gorm:"not null" json:"order_plan_id,omitempty" form:"order_plan_id"`
// 	UserID      *string   `gorm:"not null" json:"user_id,omitempty" form:"user_id"`
// 	Description string    `json:"description,omitempty" form:"description"`
// 	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
// 	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
// 	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
// 	OrderPlan   OrderPlan `gorm:"foreignKey:OrderPlanID;references:ID" json:"order_plan"`
// 	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
// }

// func (u *OrderPlanManually) BeforeCreate(tx *gorm.DB) (err error) {
// 	id, _ := g.New()
// 	u.ID = id
// 	return nil
// }

// type PermissionPlanManually struct {
// 	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
// 	UserID      *string   `gorm:"not null" json:"user_id,omitempty" form:"user_id"`
// 	Description string    `json:"description,omitempty" form:"description"`
// 	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
// 	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
// 	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
// 	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
// }

// func (u *PermissionPlanManually) BeforeCreate(tx *gorm.DB) (err error) {
// 	id, _ := g.New()
// 	u.ID = id
// 	return nil
// }
