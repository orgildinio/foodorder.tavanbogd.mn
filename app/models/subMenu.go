package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type TblMenu struct {
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodTypeID    *int           `gorm:"column:food_type_id" json:"food_type_id"`
	ID            int            `gorm:"column:id" json:"id"`
	KitchenID     *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	OrderBranchID *int           `gorm:"column:order_branch_id" json:"order_branch_id"`
	OrderRuleID   *int           `gorm:"column:order_rule_id" json:"order_rule_id"`
	OrderUserID   *int           `gorm:"column:order_user_id" json:"order_user_id"`
	PacketPriceID *int           `gorm:"column:packet_price_id" json:"packet_price_id"`
	SetDate       DB.Date        `gorm:"column:set_date" json:"set_date"`
	SetName       *string        `gorm:"column:set_name" json:"set_name"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID        *int           `gorm:"column:user_id" json:"user_id"`
}

func (t *TblMenu) TableName() string {
	return "tbl_menu"
}

type RequestFoodData struct {
	Data struct {
		CreatedAt   time.Time   `json:"created_at"`
		DeletedAt   interface{} `json:"deleted_at"`
		ID          int         `json:"id"`
		OrderRuleID int         `json:"order_rule_id"`
		SetDate     interface{} `json:"set_date"`
		SetName     string      `json:"set_name"`
		SubMenu     []struct {
			FoodTypeID   int `json:"food_type_id"`
			ID           int `json:"id"`
			MenuID       int `json:"menu_id"`
			SubMenuFoods []struct {
				FoodID    int `json:"food_id"`
				ID        int `json:"id"`
				SubMenuID int `json:"sub_menu_id"`
			} `json:"sub_menu_foods"`
		} `json:"sub_menu"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    int       `json:"user_id"`
	} `json:"data"`
	Status bool `json:"status"`
}

type SubMenu struct {
	FoodTypeID *int `gorm:"column:food_type_id" json:"food_type_id"`
	ID         int  `gorm:"column:id" json:"id"`
	MenuID     int  `gorm:"column:menu_id" json:"menu_id"`
}

func (s *SubMenu) TableName() string {
	return "sub_menu"
}

type SubMenuFoods struct {
	FoodID    *int `gorm:"column:food_id" json:"food_id"`
	ID        int  `gorm:"column:id" json:"id"`
	SubMenuID int  `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (s *SubMenuFoods) TableName() string {
	return "sub_menu_foods"
}
