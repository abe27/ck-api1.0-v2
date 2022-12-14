package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Receive struct {
	ID            string          `gorm:"primaryKey;size:21" json:"id"`
	RowID         string          `gorm:"null;size:18" json:"row_id" form:"row_id"`
	FileEdiID     *string         `gorm:"not null" form:"file_edi_id" json:"file_edi_id"`
	ReceiveTypeID *string         `gorm:"not null" form:"receive_type_id" json:"receive_type_id"`
	ReceiveDate   time.Time       `gorm:"type:date" json:"receive_date" form:"receive_date" binding:"required"`
	TransferOutNo string          `gorm:"not null;unique;size:15" json:"transfer_out_no" form:"transfer_out_no" binding:"required"`
	TexNo         string          `gorm:"size:15;" json:"tex_no" form:"tex_no"`
	Item          int64           `json:"item" form:"item" default:"0"`
	PlanCtn       int64           `json:"plan_ctn" form:"plan_ctn" default:"0"`
	ReceiveCtn    int64           `json:"receive_ctn" form:"receive_ctn" default:"0"`
	IsSync        bool            `json:"is_sync" form:"is_sync" default:"true"`
	IsActive      bool            `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt     time.Time       `json:"created_at" default:"now"`
	UpdatedAt     time.Time       `json:"updated_at" default:"now"`
	FileEdi       FileEdi         `gorm:"foreignKey:FileEdiID;references:ID" json:"file_edi"`
	ReceiveType   ReceiveType     `gorm:"foreignKey:ReceiveTypeID;references:ID" json:"receive_type"`
	ReceiveDetail []ReceiveDetail `json:"receive_detail"`
}

func (u *Receive) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type ReceiveDetail struct {
	ID               string             `gorm:"primaryKey;size:21" json:"id"`
	RowID            string             `gorm:"null;size:18" json:"row_id" form:"row_id"`
	ReceiveID        *string            `gorm:"not null;" form:"receive_id" json:"receive_id"`
	LedgerID         *string            `gorm:"not null;" form:"ledger_id" json:"ledger_id"`
	PlanQty          int64              `json:"plan_qty" form:"plan_qty"`
	PlanCtn          int64              `json:"plan_ctn" form:"plan_ctn"`
	IsActive         bool               `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt        time.Time          `json:"created_at" default:"now"`
	UpdatedAt        time.Time          `json:"updated_at" default:"now"`
	Receive          Receive            `gorm:"foreignKey:ReceiveID;references:ID" json:"receive"`
	Ledger           Ledger             `gorm:"foreignKey:LedgerID;references:ID" json:"ledger"`
	CartonNotReceive []CartonNotReceive `json:"receive_carton"`
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
	IsSync   bool `json:"is_sync" form:"is_sync" default:"true"`
	IsActive bool `form:"is_active" binding:"required"`
}
