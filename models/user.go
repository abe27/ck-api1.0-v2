package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;unique;index;size:21;" json:"id" form:"id"`
	FirstName string    `gorm:"size:50;" json:"first_name" form:"first_name"`
	LastName  string    `gorm:"size:50;" json:"last_name" form:"last_name"`
	UserName  string    `gorm:"unique;size:21;" json:"user_name" form:"user_name"`
	Email     string    `gorm:"unique;size:21;" json:"email" form:"email"`
	Password  string    `gorm:"unique;size:60;" json:"password" form:"password"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" default:"now"`
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type JwtToken struct {
	ID        string    `gorm:"primaryKey;unique;index;size:21;" json:"id" form:"id"`
	UserID    string    `gorm:"unique;size:21;" json:"user_id" form:"user_id"`
	Token     string    `json:"token" form:"token"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" default:"now"`
	User      User      `gorm:"foreignKey:AreaID;references:ID;constraint:OnDelete:CASCADE;" json:"user"`
}

func (obj *JwtToken) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type AuthSession struct {
	UserID    string `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
