package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type FileEdi struct {
	ID         string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	FactoryID  *string   `gorm:"not null;" json:"factory_id,omitempty" form:"factory_id"`
	MailboxID  *string   `gorm:"not null;" json:"mailbox_id,omitempty" form:"mailbox_id" binding:"required"`
	FileTypeID *string   `gorm:"not null;" json:"file_type_id,omitempty" form:"file_type_id"`
	BatchNo    string    `gorm:"not null;unique;size:10" json:"batch_no,omitempty" form:"batch_no" binding:"required"`
	Size       int64     `json:"size,omitempty" form:"size"`
	BatchName  string    `gorm:"size:50" json:"batch_name,omitempty" form:"batch_name"`
	CreationOn time.Time `json:"creation_on,omitempty" form:"creation_on"`
	Flags      string    `gorm:"size:5" json:"flags,omitempty" form:"flags" binding:"required"`
	FormatType string    `gorm:"size:5" json:"format_type,omitempty" form:"format_type" binding:"required"`
	Originator string    `gorm:"size:10" json:"originator,omitempty" form:"originator" binding:"required"`
	BatchPath  string    `gorm:"size:255" json:"batch_path,omitempty"`
	IsDownload bool      `json:"is_download,omitempty" form:"is_download" binding:"required"`
	IsActive   bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt  time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" default:"now"`
	Factory    Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
	Mailbox    Mailbox   `gorm:"foreignKey:MailboxID;references:ID" json:"mailbox,omitempty"`
	FileType   FileType  `gorm:"foreignKey:FileTypeID;references:ID" json:"file_type,omitempty"`
}

func (obj *FileEdi) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type GEDIReceive struct {
	Factory          string // "factory": factory,
	FacZone          string // "faczone": str(line[4 : (4 + 3)]).lstrip().rstrip(),
	ReceivingKey     string // "receivingkey": str(line[4 : (4 + 12)]).lstrip().rstrip(),
	PartNo           string // "partno": str(line[76 : (76 + 25)]).lstrip().rstrip(),
	PartName         string // "partname": str(line[101 : (101 + 25)]).lstrip().rstrip(),
	Vendor           string // "vendor": factory,
	Cd               string // "cd": cd,
	Unit             string // "unit": unit,
	Whs              string // "whs": factory,
	Tagrp            string // "tagrp": "C",
	RecType          string // "recisstype": recisstype,
	PlanType         string // "plantype": plantype,
	RecID            string // "recid": str(line[0:4]).lstrip().rstrip(),
	Aetono           string // "aetono": str(line[4 : (4 + 12)]).lstrip().rstrip(),
	Aetodt           string // "aetodt": str(line[16 : (16 + 10)]).lstrip().rstrip(),
	Aetctn           int64  // "aetctn": float(str(line[26 : (26 + 9)]).lstrip().rstrip()),
	Aetfob           int64  // "aetfob": float(str(line[35 : (35 + 9)]).lstrip().rstrip()),
	Aenewt           int64  // "aenewt": float(str(line[44 : (44 + 11)]).lstrip().rstrip()),
	Aentun           string // "aentun": str(line[55 : (55 + 5)]).lstrip().rstrip(),
	Aegrwt           int64  // "aegrwt": float(str(line[60 : (60 + 11)]).lstrip().rstrip()),
	Aegwun           string // "aegwun": str(line[71 : (71 + 5)]).lstrip().rstrip(),
	Aeypat           string // "aeypat": str(line[76 : (76 + 25)]).lstrip().rstrip(),
	Aeedes           string // "aeedes": str(self.__check_partname(factory, self.__re_partname(line[101 : (101 + 25)]))),
	Aetdes           string // "aetdes": str(self.__check_partname(factory, self.__re_partname(line[101 : (101 + 25)]))),
	Aetarf           int64  // "aetarf": float(str(line[151 : (151 + 10)]).lstrip().rstrip()),
	Aestat           int64  // "aestat": float(str(line[161 : (161 + 10)]).lstrip().rstrip()),
	Aebrnd           int64  // "aebrnd": float(str(line[171 : (171 + 10)]).lstrip().rstrip()),
	Aertnt           int64  // "aertnt": float(str(line[181 : (181 + 5)]).lstrip().rstrip()),
	Aetrty           int64  // "aetrty": float(str(line[186 : (186 + 5)]).lstrip().rstrip()),
	Aesppm           int64  // "aesppm": float(str(line[191 : (191 + 5)]).lstrip().rstrip()),
	AeQty1           int64  // "aeQty1": float(str(line[196 : (196 + 9)]).lstrip().rstrip()),
	AeQty2           int64  // "aeQty2": float(str(line[205 : (205 + 9)]).lstrip().rstrip()),
	Aeuntp           int64  // "aeuntp": float(str(line[214 : (214 + 9)]).lstrip().rstrip()),
	Aeamot           int64  // "aeamot": float(str(line[223 : (223 + 11)]).lstrip().rstrip()),
	Plnctn           int64  // "plnctn": float(str(line[26 : (26 + 9)]).lstrip().rstrip()),
	PlnQty           int64  // "plnQty": float(str(line[196 : (196 + 9)]).lstrip().rstrip()),
	Minimum          int64  // "minimum": 0,
	Maximum          int64  // "maximum": 0,
	Picshelfbin      string // "picshelfbin": "PNON",
	Stkshelfbin      string // "stkshelfbin": "SNON",
	Ovsshelfbin      string // "ovsshelfbin": "ONON",
	PicshelfbasicQty int64  // "picshelfbasicQty": 0,
	Outerpcs         int64  // "outerpcs": 0,
	AllocateQty      int64  // "allocateQty": 0,
}
