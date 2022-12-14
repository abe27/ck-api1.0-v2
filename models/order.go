package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type OrderTitle struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" default:"now"`
}

func (obj *OrderTitle) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Order struct {
	ID           string         `gorm:"primaryKey;unique;index;size:21" json:"id"`
	RowID        string         `gorm:"null;size:18" json:"row_id" form:"row_id"`
	ConsigneeID  *string        `gorm:"not null" json:"consignee_id" form:"consignee_id" binding:"required"`
	ShipmentID   *string        `gorm:"not null" json:"shipment_id" form:"shipment_id" binding:"required"`
	EtdDate      *time.Time     `gorm:"not null;type:date;" json:"etd_date" form:"etd_date" binding:"required"`
	PcID         *string        `gorm:"not null" json:"pc_id" form:"pc_id" binding:"required"`
	CommercialID *string        `gorm:"not null" json:"commercial_id" form:"commercial_id" binding:"required"`
	SampleFlgID  *string        `gorm:"not null" json:"sample_flg_id" form:"sample_flg_id" binding:"required"`
	OrderTitleID *string        `gorm:"not null" json:"title_id" form:"title_id" binding:"required"`
	Bioabt       int64          `json:"bioabt" form:"bioabt" binding:"required"`
	CarrierCode  string         `gorm:"size:255;" json:"carrier_code" form:"carrier_code" binding:"required"`
	ShipForm     string         `gorm:"size:255;" json:"ship_form" form:"ship_form" binding:"required"`
	ShipTo       string         `gorm:"size:255;" json:"ship_to" form:"ship_to" binding:"required"`
	ShipVia      string         `gorm:"size:255;" json:"ship_via" form:"ship_via" binding:"required"`
	ShipDer      string         `gorm:"size:255;" json:"ship_der" form:"ship_der" binding:"required"`
	LoadingArea  string         `gorm:"size:255;" json:"loading_area" form:"loading_area" binding:"required"`
	Privilege    string         `gorm:"size:50;" json:"privilege" form:"privilege" binding:"required"`
	ZoneCode     string         `gorm:"size:10;unique;" json:"zone_code" form:"zone_code" binding:"required"`
	RunningSeq   int64          `json:"running_seq" form:"running_seq" binding:"required"`
	IsChecked    bool           `json:"is_checked" form:"is_checked"`
	IsInvoice    bool           `json:"is_invoice" form:"is_invoice"`
	IsSync       bool           `json:"is_sync" form:"is_sync"`
	IsActive     bool           `json:"is_active" form:"is_active"`
	CreatedAt    time.Time      `json:"created_at" default:"now"`
	UpdatedAt    time.Time      `json:"updated_at" default:"now"`
	Consignee    Consignee      `gorm:"foreignKey:ConsigneeID;references:ID" json:"consignee"`
	Shipment     Shipment       `gorm:"foreignKey:ShipmentID;references:ID" json:"shipment"`
	Pc           Pc             `gorm:"foreignKey:PcID;references:ID" json:"pc"`
	Commercial   Commercial     `gorm:"foreignKey:CommercialID;references:ID" json:"commercial"`
	SampleFlg    SampleFlg      `gorm:"foreignKey:SampleFlgID;references:ID" json:"sample_flg"`
	OrderTitle   OrderTitle     `gorm:"foreignKey:OrderTitleID;references:ID" json:"order_title"`
	OrderDetail  []*OrderDetail `json:"order_detail"`
	Pallet       []*Pallet      `json:"pallet"`
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}

type OrderDetailForm struct {
	ReviseID      string  `json:"revise_id" form:"revise_id"`
	OrderID       *string `json:"order_id" form:"order_id" binding:"required"`
	Pono          *string `json:"pono" form:"pono" binding:"required"`
	LedgerID      *string `json:"ledger_id" form:"ledger_id" binding:"required"`
	OrderPlanID   *string `json:"order_plan_id" form:"order_plan_id" binding:"required"`
	OrderCtn      int64   `json:"order_ctn" form:"order_ctn" binding:"required"`
	TotalOnPallet int64   `json:"total_on_pallet" form:"total_on_pallet" binding:"required"`
}

