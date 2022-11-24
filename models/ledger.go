package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Part struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Slug        string    `gorm:"size:50;unique;not null;" json:"slug,omitempty" form:"slug" binding:"required"`
	Title       string    `gorm:"size:50;unique;not null;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (u *Part) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type Ledger struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	WhsID       *string   `gorm:"not null;" json:"whs_id,omitempty" form:"whs_id" binding:"required"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id,omitempty" form:"factory_id" binding:"required"`
	PartID      *string   `gorm:"not null;" json:"part_id,omitempty" form:"part_id" binding:"required"`
	PartTypeID  *string   `gorm:"not null;" json:"part_type_id,omitempty" form:"part_type_id" binding:"required"`
	UnitID      *string   `gorm:"not null;" json:"unit_id,omitempty" form:"unit_id" binding:"required"`
	DimWidth    float64   `json:"dim_width,omitempty" form:"dim_width" default:"0"`
	DimLength   float64   `json:"dim_length,omitempty" form:"dim_length" default:"0"`
	DimHeight   float64   `json:"dim_height,omitempty" form:"dim_height" default:"0"`
	GrossWeight float64   `json:"gross_weight,omitempty" form:"gross_weight" default:"0"`
	NetWeight   float64   `json:"net_weight,omitempty" form:"net_weight" default:"0"`
	Qty         float64   `json:"qty,omitempty" form:"qty" default:"0"`
	Ctn         float64   `json:"ctn,omitempty" form:"ctn" default:"0"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID" json:"whs,omitempty"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
	Part        Part      `gorm:"foreignKey:PartID;references:ID;" json:"part,omitempty"`
	PartType    PartType  `gorm:"foreignKey:PartTypeID;references:ID" json:"part_type,omitempty"`
	Unit        Unit      `gorm:"foreignKey:UnitID;references:ID" json:"unit,omitempty"`
}

func (u *Ledger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type Carton struct {
	ID              string        `gorm:"primaryKey;size:21" json:"id,omitempty"`
	RowID           string        `gorm:"not null;size:18" json:"row_id,omitempty" form:"row_id"`
	LedgerID        *string       `gorm:"not null;" json:"ledger_id,omitempty" form:"ledger_id" binding:"required"`
	LocationID      *string       `json:"location_id,omitempty" form:"location_id" binding:"required"`
	ReceiveDetailID *string       `son:"receive_detail_id,omitempty" form:"receive_detail_id" binding:"required"`
	LotNo           string        `gorm:"not null;size:8;" json:"lot_no,omitempty" form:"lot_no" binding:"required"`
	SerialNo        string        `gorm:"not null;size:10;unique;" json:"serial_no,omitempty" form:"serial_no" binding:"required"`
	LineNo          string        `gorm:"size:10;" json:"line_no,omitempty" form:"line_no"`
	RevisionNo      string        `gorm:"size:10;" json:"revise_no,omitempty" form:"revise_no"`
	Qty             float64       `json:"qty,omitempty" form:"qty" default:"0"`
	PalletNo        string        `json:"pallet_no,omitempty" form:"pallet_no" binding:"required"`
	IsActive        bool          `json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt       time.Time     `json:"created_at,omitempty" default:"now"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" default:"now"`
	Ledger          Ledger        `gorm:"foreignKey:LedgerID;references:ID" json:"ledger,omitempty"`
	Location        Location      `gorm:"foreignKey:LocationID;references:ID" json:"location,omitempty"`
	ReceiveDetail   ReceiveDetail `gorm:"foreignKey:ReceiveDetailID;references:ID;" json:"receive_detail,omitempty"`
}

func (u *Carton) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
