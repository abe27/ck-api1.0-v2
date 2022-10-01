package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Receive struct {
	ID            string          `gorm:"primaryKey;size:21" json:"id,omitempty"`
	FileEdiID     *string         `gorm:"not null" form:"file_edi_id" json:"file_edi_id,omitempty"`
	ReceiveTypeID *string         `gorm:"not null" form:"receive_type_id" json:"receive_type_id,omitempty"`
	ReceiveDate   time.Time       `json:"receive_date,omitempty" form:"receive_date" binding:"required"`
	TransferOutNo string          `gorm:"not null;unique;size:15" json:"transfer_out_no,omitempty" form:"transfer_out_no" binding:"required"`
	TexNo         string          `gorm:"size:15;" json:"tex_no,omitempty" form:"tex_no"`
	Item          int64           `json:"item,omitempty" form:"item" default:"0"`
	PlanCtn       int64           `json:"plan_ctn,omitempty" form:"plan_ctn" default:"0"`
	ReceiveCtn    int64           `json:"receive_ctn,omitempty" form:"receive_ctn" default:"0"`
	IsSync        bool            `json:"is_sync,omitempty" form:"is_sync" default:"true"`
	IsActive      bool            `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt     time.Time       `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	FileEdi       FileEdi         `gorm:"foreignKey:FileEdiID;references:ID" json:"file_edi,omitempty"`
	ReceiveType   ReceiveType     `gorm:"foreignKey:ReceiveTypeID;references:ID" json:"receive_type,omitempty"`
	ReceiveDetail []ReceiveDetail `json:"receive_detail,omitempty"`
}

func (u *Receive) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type ReceiveDetail struct {
	ID        string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	ReceiveID *string   `gorm:"not null;" form:"receive_id" json:"receive_id,omitempty"`
	LedgerID  *string   `gorm:"not null;" form:"ledger_id" json:"ledger_id,omitempty"`
	PlanQty   int64     `json:"plan_qty,omitempty" form:"plan_qty"`
	PlanCtn   int64     `json:"plan_ctn,omitempty" form:"plan_ctn"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Receive   Receive   `gorm:"foreignKey:ReceiveID;references:ID" json:"receive,omitempty"`
	Ledger    Ledger    `gorm:"foreignKey:LedgerID;references:ID" json:"ledger,omitempty"`
	// Cartons   []Carton  `json:"carton"`
}

func (u *ReceiveDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type FormReceiveDetail struct {
	WhsID     string `form:"whs_id" binding:"required"`
	ReceiveID string `form:"receive_id" binding:"required"`
	LedgerID  string `form:"ledger_id" binding:"required"`
	PlanQty   int    `form:"plan_qty" binding:"required"`
	PlanCtn   int    `form:"plan_ctn" binding:"required"`
	IsActive  bool   `form:"is_active" binding:"required"`
}

type ReceiveIcam struct {
	No        string `json:"no" binding:"required"`
	Date      string `json:"date" binding:"required"`
	FactoryId string `json:"factory_id" binding:"required"`
	WhsId     string `json:"whs_id" binding:"required"`
	PartId    string `json:"part_id" binding:"required"`
	Ctn       int    `json:"ctn" binding:"required"`
}

type ReceiveEntForm struct {
	IsSync   bool `json:"is_sync,omitempty" form:"is_sync" default:"true"`
	IsActive bool `form:"is_active" binding:"required"`
}
