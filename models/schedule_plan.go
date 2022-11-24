package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type SchedulePlan struct {
	ID            string      `gorm:"primaryKey;unique;index;size:21" json:"id,omitempty"`
	CustomerID    string      `gorm:"not null;" json:"customer_id,omitempty" form:"customer_id"`
	ShipmentID    string      `gorm:"not null;" json:"shipment_id" form:"shipment_id"`
	WhsID         string      `gorm:"not null;" json:"whs_id" form:"whs_id"`
	PlanningDayID string      `gorm:"not null;" json:"planning_day_id" form:"planning_day_id"`
	IsSync        bool        `json:"is_sync,omitempty" form:"is_sync"`
	IsActive      bool        `json:"is_active,omitempty" form:"is_active"`
	CreatedAt     time.Time   `json:"created_at,omitempty" default:"now"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty" default:"now"`
	Customer      Customer    `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	Shipment      Shipment    `gorm:"foreignKey:ShipmentID;references:ID" json:"shipment"`
	Whs           Whs         `gorm:"foreignKey:WhsID;references:ID" json:"whs"`
	PlanningDay   PlanningDay `gorm:"foreignKey:PlanningDayID;references:ID" json:"planning_day"`
}

func (u *SchedulePlan) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}
