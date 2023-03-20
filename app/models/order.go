package models

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID            int            `gorm:"column:id" json:"id"`
	InvoiceID     int            `gorm:"column:invoice_id" json:"invoice_id"`
	OrderQuantity int            `gorm:"column:order_quantity" json:"order_quantity"`
	PaymentType   string         `gorm:"column:payment_type" json:"payment_type"`
	PaymentStatus string         `gorm:"column:payment_status" json:"payment_status"`
	Price         int            `gorm:"column:price" json:"price"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID        int            `gorm:"column:user_id" json:"user_id"`
	OrderType     string         `gorm:"column:order_type" json:"order_type"`
}

func (o *Orders) TableName() string {
	return "orders"
}

type OrderRequest struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	OrderNumber string `json:"order_number"`
}

type OrderDetail struct {
	CartID    *int       `gorm:"column:cart_id" json:"cart_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	FoodID    *int       `gorm:"column:food_id" json:"food_id"`
	ID        int        `gorm:"column:id" json:"id"`
	KitchenID int        `gorm:"column:kitchen_id" json:"kitchen_id"`
	Price     int        `gorm:"column:price" json:"price"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	OrderID   int        `gorm:"column:order_id" json:"order_id"`
	Qty       int        `gorm:"column:qty" json:"qty"`
}

func (o *OrderDetail) TableName() string {
	return "order_detail"
}

type OrdersStatus struct {
	ID            int    `gorm:"column:id" json:"id"`
	PaymentStatus string `gorm:"column:payment_status" json:"payment_status"`
	OrderNumber   string `gorm:"column:order_number"json:"order_number"`
}

func (o *OrdersStatus) TableName() string {
	return "orders"
}
