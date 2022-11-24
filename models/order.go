package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type OrderTitle struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *OrderTitle) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Order struct {
	ID           string         `gorm:"primaryKey;unique;index;size:21" json:"id,omitempty"`
	RowID        string         `gorm:"null;size:18" json:"row_id,omitempty" form:"row_id"`
	ConsigneeID  *string        `gorm:"not null" json:"consignee_id,omitempty" form:"consignee_id" binding:"required"`
	ShipmentID   *string        `gorm:"not null" json:"shipment_id,omitempty" form:"shipment_id" binding:"required"`
	EtdDate      *time.Time     `gorm:"not null;type:date;" json:"etd_date,omitempty" form:"etd_date" binding:"required"`
	PcID         *string        `gorm:"not null" json:"pc_id,omitempty" form:"pc_id" binding:"required"`
	CommercialID *string        `gorm:"not null" json:"commercial_id,omitempty" form:"commercial_id" binding:"required"`
	SampleFlgID  *string        `gorm:"not null" json:"sample_flg_id,omitempty" form:"sample_flg_id" binding:"required"`
	OrderTitleID *string        `gorm:"not null" json:"title_id,omitempty" form:"title_id" binding:"required"`
	Bioabt       int64          `json:"bioabt,omitempty" form:"bioabt" binding:"required"`
	CarrierCode  string         `gorm:"size:255;" json:"carrier_code,omitempty" form:"carrier_code" binding:"required"`
	ShipForm     string         `gorm:"size:255;" json:"ship_form,omitempty" form:"ship_form" binding:"required"`
	ShipTo       string         `gorm:"size:255;" json:"ship_to,omitempty" form:"ship_to" binding:"required"`
	ShipVia      string         `gorm:"size:255;" json:"ship_via,omitempty" form:"ship_via" binding:"required"`
	ShipDer      string         `gorm:"size:255;" json:"ship_der,omitempty" form:"ship_der" binding:"required"`
	LoadingArea  string         `gorm:"size:255;" json:"loading_area,omitempty" form:"loading_area" binding:"required"`
	Privilege    string         `gorm:"size:50;" json:"privilege,omitempty" form:"privilege" binding:"required"`
	ZoneCode     string         `gorm:"size:10;unique;" json:"zone_code,omitempty" form:"zone_code" binding:"required"`
	RunningSeq   int64          `json:"running_seq,omitempty" form:"running_seq" binding:"required"`
	IsChecked    bool           `json:"is_checked" form:"is_checked"`
	IsInvoice    bool           `json:"is_invoice" form:"is_invoice"`
	IsSync       bool           `json:"is_sync,omitempty" form:"is_sync"`
	IsActive     bool           `json:"is_active,omitempty" form:"is_active"`
	CreatedAt    time.Time      `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Consignee    Consignee      `gorm:"foreignKey:ConsigneeID;references:ID" json:"consignee,omitempty"`
	Shipment     Shipment       `gorm:"foreignKey:ShipmentID;references:ID" json:"shipment,omitempty"`
	Pc           Pc             `gorm:"foreignKey:PcID;references:ID" json:"pc,omitempty"`
	Commercial   Commercial     `gorm:"foreignKey:CommercialID;references:ID" json:"commercial,omitempty"`
	SampleFlg    SampleFlg      `gorm:"foreignKey:SampleFlgID;references:ID" json:"sample_flg,omitempty"`
	OrderTitle   OrderTitle     `gorm:"foreignKey:OrderTitleID;references:ID" json:"order_title,omitempty"`
	OrderDetail  []*OrderDetail `json:"order_detail,omitempty"`
	Pallet       []*Pallet      `json:"pallet,omitempty"`
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}

type OrderDetailForm struct {
	ReviseID      string  `json:"revise_id,omitempty" form:"revise_id"`
	OrderID       *string `json:"order_id,omitempty" form:"order_id" binding:"required"`
	Pono          *string `json:"pono,omitempty" form:"pono" binding:"required"`
	LedgerID      *string `json:"ledger_id,omitempty" form:"ledger_id" binding:"required"`
	OrderPlanID   *string `json:"order_plan_id,omitempty" form:"order_plan_id" binding:"required"`
	OrderCtn      int64   `json:"order_ctn" form:"order_ctn" binding:"required"`
	TotalOnPallet int64   `json:"total_on_pallet,omitempty" form:"total_on_pallet" binding:"required"`
}

type AddOrderDetailForm struct {
	Pono        string `json:"pono,omitempty" form:"pono" binding:"required"`
	PartNo      string `json:"part_no,omitempty" form:"part_no" binding:"required"`
	OrderBalQty int64  `json:"order_balqty,omitempty" form:"order_balqty" binding:"required"`
	OrderCtn    int64  `json:"order_ctn" form:"order_ctn" binding:"required"`
}

type OrderDetail struct {
	ID            string    `gorm:"primaryKey;unique;index;size:21" json:"id,omitempty"`
	RowID         string    `gorm:"null;size:18" json:"row_id,omitempty" form:"row_id"`
	OrderID       *string   `gorm:"not null;" json:"order_id,omitempty" form:"order_id" binding:"required"`
	Pono          *string   `gorm:"not null;size:25" json:"pono,omitempty" form:"pono" binding:"required"`
	LedgerID      *string   `gorm:"not null;" json:"ledger_id,omitempty" form:"ledger_id" binding:"required"`
	OrderPlanID   *string   `gorm:"not null;unique;" json:"order_plan_id,omitempty" form:"order_plan_id" binding:"required"`
	OrderCtn      int64     `json:"order_ctn" form:"order_ctn" binding:"required"`
	TotalOnPallet int64     `json:"total_on_pallet,omitempty" form:"total_on_pallet" binding:"required"`
	IsMatched     bool      `json:"is_matched" form:"is_matched"`
	IsSync        bool      `json:"is_sync,omitempty" form:"is_sync"`
	IsActive      bool      `json:"is_active,omitempty" form:"is_active"`
	CreatedAt     time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Order         Order     `gorm:"foreignKey:OrderID;references:ID" json:"order,omitempty"`
	Ledger        Ledger    `gorm:"foreignKey:LedgerID;references:ID" json:"ledger,omitempty"`
	OrderPlan     OrderPlan `gorm:"foreignKey:OrderPlanID;references:ID" json:"orderplan,omitempty"`
}

func (u *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}

type OrderPrepare struct {
	ID             string       `gorm:"primaryKey;unique;index;size:21" json:"id,omitempty"`
	PalletDetailID *string      `gorm:"not null;" json:"pallet_detail_id,omitempty" form:"pallet_detail_id" binding:"required"`
	CartonID       *string      `gorm:"not null;" json:"carton_id,omitempty" form:"carton_id" binding:"required"`
	PalletOutNo    *string      `gorm:"not null;size:25" json:"pallet_out_no,omitempty" form:"pallet_out_no" binding:"required"`
	IsSync         bool         `json:"is_sync,omitempty" form:"is_sync"`
	IsActive       bool         `json:"is_active,omitempty" form:"is_active"`
	CreatedAt      time.Time    `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt      time.Time    `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	PalletDetail   PalletDetail `gorm:"foreignKey:PalletDetailID;references:ID" json:"pallet_detail,omitempty"`
	Carton         Carton       `gorm:"foreignKey:CartonID;references:ID" json:"carton,omitempty"`
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
