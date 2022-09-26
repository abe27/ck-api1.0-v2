package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type OrderPlan struct {
	ID               string      `gorm:"primaryKey;size:21;" json:"id"`
	FileEdiID        *string     `json:"file_edi_id" form:"file_edi_id"`
	WhsID            *string     `json:"whs_id" form:"whs_id"`
	OrderZoneID      *string     `json:"order_zone_id" form:"order_type_id" binding:"required"`
	ConsigneeID      *string     `json:"consignee_id" form:"consignee_id"`
	ReviseOrderID    *string     `json:"revise_order_id" form:"revise_order_id" binding:"required"`
	LedgerID         *string     `json:"ledger_id" form:"ledger_id" binding:"required"`
	PcID             *string     `json:"pc_id" form:"pc_id" binding:"required"`
	CommercialID     *string     `json:"commercial_id" form:"commercial_id" binding:"required"`
	OrderTypeID      *string     `json:"order_type_id" form:"order_type_id" binding:"required"`
	ShipmentID       *string     `json:"shipment_id" form:"shipment_id" binding:"required"`
	SampleFlgID      *string     `json:"sample_flg_id" form:"sample_flg_id" binding:"required"`
	Seq              int64       `form:"seq" json:"seq"`
	Vendor           string      `gorm:"size:5" form:"vendor" json:"vendor"`
	Cd               string      `gorm:"size:5" form:"cd" json:"cd"`
	Tagrp            string      `gorm:"size:5" form:"tagrp" json:"tagrp"`
	Sortg1           string      `gorm:"size:25" form:"sortg1" json:"sortg1"`
	Sortg2           string      `gorm:"size:25" form:"sortg2" json:"sortg2"`
	Sortg3           string      `gorm:"size:25" form:"sortg3" json:"sortg3"`
	PlanType         string      `gorm:"size:25" form:"plan_type" json:"plan_type"`
	OrderGroup       string      `gorm:"size:25" form:"order_group" json:"order_groups"`
	Pono             string      `gorm:"size:25" form:"pono" json:"pono"`
	RecId            string      `gorm:"size:25" form:"rec_id" json:"rec_id"`
	Biac             string      `gorm:"size:25" form:"biac" json:"biac"`
	EtdTap           time.Time   `gorm:"type:date;" form:"etd_tap" json:"etd_tap"`
	PartNo           string      `gorm:"size:25" form:"part_no" json:"part_no"`
	PartName         string      `gorm:"size:50" form:"part_name" json:"part_name"`
	SampFlg          string      `gorm:"column:'sample_flg';size:2" form:"sample_flg" json:"sample_flg"`
	Orderorgi        float64     `form:"orderorgi" json:"orderorgi"`
	Orderround       float64     `form:"orderround" json:"orderround"`
	FirmFlg          string      `gorm:"size:2" form:"firm_flg" json:"firm_flg"`
	ShippedFlg       string      `gorm:"size:2" form:"shipped_flg" json:"shipped_flg"`
	ShippedQty       float64     `form:"shipped_qty" json:"shipped_qty"`
	Ordermonth       time.Time   `gorm:"type:date;" form:"ordermonth" json:"ordermonth"`
	BalQty           float64     `form:"balqty" json:"balqty"`
	Bidrfl           string      `gorm:"size:2" form:"bidrfl" json:"bidrfl"`
	DeleteFlg        string      `gorm:"size:2" form:"delete_flg" json:"delete_flg"`
	Reasoncd         string      `gorm:"size:5" form:"reasoncd" json:"reasoncd"`
	Upddte           time.Time   `gorm:"type:date;" form:"upddte" json:"upddte"`
	Updtime          time.Time   `gorm:"type:Time;" form:"updtime" json:"updtime"`
	CarrierCode      string      `gorm:"size:5" form:"carrier_code" json:"carrier_code"`
	Bioabt           int64       `form:"bioat" json:"bioat"`
	Bicomd           string      `gorm:"size:2" form:"bicomd" json:"bicomd"`
	Bistdp           float64     `form:"bistdp" json:"bistdp"`
	Binewt           float64     `form:"binewt" json:"binewt"`
	Bigrwt           float64     `form:"bigrwt" json:"bigrwt"`
	Bishpc           string      `gorm:"size:25" form:"bishpc" json:"bishpc"`
	Biivpx           string      `gorm:"size:5" form:"biivpx" json:"biivpx"`
	Bisafn           string      `gorm:"size:25" form:"bisafn" json:"bisafn"`
	Biwidt           float64     `form:"biwidt" json:"biwidt"`
	Bihigh           float64     `form:"bihigh" json:"bihigh"`
	Bileng           float64     `form:"bileng" json:"bileng"`
	LotNo            string      `gorm:"size:25" form:"lotno" json:"lotno"`
	Minimum          int64       `form:"minimum" json:"minimum"`
	Maximum          int64       `form:"maximum" json:"maximum"`
	Picshelfbin      string      `gorm:"size:25" form:"picshelfbin" json:"picshelfbin"`
	Stkshelfbin      string      `gorm:"size:25" form:"stkshelfbin" json:"stkshelfbin"`
	Ovsshelfbin      string      `gorm:"size:25" form:"ovsshelfbin" json:"ovsshelfbin"`
	PicshelfbasicQty float64     `form:"picshelfbasicqty" json:"picshelfbasicqty"`
	OuterPcs         float64     `form:"outerpcs" json:"outerpcs"`
	AllocateQty      float64     `form:"allocate_qty"`
	Description      string      `form:"description" json:"description"`
	IsReviseError    bool        `json:"is_revise_error" form:"is_revise_error" default:"false"`
	IsGenerate       bool        `json:"is_generate" form:"is_generate" default:"false"`
	ByManually       bool        `json:"by_manually" form:"by_manually" default:"false"`
	IsSync           bool        `json:"is_sync" form:"is_sync" default:"false"`
	IsActive         bool        `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt        time.Time   `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt        time.Time   `json:"updated_at" form:"updated_at" default:"now"`
	FileEdi          FileEdi     `gorm:"foreignKey:FileEdiID;references:ID;" json:"file_gedi"`
	Whs              Whs         `gorm:"foreignKey:WhsID;references:ID;" json:"whs"`
	Consignee        Consignee   `gorm:"foreignKey:ConsigneeID;references:ID" json:"consignee"`
	ReviseOrder      ReviseOrder `gorm:"foreignKey:ReviseOrderID;references:ID" json:"reviseOrder"`
	Ledger           Ledger      `gorm:"foreignKey:LedgerID;references:ID" json:"ledger"`
	Pc               Pc          `gorm:"foreignKey:PcID;references:ID" json:"pc"`
	Commercial       Commercial  `gorm:"foreignKey:CommercialID;references:ID" json:"commercial"`
	OrderType        OrderType   `gorm:"foreignKey:OrderTypeID;references:ID" json:"orderType"`
	Shipment         Shipment    `gorm:"foreignKey:ShipmentID;references:ID" json:"shipment"`
	OrderZone        OrderZone   `gorm:"foreignKey:OrderZoneID;references:ID" json:"orderzone"`
	SampleFlg        SampleFlg   `gorm:"foreignKey:SampleFlgID;references:ID" json:"sampleflg"`
}

func (u *OrderPlan) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

// type OrderPlanManually struct {
// 	ID          string    `gorm:"primaryKey;size:21" json:"id"`
// 	OrderPlanID *string   `gorm:"not null" json:"order_plan_id" form:"order_plan_id"`
// 	UserID      *string   `gorm:"not null" json:"user_id" form:"user_id"`
// 	Description string    `json:"description" form:"description"`
// 	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
// 	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
// 	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
// 	OrderPlan   OrderPlan `gorm:"foreignKey:OrderPlanID;references:ID" json:"order_plan"`
// 	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
// }

// func (u *OrderPlanManually) BeforeCreate(tx *gorm.DB) (err error) {
// 	id, _ := g.New()
// 	u.ID = id
// 	return nil
// }

// type PermissionPlanManually struct {
// 	ID          string    `gorm:"primaryKey;size:21" json:"id"`
// 	UserID      *string   `gorm:"not null" json:"user_id" form:"user_id"`
// 	Description string    `json:"description" form:"description"`
// 	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
// 	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
// 	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
// 	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
// }

// func (u *PermissionPlanManually) BeforeCreate(tx *gorm.DB) (err error) {
// 	id, _ := g.New()
// 	u.ID = id
// 	return nil
// }