type AddOrderDetailForm struct {
	Pono        string `json:"pono" form:"pono" binding:"required"`
	PartNo      string `json:"part_no" form:"part_no" binding:"required"`
	OrderBalQty int64  `json:"order_balqty" form:"order_balqty" binding:"required"`
	OrderCtn    int64  `json:"order_ctn" form:"order_ctn" binding:"required"`
}

type OrderDetail struct {
	ID            string    `gorm:"primaryKey;unique;index;size:21" json:"id"`
	RowID         string    `gorm:"null;size:18" json:"row_id" form:"row_id"`
	OrderID       *string   `gorm:"not null;" json:"order_id" form:"order_id" binding:"required"`
	Pono          *string   `gorm:"not null;size:25" json:"pono" form:"pono" binding:"required"`
	LedgerID      *string   `gorm:"not null;" json:"ledger_id" form:"ledger_id" binding:"required"`
	OrderPlanID   *string   `gorm:"not null;unique;" json:"order_plan_id" form:"order_plan_id" binding:"required"`
	OrderCtn      int64     `json:"order_ctn" form:"order_ctn" binding:"required"`
	TotalOnPallet int64     `json:"total_on_pallet" form:"total_on_pallet" binding:"required"`
	IsChecked     bool      `json:"is_checked" form:"is_checked" default:"false"`
	IsMatched     bool      `json:"is_matched" form:"is_matched" default:"false"`
	IsSync        bool      `json:"is_sync" form:"is_sync" default:"false"`
	IsActive      bool      `json:"is_active" form:"is_active"`
	CreatedAt     time.Time `json:"created_at" default:"now"`
	UpdatedAt     time.Time `json:"updated_at" default:"now"`
	Order         Order     `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	Ledger        Ledger    `gorm:"foreignKey:LedgerID;references:ID" json:"ledger"`
	OrderPlan     OrderPlan `gorm:"foreignKey:OrderPlanID;references:ID" json:"orderplan"`
}

func (u *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}

type OrderPrepare struct {
	ID             string       `gorm:"primaryKey;unique;index;size:21" json:"id"`
	PalletDetailID *string      `gorm:"not null;" json:"pallet_detail_id" form:"pallet_detail_id" binding:"required"`
	CartonID       *string      `gorm:"not null;" json:"carton_id" form:"carton_id" binding:"required"`
	PalletOutNo    *string      `gorm:"not null;size:25" json:"pallet_out_no" form:"pallet_out_no" binding:"required"`
	IsSync         bool         `json:"is_sync" form:"is_sync"`
	IsActive       bool         `json:"is_active" form:"is_active"`
	CreatedAt      time.Time    `json:"created_at" default:"now"`
	UpdatedAt      time.Time    `json:"updated_at" default:"now"`
	PalletDetail   PalletDetail `gorm:"foreignKey:PalletDetailID;references:ID" json:"pallet_detail"`
	Carton         Carton       `gorm:"foreignKey:CartonID;references:ID" json:"carton"`
}

func (u *OrderPrepare) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}

type OrderShort struct {
	OrderShippingID string    `json:"order_shipping_id" form:"order_shipping_id"`
	OrderDetailID   string    `json:"order_detail_id" form:"order_detail_id"`
	OrderPlanID     string    `json:"order_plan_id" form:"order_plan_id"`
	OrderEtd        time.Time `json:"order_etd" form:"order_etd"`
	OrderCtn        int64     `json:"order_ctn" form:"order_ctn"`
	Ctn             int64     `json:"ctn" form:"ctn"`
}

type OrderSyncForm struct {
	IsSync bool `json:"is_sync" form:"is_sync"`
}
