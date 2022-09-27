package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Area struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	InvPrefix   string    `gorm:"size:5;" json:"inv_prefix,omitempty" form:"inv_prefix" binding:"required"`
	LabelPrefix string    `gorm:"size:5;" json:"label_prefix,omitempty" form:"label_prefix" binding:"required"`
	PartUnit    string    `gorm:"size:50;" json:"part_unit,omitempty" form:"part_unit" binding:"required"`
	CdCode      string    `gorm:"size:5;" json:"cd_code,omitempty" form:"cd_code" binding:"required"`
	PartType    string    `gorm:"size:50;" json:"part_type,omitempty" form:"part_type" binding:"required"`
	Sortg1      string    `gorm:"size:50;" json:"sortg1,omitempty" form:"sortg1" binding:"required"`
	Sortg2      string    `gorm:"size:50;" json:"sortg2,omitempty" form:"sortg2" binding:"required"`
	Sortg3      string    `gorm:"size:50;" json:"sortg3,omitempty" form:"sortg3" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
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

type ReceiveType struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id"`
	WhsID       string    `gorm:"not null;" json:"whs_id" form:"whs_id"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"whs"`
}

func (obj *ReceiveType) BeforeCreate(tx *gorm.DB) (err error) {
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

type Pc struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:50" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Pc) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Commercial struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:50" json:"title" form:"title" binding:"required"`
	Prefix      string    `gorm:"size:5" json:"prefix" form:"prefix" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Commercial) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type SampleFlg struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:50" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *SampleFlg) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ReviseOrder struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *ReviseOrder) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Shipment struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *Shipment) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderType struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *OrderType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

// N=All,F=3 Front,E=3 End,O=Sprit Order
type OrderGroupType struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id"`
	Title       string    `gorm:"not null;unique;size:15" json:"title" form:"title" binding:"required"`
	Description string    `json:"description" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *OrderGroupType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderGroup struct {
	ID               string         `gorm:"primaryKey;size:21" json:"id"`
	UserID           *string        `gorm:"not null;" json:"user_id" form:"user_id" binding:"required"`
	ConsigneeID      *string        `gorm:"not null;" json:"consignee_id" form:"consignee_id" binding:"required"`
	OrderGroupTypeID *string        `gorm:"not null;" json:"order_group_type_id" form:"order_group_type_id" binding:"required"`
	SubOrder         string         `gorm:"not null;size:15" json:"sub_order" form:"sub_order" binding:"required"`
	Description      string         `json:"description" form:"description" binding:"required"`
	IsActive         bool           `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt        time.Time      `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt        time.Time      `json:"updated_at" form:"updated_at" default:"now"`
	User             User           `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Consignee        Consignee      `gorm:"foreignKey:ConsigneeID;references:ID" json:"consignee"`
	OrderGroupType   OrderGroupType `gorm:"foreignKey:OrderGroupTypeID;references:ID" json:"order_group_type"`
}

func (obj *OrderGroup) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type FormGroupConsignee struct {
	UserID           string `json:"user_id" form:"user_id" binding:"required"`
	WhsID            string `json:"whs_id" form:"whs_id"`
	FactoryID        string `json:"factory_id" form:"factory_id"`
	AffcodeID        string `json:"affcode_id" form:"affcode_id" binding:"required"`
	CustcodeID       string `json:"custcode_id" form:"custcode_id" binding:"required"`
	OrderGroupTypeID string `json:"order_group_type_id" form:"order_group_type_id" binding:"required"`
	SubOrder         string `gorm:"size:15" json:"sub_order" form:"sub_order" binding:"required"`
	Description      string `json:"description" form:"description" binding:"required"`
	IsActive         bool   `json:"is_active" form:"is_active" binding:"required"`
}
