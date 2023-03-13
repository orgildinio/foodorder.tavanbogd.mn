package models

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	Qty       *int           `gorm:"column:qty" json:"qty"`
	CartType  *string        `gorm:"column:cart_type" json:"cart_type"`
	CreatedAt *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int            `gorm:"column:id" json:"id"`
	KitchenID *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	MenuID    *int           `gorm:"column:menu_id" json:"menu_id"`
	//OrderNumber *string        `gorm:"column:order_number" json:"order_number"`
	OrderStatus *string    `gorm:"column:order_status" json:"order_status"`
	PaymentAt   *time.Time `gorm:"column:payment_at" json:"payment_at"`
	PaymentType *string    `gorm:"column:payment_type" json:"payment_type"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID      int        `gorm:"column:user_id" json:"user_id"`
	CartID      *int       `gorm:"column:cart_id" json:"cart_id"`
	//InvoiceId   string         `gorm:"column:invoice_id" json:"invoice_id"`
}

func (o *Orders) TableName() string {
	return "orders"
}

type ViewOrders struct {
	InvoiceID      *string        `gorm:"column:Invoice_id" json:"Invoice_id"`
	CancelledAt    *time.Time     `gorm:"column:cancelled_at" json:"cancelled_at"`
	CartID         *int           `gorm:"column:cart_id" json:"cart_id"`
	CartType       *string        `gorm:"column:cart_type" json:"cart_type"`
	CreatedAt      *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodBalanceQty *int           `gorm:"column:food_balance_qty" json:"food_balance_qty"`
	FoodID         *int           `gorm:"column:food_id" json:"food_id"`
	ID             *int           `gorm:"column:id" json:"id"`
	KitchenID      *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	MenuID         *int           `gorm:"column:menu_id" json:"menu_id"`
	OrderNumber    string         `gorm:"column:order_number" json:"order_number"`
	OrderQty       *int           `gorm:"column:order_qty" json:"order_qty"`
	OrderStatus    *string        `gorm:"column:order_status" json:"order_status"`
	PaymentAt      *time.Time     `gorm:"column:payment_at" json:"payment_at"`
	PaymentStatus  *string        `gorm:"column:payment_status" json:"payment_status"`
	PaymentType    *string        `gorm:"column:payment_type" json:"payment_type"`
	TotalPrice     *int           `gorm:"column:total_price" json:"total_price"`
	UpdatedAt      *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID         *int           `gorm:"column:user_id" json:"user_id"`
}

func (v *ViewOrders) TableName() string {
	return "view_orders"
}

type CheckOrders struct {
	CartType    *string        `gorm:"column:cart_type" json:"cart_type"`
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID          int            `gorm:"column:id" json:"id"`
	KitchenID   *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	MenuID      *int           `gorm:"column:menu_id" json:"menu_id"`
	OrderNumber *string        `gorm:"column:order_number" json:"order_number"`
	OrderStatus *string        `gorm:"column:order_status" json:"order_status"`
	PaymentAt   *time.Time     `gorm:"column:payment_at" json:"payment_at"`
	PaymentType *string        `gorm:"column:payment_type" json:"payment_type"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID      int            `gorm:"column:user_id" json:"user_id"`
	CartID      *int           `gorm:"column:cart_id" json:"cart_id"`
	//InvoiceId   string         `gorm:"column:invoice_id" json:"invoice_id"`
}

func (o *CheckOrders) TableName() string {
	return "orders"
}
