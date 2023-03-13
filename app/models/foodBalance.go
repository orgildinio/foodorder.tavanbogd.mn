package models

import (
	"gorm.io/gorm"
	"time"
)

type FoodBalance struct {
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodID    int            `gorm:"column:food_id" json:"food_id"`
	FoodPrice *float32       `gorm:"column:food_price" json:"food_price"`
	ID        int            `gorm:"column:id" json:"id"`
	KitchenID int            `gorm:"column:kitchen_id" json:"kitchen_id"`
	Qty       int            `gorm:"column:qty" json:"qty"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (f *FoodBalance) TableName() string {
	return "food_balance"
}

type FoodBalanced struct {
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodID    *int           `gorm:"column:food_id" json:"food_id"`
	FoodPrice *float32       `gorm:"column:food_price" json:"food_price"`
	ID        int            `gorm:"column:id" json:"id"`
	KitchenID int            `gorm:"column:kitchen_id" json:"kitchen_id"`
	Qty       *float32       `gorm:"column:qty" json:"qty"`
	UpdatedAt *time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (f *FoodBalanced) TableName() string {
	return "food_balance"
}
