package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Affcode struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Affcode) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Customer struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type CustomerAddress struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:50" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *CustomerAddress) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Consignee struct {
	ID                string          `gorm:"primaryKey;size:21" json:"id,omitempty"`
	WhsID             *string         `gorm:"not null;" json:"whs_id,omitempty" form:"whs_id"`
	FactoryID         *string         `gorm:"not null;" json:"factory_id,omitempty" form:"factory_id"`
	AffcodeID         *string         `gorm:"not null;" json:"affcode_id,omitempty" form:"affcode_id" binding:"required"`
	CustomerID        *string         `gorm:"not null;" json:"customer_id,omitempty" form:"customer_id" binding:"required"`
	CustomerAddressID *string         `gorm:"null;" json:"customer_ddress_id,omitempty" form:"customer_address_id"`
	Prefix            string          `gorm:"not null" json:"prefix,omitempty" form:"prefix" binding:"required"`
	IsActive          bool            `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt         time.Time       `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt         time.Time       `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Whs               Whs             `gorm:"foreignKey:WhsID;references:ID" json:"whs,omitempty"`
	Factory           Factory         `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
	Affcode           Affcode         `gorm:"foreignKey:AffcodeID;references:ID" json:"affcode,omitempty"`
	Customer          Customer        `gorm:"foreignKey:CustomerID;references:ID" json:"customer,omitempty"`
	CustomerAddress   CustomerAddress `gorm:"foreignKey:CustomerAddressID;references:ID" json:"customer_address,omitempty"`
	OrderGroup        []*OrderGroup   `json:"order_group,omitempty"`
}

func (obj *Consignee) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderZone struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Value       int64     `gorm:"not null;" json:"value,omitempty" form:"value" binding:"required"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id,omitempty" form:"factory_id" binding:"required"`
	WhsID       *string   `gorm:"not null;" json:"whs_id,omitempty" form:"whs_id" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID" json:"whs,omitempty"`
}

func (obj *OrderZone) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type LastInvoice struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	FactoryID   *string   `gorm:"not null;" json:"factory_id,omitempty" form:"factory_id" binding:"required"`
	AffcodeID   *string   `gorm:"not null;" json:"affcode_id,omitempty" form:"affcode_id" binding:"required"`
	OnYear      int64     `gorm:"not null;" json:"on_year,omitempty" form:"on_year" binding:"required"`
	LastRunning int64     `json:"last_running,omitempty" form:"last_running" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
	Affcode     Affcode   `gorm:"foreignKey:AffcodeID;references:ID" json:"affcode,omitempty"`
}

func (obj *LastInvoice) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
