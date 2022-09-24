package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Area struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Area) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Whs struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Value       string    `gorm:"size:5;" json:"value,omitempty" form:"value" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Factory struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	InvPrefix   string    `gorm:"size:5;" json:"inv_prefix,omitempty" form:"inv_prefix" binding:"required"`
	LabelPrefix string    `gorm:"size:5;" json:"label_prefix,omitempty" form:"label_prefix" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Factory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PrefixName struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *PrefixName) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Position struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Position) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Department struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Department) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Unit struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PartType struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *PartType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type FileType struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:50" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *FileType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Mailbox struct {
	ID        string    `gorm:"primaryKey;size:21" json:"id"`
	Mailbox   string    `gorm:"not null;unique;size:21" json:"mailbox" form:"mailbox" binding:"required"`
	Password  string    `gorm:"size:50" json:"password" form:"password" binding:"required"`
	HostUrl   string    `json:"host_url" form:"host_url" binding:"required"`
	AreaID    *string   `json:"area_id" form:"area_id" binding:"required"`
	IsActive  bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	Area      Area      `gorm:"foreignKey:AreaID;references:ID;constraint:OnDelete:CASCADE;" json:"area"`
}

func (obj *Mailbox) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}