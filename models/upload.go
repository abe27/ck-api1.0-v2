package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type CartonNotReceive struct {
	ID            string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	TransferOutNo string    `gorm:"not null;size:25" json:"transfer_out_no,omitempty" form:"transfer_out_no" binding:"required"`
	PartNo        string    `gorm:"not null;" json:"part_no,omitempty" form:"part_no" binding:"required"`
	LotNo         string    `gorm:"not null;size:8;" json:"lot_no,omitempty" form:"lot_no" binding:"required"`
	SerialNo      string    `gorm:"not null;size:10;" json:"serial_no,omitempty" form:"serial_no" binding:"required"`
	Qty           int64     `json:"qty,omitempty" form:"qty" binding:"required"`
	IsSync        bool      `json:"is_sync,omitempty" form:"is_sync" default:"false"`
	CreatedAt     time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (u *CartonNotReceive) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
