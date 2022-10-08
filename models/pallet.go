package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Pallet struct {
	ID           string     `gorm:"primaryKey;size:21" json:"id,omitempty"`
	OrderID      *string    `gorm:"not null;unique;" json:"order_id,omitempty" form:"order_id" binding:"required"`
	PalletTypeID *string    `gorm:"not null;" json:"pallet_type_id,omitempty" form:"pallet_type_id" binding:"required"`
	PalletNo     int64      `gorm:"not null;" json:"pallet_no,omitempty" form:"pallet_no" binding:"required"`
	PalletTotal  int64      `gorm:"not null;" json:"pallet_total,omitempty" form:"pallet_total" binding:"required"`
	IsSync       bool       `json:"is_sync,omitempty" form:"is_sync" binding:"required"`
	IsActive     bool       `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt    time.Time  `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Order        Order      `gorm:"foreignKey:OrderID;references:ID" json:"order,omitempty"`
	PalletType   PalletType `gorm:"foreignKey:PalletTypeID;references:ID"  json:"pallet_type"`
}

func (obj *Pallet) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PalletDetail struct {
	ID            string      `gorm:"primaryKey;size:21" json:"id,omitempty"`
	PalletID      *string     `json:"pallet_id,omitempty" form:"pallet_id" binding:"required"`
	OrderDetailID *string     `gorm:"not null;" json:"order_detail_id,omitempty" form:"order_detail_id" binding:"required"`
	SeqNo         int64       `gorm:"not null;" json:"seq_no,omitempty" form:"seq_no" binding:"required"`
	IsPrintLabel  bool        `json:"is_print_label,omitempty" form:"is_print_label" binding:"required"`
	IsActive      bool        `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt     time.Time   `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Pallet        Pallet      `gorm:"foreignKey:PalletID;references:ID" json:"pallet,omitempty"`
	OrderDetail   OrderDetail `gorm:"foreignKey:OrderDetailID;references:ID"  json:"order_detail,omitempty"`
}

func (obj *PalletDetail) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
