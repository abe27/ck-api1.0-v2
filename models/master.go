package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Area struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Area) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Whs struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Value       string    `gorm:"size:5;" json:"value,omitempty" form:"value" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Factory struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
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
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Factory) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PrefixName struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *PrefixName) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Position struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Position) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Department struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Department) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Unit struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PartType struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *PartType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ReceiveType struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	WhsID       string    `gorm:"not null;" json:"whs_id,omitempty" form:"whs_id"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Whs         Whs       `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"whs,omitempty"`
}

func (obj *ReceiveType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type FileType struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:50" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *FileType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Mailbox struct {
	ID        string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Mailbox   string    `gorm:"not null;unique;size:21" json:"mailbox,omitempty" form:"mailbox" binding:"required"`
	Password  string    `gorm:"size:50" json:"password,omitempty" form:"password" binding:"required"`
	HostUrl   string    `json:"host_url,omitempty" form:"host_url" binding:"required"`
	AreaID    *string   `json:"area_id,omitempty" form:"area_id" binding:"required"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
	Area      Area      `gorm:"foreignKey:AreaID;references:ID;constraint:OnDelete:CASCADE;" json:"area,omitempty"`
}

func (obj *Mailbox) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Pc struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:50" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Pc) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Commercial struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:50" json:"title,omitempty" form:"title" binding:"required"`
	Prefix      string    `gorm:"size:5" json:"prefix,omitempty" form:"prefix" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Commercial) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type SampleFlg struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:50" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *SampleFlg) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ReviseOrder struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *ReviseOrder) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Shipment struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Shipment) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderType struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *OrderType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

// N=All,F=3 Front,E=3 End,O=Sprit Order
type OrderGroupType struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Title       string    `gorm:"not null;unique;size:15" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *OrderGroupType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OrderGroup struct {
	ID               string          `gorm:"primaryKey;size:21" json:"id,omitempty"`
	UserID           *string         `gorm:"not null;" json:"user_id,omitempty" form:"user_id" binding:"required"`
	ConsigneeID      *string         `gorm:"not null;" json:"consignee_id,omitempty" form:"consignee_id" binding:"required"`
	OrderGroupTypeID *string         `gorm:"not null;" json:"order_group_type_id,omitempty" form:"order_group_type_id" binding:"required"`
	SubOrder         string          `gorm:"not null;size:15" json:"sub_order,omitempty" form:"sub_order" binding:"required"`
	Description      string          `json:"description,omitempty" form:"description" binding:"required"`
	IsActive         bool            `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt        time.Time       `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt        time.Time       `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	User             *User           `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Consignee        *Consignee      `gorm:"foreignKey:ConsigneeID;references:ID" json:"consignee,omitempty"`
	OrderGroupType   *OrderGroupType `gorm:"foreignKey:OrderGroupTypeID;references:ID" json:"order_group_type,omitempty"`
}

func (obj *OrderGroup) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type FormGroupConsignee struct {
	UserID           string `json:"user_id,omitempty" form:"user_id" binding:"required"`
	WhsID            string `json:"whs_id,omitempty" form:"whs_id"`
	FactoryID        string `json:"factory_id,omitempty" form:"factory_id"`
	AffcodeID        string `json:"affcode_id,omitempty" form:"affcode_id" binding:"required"`
	CustcodeID       string `json:"custcode_id,omitempty" form:"custcode_id" binding:"required"`
	OrderGroupTypeID string `json:"order_group_type_id,omitempty" form:"order_group_type_id" binding:"required"`
	SubOrder         string `gorm:"size:15" json:"sub_order,omitempty" form:"sub_order" binding:"required"`
	Description      string `json:"description,omitempty" form:"description" binding:"required"`
	IsActive         bool   `json:"is_active,omitempty" form:"is_active" binding:"required"`
}

type Location struct {
	ID          string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `gorm:"null" json:"description,omitempty" form:"description"`
	MaxLimit    int64     `gorm:"null" json:"max_limit,omitempty" form:"max_limit" binding:"required"`
	IsActive    bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
}

func (obj *Location) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type AutoGenerateInvoice struct {
	ID         string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	FactoryID  *string   `gorm:"not null;unique;" json:"factory_id,omitempty" form:"factory_id,omitempty"`
	IsGenerate bool      `json:"is_generate" form:"is_generate" default:"true"`
	IsActive   bool      `json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt  time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Factory    Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
}

func (obj *AutoGenerateInvoice) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type LineNotifyToken struct {
	ID        string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	WhsID     *string   `gorm:"not null;" json:"whs_id,omitempty" form:"whs_id,omitempty"`
	FactoryID *string   `gorm:"not null;" json:"factory_id,omitempty" form:"factory_id,omitempty"`
	Token     string    `gorm:"not null;unique;" json:"token,omitempty" form:"token"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Whs       Whs       `gorm:"foreignKey:WhsID;references:ID" json:"whs,omitempty"`
	Factory   Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
}

func (obj *LineNotifyToken) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PalletType struct {
	ID               string    `gorm:"primaryKey,unique;size:21;" json:"id,omitempty"`
	Type             string    `gorm:"not null;unique;" json:"type,omitempty" form:"type"`
	Floors           int64     `json:"floors" form:"floors" default:"0"`
	BoxSizeWidth     float64   `json:"box_size_width" form:"box_size_width" default:"0"`
	BoxSizeLength    float64   `json:"box_size_length" form:"box_size_length" default:"0"`
	BoxSizeHight     float64   `json:"box_size_hight" form:"box_size_hight" default:"0"`
	PalletSizeWidth  float64   `json:"pallet_size_width" form:"pallet_size_width" default:"0"`
	PalletSizeLength float64   `json:"pallet_size_length" form:"pallet_size_length" default:"0"`
	PalletSizeHight  float64   `json:"pallet_size_hight" form:"pallet_size_hight" default:"0"`
	LimitTotal       int64     `json:"limit_total" form:"limit_total" default:"0"`
	IsActive         bool      `json:"is_activey" form:"is_active" default:"true"`
	CreatedAt        time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt        time.Time `json:"updated_at" form:"updated_at" default:"now"`
	// ชนดิ กลอ่ ง จ านวนชนั้ BOX SIZE PALLET SIZE BOX/PALLET
}

func (obj *PalletType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type LastFticket struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	FactoryID   *string   `gorm:"not null;unique;" json:"factory_id,omitempty" form:"factory_id" binding:"required"`
	OnYear      int64     `gorm:"not null;" json:"on_year,omitempty" form:"on_year"`
	LastRunning int64     `json:"last_running,omitempty" form:"last_running" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	Factory     Factory   `gorm:"foreignKey:FactoryID;references:ID" json:"factory,omitempty"`
}

func (obj *LastFticket) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
