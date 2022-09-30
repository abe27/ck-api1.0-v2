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
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
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
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
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
