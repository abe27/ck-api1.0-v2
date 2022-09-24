package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FileEdi struct {
	ID         string    `gorm:"primaryKey;size:21" json:"id"`
	FactoryID  *string   `json:"factory_id" form:"factory_id"`
	MailboxID  *string   `json:"mailbox_id" form:"mailbox_id" binding:"required"`
	FileTypeID *string   `json:"file_type_id" form:"file_type_id"`
	BatchNo    string    `gorm:"unique;primaryKey;size:10" json:"batch_no" form:"batch_no" binding:"required"`
	Size       int64     `json:"size" form:"size"`
	BatchName  string    `gorm:"size:50" json:"batch_name" form:"batch_name"`
	CreationOn time.Time `json:"creation_on" form:"creation_on"`
	Flags      string    `gorm:"size:5" json:"flags" form:"flags" binding:"required"`
	FormatType string    `gorm:"size:5" json:"format_type" form:"format_type" binding:"required"`
	Originator string    `gorm:"size:10" json:"originator" form:"originator" binding:"required"`
	BatchPath  string    `gorm:"size:255" json:"batch_path"`
	IsDownload bool      `json:"is_download" form:"is_download" binding:"required"`
	IsActive   bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt  time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt  time.Time `json:"updated_at" form:"updated_at" default:"now"`
	Factory    Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory"`
	Mailbox    Mailbox   `gorm:"foreignKey:MailboxID;references:ID" json:"mailbox"`
	FileType   FileType  `gorm:"foreignKey:FileTypeID;references:ID" json:"file_type"`
}

func (obj *FileEdi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
