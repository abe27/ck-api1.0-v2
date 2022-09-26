package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Part struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Slug        string    `gorm:"size:50;unique;not null;" json:"slug" form:"slug" binding:"required"`
	Title       string    `gorm:"size:50;unique;not null;" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (u *Part) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type Ledger struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	FactoryID   *string   `json:"factory_id" form:"factory_id" binding:"required"`
	PartID      *string   `json:"part_id" form:"part_id" binding:"required"`
	PartTypeID  *string   `json:"part_type_id" form:"part_type_id" binding:"required"`
	UnitID      *string   `json:"unit_id" form:"unit_id" binding:"required"`
	DimWidth    float64   `json:"dim_width" form:"dim_width" default:"0"`
	DimLength   float64   `json:"dim_length" form:"dim_length" default:"0"`
	DimHeight   float64   `json:"dim_height" form:"dim_height" default:"0"`
	GrossWeight float64   `json:"gross_weight" form:"gross_weight" default:"0"`
	NetWeight   float64   `json:"net_weight" form:"net_weight" default:"0"`
	Qty         float64   `json:"qty" form:"qty" default:"0"`
	Ctn         float64   `json:"ctn" form:"ctn" default:"0"`
	IsActive    bool      `json:"is_active" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Part        Part      `gorm:"foreignKey:PartID;references:ID;" json:"part"`
	PartType    PartType  `gorm:"foreignKey:PartTypeID;references:ID" json:"part_type"`
	Unit        Unit      `gorm:"foreignKey:UnitID;references:ID" json:"unit"`
}

func (u *Ledger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
