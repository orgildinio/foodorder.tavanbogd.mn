package models

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	MenuID      int    `gorm:"column:menu_id" json:"menu_id"`
	OrderStatus string `gorm:"column:order_status" json:"order_status"`
	GtBranch    string `gorm:"column:gt_branch" json:"gt_branch"`
	UserID      int    `gorm:"column:user_id" json:"user_id"`
	//CancelledAt *time.Time     `gorm:"column:cancelled_at" json:"cancelled_at"`
	//CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	//DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	//ID          int            `gorm:"column:id" json:"id"`

	//OrderNumber *string `gorm:"column:order_number" json:"order_number"`

	//PaymentAt   *time.Time     `gorm:"column:payment_at" json:"payment_at"`
	//PaymentType *string        `gorm:"column:payment_type" json:"payment_type"`
	//UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	//UserID      *int           `gorm:"column:user_id" json:"user_id"`
}

func (o *Orders) TableName() string {
	return "orders"
}

type OrdersView struct {
	CancelledAt *time.Time     `gorm:"column:cancelled_at" json:"cancelled_at"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	GtBranch    *string        `gorm:"column:gt_branch" json:"gt_branch"`
	ID          *int           `gorm:"column:id" json:"id"`
	MenuID      *int           `gorm:"column:menu_id" json:"menu_id"`
	OrderNumber *string        `gorm:"column:order_number" json:"order_number"`
	OrderStatus *string        `gorm:"column:order_status" json:"order_status"`
	PaymentAt   *time.Time     `gorm:"column:payment_at" json:"payment_at"`
	PaymentType *string        `gorm:"column:payment_type" json:"payment_type"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID      *int           `gorm:"column:user_id" json:"user_id"`
}

func (o *OrdersView) TableName() string {
	return "orders_view"
}
