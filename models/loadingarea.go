package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type OrderLoadingArea struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	OrderZoneID *string   `json:"order_zone_id" form:"order_zone_id" binding:"required"`
	Prefix      string    `json:"prefix" form:"prefix" binding:"required"`
	LoadingArea string    `json:"loading_area" form:"loading_area" binding:"required"`
	Privilege   string    `json:"privilege" form:"privilege" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
	OrderZone   OrderZone `gorm:"foreignKey:OrderZoneID;references:ID" json:"order_zone"`
}

func (obj *OrderLoadingArea) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderLoadingAreaForm struct {
	ID          string `gorm:"primaryKey;size:21" json:"id"`
	Bioat       int64  `json:"bioat" form:"bioat" binding:"required"`
	Factory     string `json:"factory" form:"factory"`
	Prefix      string `json:"prefix" form:"prefix" binding:"required"`
	LoadingArea string `json:"loading_area" form:"loading_area" binding:"required"`
	Privilege   string `json:"privilege" form:"privilege" binding:"required"`
	IsActive    bool   `json:"is_active" form:"is_active" binding:"required"`
}
