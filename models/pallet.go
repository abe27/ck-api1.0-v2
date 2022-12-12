package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Pallet struct {
	ID           string         `gorm:"primaryKey;size:21" json:"id"`
	OrderID      *string        `gorm:"not null;" json:"order_id" form:"order_id" binding:"required"`
	PalletTypeID *string        `gorm:"not null;" json:"pallet_type_id" form:"pallet_type_id" binding:"required"`
	PalletPrefix string         `gorm:"not null;size:1;" json:"pallet_prefix" form:"pallet_prefix" default:"P"`
	PalletNo     int64          `gorm:"not null;" json:"pallet_no" form:"pallet_no" binding:"required"`
	PalletTotal  int64          `gorm:"not null;" json:"pallet_total" form:"pallet_total" binding:"required"`
	IsSync       bool           `json:"is_sync" form:"is_sync" binding:"required"`
	IsActive     bool           `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt    time.Time      `json:"created_at" default:"now"`
	UpdatedAt    time.Time      `json:"updated_at" default:"now"`
	Order        Order          `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	PalletType   PalletType     `gorm:"foreignKey:PalletTypeID;references:ID"  json:"pallet_type"`
	PalletDetail []PalletDetail `json:"pallet_detail"`
}

func (obj *Pallet) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PalletDetail struct {
	ID            string      `gorm:"primaryKey;size:21" json:"id"`
	PalletID      *string     `json:"pallet_id" form:"pallet_id" binding:"required"`
	OrderDetailID *string     `gorm:"not null;" json:"order_detail_id" form:"order_detail_id" binding:"required"`
	SeqNo         int64       `gorm:"not null;" json:"seq_no" form:"seq_no" binding:"required"`
	IsPrintLabel  bool        `json:"is_print_label" form:"is_print_label" binding:"required"`
	IsActive      bool        `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt     time.Time   `json:"created_at" default:"now"`
	UpdatedAt     time.Time   `json:"updated_at" default:"now"`
	Pallet        Pallet      `gorm:"foreignKey:PalletID;references:ID" json:"pallet"`
	OrderDetail   OrderDetail `gorm:"foreignKey:OrderDetailID;references:ID"  json:"order_detail"`
}

func (obj *PalletDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
