package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Affcode struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Affcode) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Customer struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type CustomerAddress struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:50" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *CustomerAddress) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Consignee struct {
	ID                string          `gorm:"primaryKey;size:21" json:"id"`
	WhsID             *string         `gorm:"not null;" json:"whs_id" form:"whs_id"`
	FactoryID         *string         `gorm:"not null;" json:"factory_id" form:"factory_id"`
	AffcodeID         *string         `gorm:"not null;" json:"affcode_id" form:"affcode_id" binding:"required"`
	CustomerID        *string         `gorm:"not null;" json:"customer_id" form:"customer_id" binding:"required"`
	CustomerAddressID *string         `gorm:"null;" json:"customer_ddress_id" form:"customer_address_id"`
	Prefix            string          `gorm:"not null" json:"prefix" form:"prefix" binding:"required"`
	IsActive          bool            `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt         time.Time       `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt         time.Time       `json:"updated_at" form:"updated_at" default:"now"`
	Whs               Whs             `gorm:"foreignKey:WhsID;references:ID" json:"whs"`
	Factory           Factory         `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Affcode           Affcode         `gorm:"foreignKey:AffcodeID;references:ID" json:"affcode"`
	Customer          Customer        `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	CustomerAddress   CustomerAddress `gorm:"foreignKey:CustomerAddressID;references:ID" json:"customer_address"`
}

func (obj *Consignee) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderZone struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Value       int64     `gorm:"not null;" json:"value" form:"value" binding:"required"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id" form:"factory_id" binding:"required"`
	WhsID       *string   `gorm:"not null;" json:"whs_id" form:"whs_id" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID" json:"whs"`
}

func (obj *OrderZone) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type LastInvoice struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id" form:"factory_id" binding:"required"`
	AffcodeID   *string   `gorm:"not null;" json:"affcode_id" form:"affcode_id" binding:"required"`
	OnYear      int64     `gorm:"not null;" json:"on_year" form:"on_year" binding:"required"`
	LastRunning int64     `json:"last_running" form:"last_running" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Affcode     Affcode   `gorm:"foreignKey:AffcodeID;references:ID" json:"affcode"`
}

func (obj *LastInvoice) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
