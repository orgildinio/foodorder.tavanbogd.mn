package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type TblMenuGtHoer struct {
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodTypeID  *int           `gorm:"column:food_type_id" json:"food_type_id"`
	ID          int            `gorm:"column:id" json:"id"`
	MainMenuID  *int           `gorm:"column:main_menu_id" json:"main_menu_id"`
	OrderRuleID *int           `gorm:"column:order_rule_id" json:"order_rule_id"`
	SetDate     DB.Date        `gorm:"column:set_date" json:"set_date"`
	SetName     *string        `gorm:"column:set_name" json:"set_name"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID      *int           `gorm:"column:user_id" json:"user_id"`
}

func (t *TblMenuGtHoer) TableName() string {
	return "tbl_menu_gt_hoer"
}

type SubMenuGtHoer struct {
	FoodTypeID *int `gorm:"column:food_type_id" json:"food_type_id"`
	ID         int  `gorm:"column:id" json:"id"`
	MenuID     *int `gorm:"column:menu_id" json:"menu_id"`
}

func (s *SubMenuGtHoer) TableName() string {
	return "sub_menu_gt_hoer"
}

type SubMenuFoodsGtHoer struct {
	FoodID       *int     `gorm:"column:food_id" json:"food_id"`
	FoodQuantity *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	ID           int      `gorm:"column:id" json:"id"`
	SubMenuID    *int     `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (s *SubMenuFoodsGtHoer) TableName() string {
	return "sub_menu_foods_gt_hoer"
}
