package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	MenuID      int    `gorm:"column:menu_id" json:"menu_id"`
	OrderStatus string `gorm:"column:order_status" json:"order_status"`
	GtBranch    string `gorm:"column:gt_branch" json:"gt_branch"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`
	//OrderNumber string `gorm:"column:order_number" json:"order_number"`
	//CancelledAt *time.Time     `gorm:"column:cancelled_at" json:"cancelled_at"`
	//CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	//DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	//ID          int            `gorm:"column:id" json:"id"`
	//PaymentAt   *time.Time     `gorm:"column:payment_at" json:"payment_at"`
	//PaymentType *string        `gorm:"column:payment_type" json:"payment_type"`
	//UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	//UserID      *int           `gorm:"column:user_id" json:"user_id"`
}

func (o *Orders) TableName() string {
	return "orders"
}

type ActiveOrder struct {
	CancelledAt       time.Time      `gorm:"column:cancelled_at" json:"cancelled_at"`
	CreatedAt         *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	GtBranch          *string        `gorm:"column:gt_branch" json:"gt_branch"`
	ID                *int           `gorm:"column:id" json:"id"`
	MenuID            *int           `gorm:"column:menu_id" json:"menu_id"`
	MorningOrderEnd   *string        `gorm:"column:morning_order_end" json:"morning_order_end"`
	MorningOrderStart *string        `gorm:"column:morning_order_start" json:"morning_order_start"`
	OrderCancelTime   *string        `gorm:"column:order_cancel_time" json:"order_cancel_time"`
	OrderNumber       *string        `gorm:"column:order_number" json:"order_number"`
	OrderRuleID       *int           `gorm:"column:order_rule_id" json:"order_rule_id"`
	OrderStatus       *string        `gorm:"column:order_status" json:"order_status"`
	PaymentAt         *time.Time     `gorm:"column:payment_at" json:"payment_at"`
	PaymentType       *string        `gorm:"column:payment_type" json:"payment_type"`
	SetDate           DB.Date        `gorm:"column:set_date" json:"set_date"`
	UpdatedAt         *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID            *int           `gorm:"column:user_id" json:"user_id"`
}

func (a *ActiveOrder) TableName() string {
	return "active_order"
}

type AcitiveMenu struct {
	FoodName          *string  `gorm:"column:food_name" json:"food_name"`
	ID                int      `gorm:"column:id" json:"id"`
	MorningOrderEnd   *string  `gorm:"column:morning_order_end" json:"morning_order_end"`
	MorningOrderStart *string  `gorm:"column:morning_order_start" json:"morning_order_start"`
	OrderCancelTime   *string  `gorm:"column:order_cancel_time" json:"order_cancel_time"`
	QuantityGtNeg     *float32 `gorm:"column:quantity_gt_neg" json:"quantity_gt_neg"`
	SetDate           DB.Date  `gorm:"column:set_date" json:"set_date"`
}

func (a *AcitiveMenu) TableName() string {
	return "acitive_menu"
}

type TblOrderRule struct {
	CanecelingTime    string  `gorm:"column:caneceling_time" json:"caneceling_time"`
	FoodOrderTimeName *string `gorm:"column:food_order_time_name" json:"food_order_time_name"`
	ID                int     `gorm:"column:id" json:"id"`
	MorningOrderEnd   *string `gorm:"column:morning_order_end" json:"morning_order_end"`
	MorningOrderStart *string `gorm:"column:morning_order_start" json:"morning_order_start"`

	UserID *int `gorm:"column:user_id" json:"user_id"`
}

func (t *TblOrderRule) TableName() string {
	return "tbl_order_rule"
}
