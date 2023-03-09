package models

import (
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
