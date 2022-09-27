package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
	UserName  string    `gorm:"not null;column:username;index;unique;size:10" json:"user_name" form:"user_name"`
	Email     string    `gorm:"not null;unique;size:50;" json:"email" form:"email"`
	Password  string    `gorm:"not null;unique;size:60;" json:"-" form:"password"`
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
	ID        string    `gorm:"primaryKey;size:60;" json:"id"`
	UserID    *string   `gorm:"not null;unique;" json:"user_id" form:"user_id" binding:"required"`
	Token     string    `gorm:"not null;unique;" json:"token" form:"token"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

// func (obj *JwtToken) BeforeCreate(tx *gorm.DB) (err error) {
// 	id, _ := g.New()
// 	obj.ID = id
// 	return
// }

type Administrator struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
	UserID    *string   `gorm:"unique;" json:"user_id" form:"user_id"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (obj *Administrator) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Profile struct {
	ID           string     `gorm:"primaryKey;size:21;" json:"id"`
	AvatarURL    string     `json:"avatar_url" form:"avatar_url"`
	PrefixNameID *string    `json:"prefix_name_id" form:"prefix_name_id"`
	FirstName    string     `gorm:"not null;size:50;" json:"first_name" form:"first_name"`
	LastName     string     `gorm:"not null;size:50;" json:"last_name" form:"last_name"`
	UserID       *string    `gorm:"not null;unique;" json:"user_id" form:"user_id"`
	PositionID   *string    `json:"position_id" form:"position_id"`
	DepartmentID *string    `json:"department_id" form:"department_id"`
	AreaID       *string    `json:"area_id" form:"area_id"`
	WhsID        *string    `json:"whs_id" form:"whs_id"`
	FactoryID    *string    `json:"factory_id" form:"factory_id"`
	IsActive     bool       `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt    time.Time  `json:"created_at" form:"created_at" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at" form:"updated_at" default:"now"`
	User         User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Area         Area       `gorm:"foreignKey:AreaID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"area"`
	Whs          Whs        `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"whs"`
	Factory      Factory    `gorm:"foreignKey:FactoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"factory"`
	Position     Position   `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"position"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"department"`
	PrefixName   PrefixName `gorm:"foreignKey:PrefixNameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"prefix_name"`
}

func (obj *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type AuthSession struct {
	Header   string      `json:"header"`
	User     interface{} `json:"user_id,omitempty"`
	Profile  interface{} `json:"profile,omitempty"`
	JwtType  string      `json:"jwt_type,omitempty"`
	JwtToken string      `json:"jwt_token,omitempty"`
	IsAdmin  bool        `json:"is_admin,omitempty"`
}

type UserForm struct {
	UserName  string `json:"username" form:"username" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
}

type UserLoginForm struct {
	UserName   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	IsRemember bool   `json:"is_remember" form:"is_remember"`
}
