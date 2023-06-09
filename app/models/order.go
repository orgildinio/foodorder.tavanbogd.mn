package models

import (
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	CreatedAt          *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID                 int            `gorm:"column:id" json:"id"`
	IsSelled           string         `gorm:"column:is_selled" json:"is_selled"`
	InvoiceID          string         `gorm:"column:invoice_id" json:"invoice_id"`
	OrderQuantity      int            `gorm:"column:order_quantity" json:"order_quantity"`
	PaymentType        string         `gorm:"column:payment_type" json:"payment_type"`
	PaymentStatus      string         `gorm:"column:payment_status" json:"payment_status"`
	Price              int            `gorm:"column:price" json:"price"`
	UpdatedAt          *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID             int            `gorm:"column:user_id" json:"user_id"`
	OrderType          string         `gorm:"column:order_type" json:"order_type"`
	CartID             int            `gorm:"column:cart_id" json:"cart_id"`
	KitchenID          int            `gorm:"column:kitchen_id" json:"kitchen_id"`
	OrgRegisterNumber  string         `gorm:"column:org_register_number" json:"org_register_number"`
	SuccessTime        string         `gorm:"column:success_time" json:"success_time"`
	IsDelivery         *string        `gorm:"column:is_delivery" json:"is_delivery"`
	CompanyID          *int           `gorm:"column:company_id" json:"company_id"`
	EbarimtType        *string        `gorm:"column:ebarimt_type" json:"ebarimt_type"`
	EbarimtOrgRegister *int           `gorm:"column:ebarimt_org_register" json:"ebarimt_org_register"`
}

func (o *Orders) TableName() string {
	return "orders"
}

type OrderEbarimt struct {
	ID                int        `gorm:"column:id" json:"id"`
	OrderID           int        `gorm:"column:order_id" json:"order_id"`
	Ebarimt           string     `gorm:"column:ebarimt" json:"ebarimt"`
	EbarimtType       string     `gorm:"column:ebarimt_type" json:"ebarimt_type"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	OrgRegisterNumber int        `gorm:"column:org_register_number" json:"org_register_number"`
}

func (o *OrderEbarimt) TableName() string {
	return "order_ebarimt"
}

type OrderRequest struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	OrderNumber string `json:"order_number"`
}

type OrderDetail struct {
	CartID    int        `gorm:"column:cart_id" json:"cart_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	FoodID    int        `gorm:"column:food_id" json:"food_id"`
	ID        int        `gorm:"column:id" json:"id"`
	KitchenID int        `gorm:"column:kitchen_id" json:"kitchen_id"`
	Price     int        `gorm:"column:price" json:"price"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
	OrderID   int        `gorm:"column:order_id" json:"order_id"`
	Qty       int        `gorm:"column:qty" json:"qty"`
	MenuID    int        `gorm:"column:menu_id" json:"menu_id"`
	OrderType string     `gorm:"column:order_type" json:"order_type"`
}

func (o *OrderDetail) TableName() string {
	return "order_detail"
}

type OrdersStatus struct {
	ID            int            `gorm:"column:id" json:"id"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	PaymentStatus string         `gorm:"column:payment_status" json:"payment_status"`
	OrderNumber   string         `gorm:"column:order_number" json:"order_number"`
}

func (o *OrdersStatus) TableName() string {
	return "orders"
}

type ViewOrder struct {
	CartID             *int           `gorm:"column:cart_id" json:"cart_id"`
	CreatedAt          *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID                 int            `gorm:"column:id" json:"id"`
	InvoiceID          *int           `gorm:"column:invoice_id" json:"invoice_id"`
	KitchenID          *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	OrderNumber        string         `gorm:"column:order_number" json:"order_number"`
	OrderQuantity      int            `gorm:"column:order_quantity" json:"order_quantity"`
	OrderType          *string        `gorm:"column:order_type" json:"order_type"`
	PaymentStatus      *string        `gorm:"column:payment_status" json:"payment_status"`
	PaymentType        string         `gorm:"column:payment_type" json:"payment_type"`
	Price              int            `gorm:"column:price" json:"price"`
	UpdatedAt          *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID             int            `gorm:"column:user_id" json:"user_id"`
	EbarimtType        string         `gorm:"column:ebarimt_type" json:"ebarimt_type"`
	EbarimtOrgRegister string         `gorm:"column:ebarimt_org_register" json:"ebarimt_org_register"`
}

func (v *ViewOrder) TableName() string {
	return "view_order"
}

type ReceptionRequestData struct {
	ID            *int   `json:"id"`
	UserID        *int   `json:"user_id"`
	PaymentStatus string `json:"payment_status"`
}

