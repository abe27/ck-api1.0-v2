package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type CartonHistory struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	RowID       string    `gorm:"not null;size:18" json:"row_id,omitempty" form:"row_id" binding:"required"`
	Whs         string    `gorm:"not null;" json:"whs,omitempty" form:"whs" binding:"required"`
	PartNo      string    `gorm:"not null;" json:"part_no,omitempty" form:"part_no" binding:"required"`
	LotNo       string    `gorm:"not null;size:8;" json:"lot_no,omitempty" form:"lot_no" binding:"required"`
	SerialNo    string    `gorm:"not null;size:10;" json:"serial_no,omitempty" form:"serial_no" binding:"required"`
	DieNo       string    `gorm:"null;size:10;" json:"die_no,omitempty" form:"die_no" binding:"required"`
	RevisionNo  string    `gorm:"null;size:25;" json:"rev_no,omitempty" form:"rev_no" binding:"required"`
	Qty         int64     `json:"qty,omitempty" form:"qty" binding:"required"`
	Shelve      string    `gorm:"null;size:20;" json:"shelve,omitempty" form:"shelve" binding:"required"`
	IpAddress   string    `gorm:"null;size:25" json:"ip_address,omitempty" form:"ip_address"`
	EmpID       string    `gorm:"null;size:25" json:"emp_id,omitempty" form:"emp_id"`
	RefNo       string    `gorm:"null;size:25" json:"ref_no,omitempty" form:"ref_no"`
	ReceiveNo   string    `gorm:"null;size:25" json:"receive_no,omitempty" form:"receive_no"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (u *CartonHistory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}

type SyncLogger struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:100" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsSuccess   bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *SyncLogger) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
