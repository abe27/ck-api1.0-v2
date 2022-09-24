package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type CartonHistory struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Whs         string    `json:"whs" form:"whs" binding:"required"`
	PartNo      string    `json:"part_no" form:"part_no" binding:"required"`
	LotNo       string    `gorm:"size:8;" json:"lot_no" form:"lot_no" binding:"required"`
	SerialNo    string    `gorm:"size:10;" json:"serial_no" form:"serial_no" binding:"required"`
	DieNo       string    `gorm:"size:10;" json:"die_no" form:"die_no" binding:"required"`
	RevisionNo  string    `gorm:"size:25;" json:"rev_no" form:"rev_no" binding:"required"`
	Qty         int64     `json:"qty" form:"qty" binding:"required"`
	Shelve      string    `gorm:"size:20;" json:"shelve" form:"shelve" binding:"required"`
	IpAddress   string    `gorm:"size:25" json:"ip_address" form:"ip_address"`
	EmpID       string    `gorm:"size:25" json:"emp_id" form:"emp_id"`
	RefNo       string    `gorm:"size:25" json:"ref_no" form:"ref_no"`
	Description string    `json:"description" form:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (u *CartonHistory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type SyncLogger struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;size:100" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsSuccess   bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *SyncLogger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
