package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type ImportInvoiceTap struct {
	ID          string    `gorm:"primaryKey;size:21" json:"id,omitempty"`
	Bhivno      string    `gorm:"not null;size:10" json:"bhivno,omitempty" form:"bhivno" binding:"required"`
	Bhodpo      string    `gorm:"not null;size:25" json:"bhodpo,omitempty" form:"bhodpo" binding:"required"`
	Bhivdt      time.Time `gorm:"not null;" json:"bhivdt,omitempty" form:"bhivdt" binding:"required"`
	Bhconn      string    `gorm:"size:50;" json:"bhconn,omitempty" form:"bhconn"`
	Bhcons      string    `gorm:"size:50" json:"bhconns,omitempty" form:"bhconns"`
	Bhsven      string    `gorm:"size:50" json:"bhsven,omitempty" form:"bhsven"`
	Bhshpf      string    `gorm:"size:50" json:"bhshpf,omitempty" form:"bhshpf"`
	Bhsafn      string    `gorm:"size:50" json:"bhafn,omitempty" form:"bhafn" binding:"required"`
	Bhshpt      string    `gorm:"size:50" json:"bhshpt,omitempty" form:"bhshpt"`
	Bhfrtn      string    `gorm:"size:50" json:"bhfrtn,omitempty" form:"bhfrtn"`
	Bhcon       int64     `json:"bhcon,omitempty" form:"bhcon"`
	Bhpaln      string    `json:"bhpaln,omitempty" form:"bhpaln"`
	Bhpnam      string    `json:"bhpnam,omitempty" form:"bhpnam"`
	Bhypat      string    `json:"bhypat,omitempty" form:"bhypat" binding:"required"`
	Bhctn       int64     `json:"bhctn,omitempty" form:"bhctn" binding:"required"`
	Bhwidt      int64     `json:"bhwidt,omitempty" form:"bhwidt"`
	Bhleng      int64     `json:"bhleng,omitempty" form:"bhleng"`
	Bhhigh      int64     `json:"bhhigh,omitempty" form:"bhhigh"`
	Bhgrwt      float64   `json:"bhgrwt,omitempty" form:"bhgrwt"`
	Bhcbmt      float64   `json:"bhcbmt,omitempty" form:"bhcbmt"`
	OrderPlanID *string   `json:"order_plan_id,omitempty" form:"order_plan_id"`
	IsMatched   bool      `json:"is_matched,omitempty" form:"is_matched" default:"true"`
	IsChecked   bool      `json:"is_checked,omitempty" form:"is_checked" default:"true"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" default:"true"`
	CreatedAt   time.Time `json:"created_at,omitempty" form:"created_at" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at" default:"now"`
	OrderPlan   OrderPlan `gorm:"foreignKey:OrderPlanID;references:ID" json:"order_plan,omitempty"`
}

func (u *ImportInvoiceTap) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	u.ID = id
	return nil
}
