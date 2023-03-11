package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type CartZahialgat struct {
	CartOrderNumber *string        `gorm:"column:cart_order_number" json:"cart_order_number"`
	CreatedAt       *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodID          *int           `gorm:"column:food_id" json:"food_id"`
	ID              int            `gorm:"column:id" json:"id"`
	KitchenID       *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	Ognoo           *time.Time     `gorm:"column:ognoo" json:"ognoo"`
	Price           *int           `gorm:"column:price" json:"price"`
	Qty             int            `gorm:"column:qty" json:"qty"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID          *int           `gorm:"column:user_id" json:"user_id"`
}

func (c *CartZahialgat) TableName() string {
	return "cart_zahialgat"
}

type SetFoodForCart struct {
	FoodID            *int    `gorm:"column:food_id" json:"food_id"`
	FoodName          *string `gorm:"column:food_name" json:"food_name"`
	FoodOrderTimeName *string `gorm:"column:food_order_time_name" json:"food_order_time_name"`
	ID                *int    `gorm:"column:id" json:"id"`
	MorningOrderEnd   *string `gorm:"column:morning_order_end" json:"morning_order_end"`
	MorningOrderStart *string `gorm:"column:morning_order_start" json:"morning_order_start"`
	SetDate           DB.Date `gorm:"column:set_date" json:"set_date"`
	SetName           *string `gorm:"column:set_name" json:"set_name"`
	SetQty            *int    `gorm:"column:set_qty" json:"set_qty"`
	SubMenuFoodID     *int    `gorm:"column:sub_menu_food_id" json:"sub_menu_food_id"`
}

func (s *SetFoodForCart) TableName() string {
	return "set_food_for_cart"
}

type CartSet struct {
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodID        *int           `gorm:"column:food_id" json:"food_id"`
	ID            int            `gorm:"column:id" json:"id"`
	KitchenID     *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	MenuID        *int           `gorm:"column:menu_id" json:"menu_id"`
	Ognoo         *time.Time     `gorm:"column:ognoo" json:"ognoo"`
	Price         *int           `gorm:"column:price" json:"price"`
	Qty           int            `gorm:"column:qty" json:"qty"`
	SubMenuFoodID *int           `gorm:"column:sub_menu_food_id" json:"sub_menu_food_id"`
	SubMenuID     *int           `gorm:"column:sub_menu_id" json:"sub_menu_id"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID        *int           `gorm:"column:user_id" json:"user_id"`
}

func (c *CartSet) TableName() string {
	return "cart_set"
}

type ViewCartSet struct {
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FirstName     *string        `gorm:"column:first_name" json:"first_name"`
	FoodID        *int           `gorm:"column:food_id" json:"food_id"`
	FoodName      *string        `gorm:"column:food_name" json:"food_name"`
	FoodType      *string        `gorm:"column:food_type" json:"food_type"`
	ID            int            `gorm:"column:id" json:"id"`
	KitchenID     *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	LastName      *string        `gorm:"column:last_name" json:"last_name"`
	MenuID        *int           `gorm:"column:menu_id" json:"menu_id"`
	Ognoo         *time.Time     `gorm:"column:ognoo" json:"ognoo"`
	Price         *int           `gorm:"column:price" json:"price"`
	Qty           int            `gorm:"column:qty" json:"qty"`
	SetDate       DB.Date        `gorm:"column:set_date" json:"set_date"`
	SetName       *string        `gorm:"column:set_name" json:"set_name"`
	SubMenuFoodID *int           `gorm:"column:sub_menu_food_id" json:"sub_menu_food_id"`
	SubMenuID     *int           `gorm:"column:sub_menu_id" json:"sub_menu_id"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID        *int           `gorm:"column:user_id" json:"user_id"`
}

func (v *ViewCartSet) TableName() string {
	return "view_cart_set"
}
