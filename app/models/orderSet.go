package models

import (
	"gorm.io/gorm"
	"time"
)

type OrderSet struct {
	CancelledAt *time.Time     `gorm:"column:cancelled_at" json:"cancelled_at"`
	CartID      int            `gorm:"column:cart_id" json:"cart_id"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID          int            `gorm:"column:id" json:"id"`
	MenuID      *int           `gorm:"column:menu_id" json:"menu_id"`
	//OrderNumber *string        `gorm:"column:order_number" json:"order_number"`
	PaymentType *string    `gorm:"column:payment_type" json:"payment_type"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID      *int       `gorm:"column:user_id" json:"user_id"`
	OrderStatus *string    `gorm:"column:order_status" json:"order_status"`
}

func (o *OrderSet) TableName() string {
	return "order_set"
}

type CartSubMenu struct {
	FoodTypeID *int `gorm:"column:food_type_id" json:"food_type_id"`
	ID         int  `gorm:"column:id" json:"id"`
	MenuID     *int `gorm:"column:menu_id" json:"menu_id"`
}

func (c *CartSubMenu) TableName() string {
	return "cart_sub_menu"
}

type CartSubMenuFood struct {
	FoodID    *int `gorm:"column:food_id" json:"food_id"`
	ID        int  `gorm:"column:id" json:"id"`
	SubMenuID *int `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (c *CartSubMenuFood) TableName() string {
	return "cart_sub_menu_food"
}

type OrderSetCheck struct {
	CancelledAt *time.Time     `gorm:"column:cancelled_at" json:"cancelled_at"`
	CartID      int            `gorm:"column:cart_id" json:"cart_id"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID          int            `gorm:"column:id" json:"id"`
	MenuID      *int           `gorm:"column:menu_id" json:"menu_id"`
	OrderNumber *string        `gorm:"column:order_number" json:"order_number"`
	PaymentType *string        `gorm:"column:payment_type" json:"payment_type"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID      *int           `gorm:"column:user_id" json:"user_id"`
	OrderStatus *string        `gorm:"column:order_status" json:"order_status"`
}

func (o *OrderSetCheck) TableName() string {
	return "order_set"
}
