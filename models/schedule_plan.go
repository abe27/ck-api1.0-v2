package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type SchedulePlan struct {
	ID        string    `gorm:"primaryKey;unique;index;size:21" json:"id,omitempty"`
	IsSync    bool      `json:"is_sync,omitempty" form:"is_sync"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (u *SchedulePlan) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return
}
