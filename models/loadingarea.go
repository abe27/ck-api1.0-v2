package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type OrderLoadingArea struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	OrderZoneID *string   `gorm:"not null;" json:"order_zone_id,omitempty" form:"order_zone_id" binding:"required"`
	Prefix      string    `gorm:"not null;" json:"prefix,omitempty" form:"prefix" binding:"required"`
	LoadingArea string    `gorm:"not null;" json:"loading_area,omitempty" form:"loading_area" binding:"required"`
	Privilege   string    `gorm:"not null;" json:"privilege,omitempty" form:"privilege" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
	OrderZone   OrderZone `gorm:"foreignKey:OrderZoneID;references:ID" json:"order_zone,omitempty"`
}

func (obj *OrderLoadingArea) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderLoadingAreaForm struct {
	ID          string `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Bioat       int64  `json:"bioat,omitempty" form:"bioat" binding:"required"`
	Factory     string `json:"factory,omitempty" form:"factory"`
	Prefix      string `json:"prefix,omitempty" form:"prefix" binding:"required"`
	LoadingArea string `json:"loading_area,omitempty" form:"loading_area" binding:"required"`
	Privilege   string `json:"privilege,omitempty" form:"privilege" binding:"required"`
	IsActive    bool   `json:"is_active,omitempty" form:"is_active" binding:"required"`
}