type LutPacketPrice struct {
	ID          int     `gorm:"column:id" json:"id"`
	PacketPrice float32 `gorm:"column:packet_price" json:"packet_price"`
}

func (l *LutPacketPrice) TableName() string {
	return "lut_packet_price"
}

type OrderDetailSet struct {
	CartID    int `gorm:"column:cart_id" json:"cart_id"`
	FoodID    int `gorm:"column:food_id" json:"food_id"`
	ID        int `gorm:"column:id" json:"id"`
	KitchenID int `gorm:"column:kitchen_id" json:"kitchen_id"`
	OrderID   int `gorm:"column:order_id" json:"order_id"`
	UserID    int `gorm:"column:user_id" json:"user_id"`
	Quantity  int `gorm:"column:quantity" json:"quantity"`
}

func (o *OrderDetailSet) TableName() string {
	return "order_detail_set"
}

type ViewOrderDetail struct {
	CartID            *int    `gorm:"column:cart_id" json:"cart_id"`
	FoodIamge         *string `gorm:"column:food_iamge" json:"food_iamge"`
	FoodID            *int    `gorm:"column:food_id" json:"food_id"`
	FoodName          string  `gorm:"column:food_name" json:"food_name"`
	FoodOrderTimeName *string `gorm:"column:food_order_time_name" json:"food_order_time_name"`
	ID                *int    `gorm:"column:id" json:"id"`
	KitcheAddress     *string `gorm:"column:kitche_address" json:"kitche_address"`
	KitchenID         *int    `gorm:"column:kitchen_id" json:"kitchen_id"`
	KitchenImage      *string `gorm:"column:kitchen_image" json:"kitchen_image"`
	KitckenName       *string `gorm:"column:kitcken_name" json:"kitcken_name"`
	MenuID            *int    `gorm:"column:menu_id" json:"menu_id"`
	OrderID           int     `gorm:"column:order_id" json:"order_id"`
	Price             int     `gorm:"column:price" json:"price"`
	Qty               int     `gorm:"column:qty" json:"qty"`
	RuleImages        *string `gorm:"column:rule_images" json:"rule_images"`
	SetName           string  `gorm:"column:set_name" json:"set_name"`
}

func (v *ViewOrderDetail) TableName() string {
	return "view_order_detail"
}

type UpdateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type MyOrders struct {
	CreatedAt         *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID                int            `gorm:"column:id" json:"id"`
	IsSelled          string         `gorm:"column:is_selled" json:"is_selled"`
	InvoiceID         string         `gorm:"column:invoice_id" json:"invoice_id"`
	OrderQuantity     int            `gorm:"column:order_quantity" json:"order_quantity"`
	PaymentType       string         `gorm:"column:payment_type" json:"payment_type"`
	PaymentStatus     string         `gorm:"column:payment_status" json:"payment_status"`
	Price             int            `gorm:"column:price" json:"price"`
	UpdatedAt         *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID            int            `gorm:"column:user_id" json:"user_id"`
	OrderType         string         `gorm:"column:order_type" json:"order_type"`
	CartID            int            `gorm:"column:cart_id" json:"cart_id"`
	KitchenID         int            `gorm:"column:kitchen_id" json:"kitchen_id"`
	OrgRegisterNumber *int           `gorm:"column:org_register_number" json:"org_register_number"`
	SuccessTime       string         `gorm:"column:success_time" json:"success_time"`
	IsDelivery        *string        `gorm:"column:is_delivery" json:"is_delivery"`
	CompanyID         *int           `gorm:"column:company_id" json:"company_id"`
	UldsenSec         float32        `gorm:"column:uldsen_sec" json:"uldsen_sec"`
	OrderNumber       string         `gorm:"column:order_number" json:"order_number"`
}

func (o *MyOrders) TableName() string {
	return "orders"
}

type OrganizationInfo struct {
	VatpayerRegisteredDate string      `json:"vatpayerRegisteredDate"`
	LastReceiptDate        interface{} `json:"lastReceiptDate"`
	ReceiptFound           bool        `json:"receiptFound"`
	Name                   string      `json:"name"`
	FreeProject            bool        `json:"freeProject"`
	Citypayer              bool        `json:"citypayer"`
	Vatpayer               bool        `json:"vatpayer"`
	Found                  bool        `json:"found"`
}

type EbarimtType struct {
	EbarimtType        string `json:"ebarimt_type"`
	EbarimtOrgRegister string `json:"ebarimt_org_register"`
	OrderNumber        string `json:"order_number"`
}
