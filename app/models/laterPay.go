package models

import "time"

type OrderLaterPay struct {
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	ID            int        `gorm:"column:id" json:"id"`
	OrderID       int        `gorm:"column:order_id" json:"order_id"`
	OrderNumber   string     `gorm:"column:order_number" json:"order_number"`
	PaymentStatus string     `gorm:"column:payment_status" json:"payment_status"`
	Price         int        `gorm:"column:price" json:"price"`
	Qty           int        `gorm:"column:qty" json:"qty"`
	SuccessDate   time.Time  `gorm:"column:success_date" json:"success_date"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID        int        `gorm:"column:user_id" json:"user_id"`
}

func (o *OrderLaterPay) TableName() string {
	return "order_later_pay"
}
