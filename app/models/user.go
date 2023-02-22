package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Avatar         *string        `gorm:"column:avatar" json:"avatar"`
	Bio            *string        `gorm:"column:bio" json:"bio"`
	Birthday       DB.Date        `gorm:"column:birthday" json:"birthday"`
	CorpID         *int           `gorm:"column:corp_id" json:"corp_id"`
	CreatedAt      *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Email          string         `gorm:"column:email" json:"email"`
	EmpID          *int           `gorm:"column:emp_id" json:"emp_id"`
	EmployeeID     *int           `gorm:"column:employee_id" json:"employee_id"`
	FcmToken       *string        `gorm:"column:fcm_token" json:"fcm_token"`
	FirstName      *string        `gorm:"column:first_name" json:"first_name"`
	Gender         *string        `gorm:"column:gender" json:"gender"`
	ID             int            `gorm:"column:id" json:"id"`
	LastName       *string        `gorm:"column:last_name" json:"last_name"`
	Login          string         `gorm:"column:login" json:"login"`
	MmkCode        *string        `gorm:"column:mmk_code" json:"mmk_code"`
	NewCode        *string        `gorm:"column:new_code" json:"new_code"`
	Password       string         `gorm:"column:password" json:"password"`
	Phone          string         `gorm:"column:phone" json:"phone"`
	RegisterNumber *string        `gorm:"column:register_number" json:"register_number"`
	Role           *int           `gorm:"column:role" json:"role"`
	Status         *string        `gorm:"column:status" json:"status"`
	UpdatedAt      *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UragOwog       *string        `gorm:"column:urag_owog" json:"urag_owog"`
}

func (d *Users) TableName() string {
	return "users"
}

type UserRegiser struct {
	LastName        string  `gorm:"column:last_name" json:"last_name"`
	FirstName       string  `gorm:"column:first_name" json:"first_name"`
	RegisterNumber  string  `gorm:"column:register_number" json:"register_number"`
	Email           string  `gorm:"column:email" json:"email"`
	Phone           string  `gorm:"column:phone" json:"phone"`
	Password        *string `gorm:"column:password" json:"password"`
	ConfirmPassword *string `gorm:"column:confirm_password" json:"confirm_password"`
	ConfirmCode     string  `gorm:"column:confirm_code" json:"confirm_code"`
	Login           string  `gorm:"column:login" json:"login"`
}
