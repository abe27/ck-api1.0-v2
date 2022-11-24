package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type AssetType struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *AssetType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
