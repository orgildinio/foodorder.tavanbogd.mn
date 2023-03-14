package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type TblMenu struct {
	CreatedAt     *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodTypeID    *int           `gorm:"column:food_type_id" json:"food_type_id"`
	ID            *int           `gorm:"column:id" json:"id"`
	KitchenID     *int           `gorm:"column:kitchen_id" json:"kitchen_id"`
	OrderBranchID *int           `gorm:"column:order_branch_id" json:"order_branch_id"`
	OrderRuleID   *int           `gorm:"column:order_rule_id" json:"order_rule_id"`
	OrderUserID   *int           `gorm:"column:order_user_id" json:"order_user_id"`
	PacketPriceID *int           `gorm:"column:packet_price_id" json:"packet_price_id"`
	SetDate       DB.Date        `gorm:"column:set_date" json:"set_date"`
	SetName       *string        `gorm:"column:set_name" json:"set_name"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID        *int           `gorm:"column:user_id" json:"user_id"`
}

func (t *TblMenu) TableName() string {
	return "tbl_menu"
}

type RequestFoodData struct {
	Data struct {
		CreatedAt   time.Time   `json:"created_at"`
		DeletedAt   interface{} `json:"deleted_at"`
		ID          int         `json:"id"`
		OrderRuleID int         `json:"order_rule_id"`
		SetDate     interface{} `json:"set_date"`
		SetName     string      `json:"set_name"`
		SubMenu     []struct {
			FoodTypeID   int `json:"food_type_id"`
			ID           int `json:"id"`
			MenuID       int `json:"menu_id"`
			SubMenuFoods []struct {
				FoodID    int `json:"food_id"`
				ID        int `json:"id"`
				SubMenuID int `json:"sub_menu_id"`
			} `json:"sub_menu_foods"`
		} `json:"sub_menu"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    int       `json:"user_id"`
	} `json:"data"`
	Status bool `json:"status"`
}

type SubMenu struct {
	FoodTypeID *int `gorm:"column:food_type_id" json:"food_type_id"`
	ID         int  `gorm:"column:id" json:"id"`
	MenuID     int  `gorm:"column:menu_id" json:"menu_id"`
}

func (s *SubMenu) TableName() string {
	return "sub_menu"
}

type SubMenuFoods struct {
	FoodID    *int `gorm:"column:food_id" json:"food_id"`
	ID        int  `gorm:"column:id" json:"id"`
	SubMenuID int  `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (s *SubMenuFoods) TableName() string {
	return "sub_menu_foods"
}

type ViewCartSetFood struct {
	DesertBalanceID *int       `gorm:"column:desert_balance_id" json:"desert_balance_id"`
	DesertFoodID    *int       `gorm:"column:desert_food_id" json:"desert_food_id"`
	DesertFoodType  *string    `gorm:"column:desert_food_type" json:"desert_food_type"`
	DesertImage     *string    `gorm:"column:desert_image" json:"desert_image"`
	DesertKitchenID *int       `gorm:"column:desert_kitchen_id" json:"desert_kitchen_id"`
	DesertName      *string    `gorm:"column:desert_name" json:"desert_name"`
	DesertQty       *float32   `gorm:"column:desert_qty" json:"desert_qty"`
	FoodIDDesert    *int       `gorm:"column:food_id_desert" json:"food_id_desert"`
	FoodIDHoer      *int       `gorm:"column:food_id_hoer" json:"food_id_hoer"`
	FoodIDNeg       *int       `gorm:"column:food_id_neg" json:"food_id_neg"`
	FoodIDUuh       *int       `gorm:"column:food_id_uuh" json:"food_id_uuh"`
	FoodIDZuush     *int       `gorm:"column:food_id_zuush" json:"food_id_zuush"`
	HoerBalanceID   *int       `gorm:"column:hoer_balance_id" json:"hoer_balance_id"`
	HoerFoodID      *int       `gorm:"column:hoer_food_id" json:"hoer_food_id"`
	HoerFoodName    *string    `gorm:"column:hoer_food_name" json:"hoer_food_name"`
	HoerFoodType    *string    `gorm:"column:hoer_food_type" json:"hoer_food_type"`
	HoerImage       *string    `gorm:"column:hoer_image" json:"hoer_image"`
	HoerKitchenID   *int       `gorm:"column:hoer_kitchen_id" json:"hoer_kitchen_id"`
	HoerQty         *float32   `gorm:"column:hoer_qty" json:"hoer_qty"`
	ID              int        `gorm:"column:id" json:"id"`
	KitchenID       *int       `gorm:"column:kitchen_id" json:"kitchen_id"`
	MenuID          *int       `gorm:"column:menu_id" json:"menu_id"`
	NegBalanceID    *int       `gorm:"column:neg_balance_id" json:"neg_balance_id"`
	NegFoodID       *int       `gorm:"column:neg_food_id" json:"neg_food_id"`
	NegFoodImage    *string    `gorm:"column:neg_food_image" json:"neg_food_image"`
	NegFoodName     *string    `gorm:"column:neg_food_name" json:"neg_food_name"`
	NegFoodType     *string    `gorm:"column:neg_food_type" json:"neg_food_type"`
	NegKitchenID    *int       `gorm:"column:neg_kitchen_id" json:"neg_kitchen_id"`
	NegQty          *float32   `gorm:"column:neg_qty" json:"neg_qty"`
	Ognoo           *time.Time `gorm:"column:ognoo" json:"ognoo"`
	PacketPrice     *float32   `gorm:"column:packet_price" json:"packet_price"`
	Price           *int       `gorm:"column:price" json:"price"`
	Qty             *int       `gorm:"column:qty" json:"qty"`
	RuleImages      *string    `gorm:"column:rule_images" json:"rule_images"`
	SetDate         DB.Date    `gorm:"column:set_date" json:"set_date"`
	SetName         *string    `gorm:"column:set_name" json:"set_name"`
	UserID          *int       `gorm:"column:user_id" json:"user_id"`
	UuhBalanceID    *int       `gorm:"column:uuh_balance_id" json:"uuh_balance_id"`
	UuhFoodID       *int       `gorm:"column:uuh_food_id" json:"uuh_food_id"`
	UuhFoodName     *string    `gorm:"column:uuh_food_name" json:"uuh_food_name"`
	UuhFoodType     *string    `gorm:"column:uuh_food_type" json:"uuh_food_type"`
	UuhImage        *string    `gorm:"column:uuh_image" json:"uuh_image"`
	UuhKitchenID    *int       `gorm:"column:uuh_kitchen_id" json:"uuh_kitchen_id"`
	UuhQty          *float32   `gorm:"column:uuh_qty" json:"uuh_qty"`
	ZuushBalanceID  *int       `gorm:"column:zuush_balance_id" json:"zuush_balance_id"`
	ZuushFoodID     *int       `gorm:"column:zuush_food_id" json:"zuush_food_id"`
	ZuushFoodName   *string    `gorm:"column:zuush_food_name" json:"zuush_food_name"`
	ZuushFoodType   *string    `gorm:"column:zuush_food_type" json:"zuush_food_type"`
	ZuushImage      *string    `gorm:"column:zuush_image" json:"zuush_image"`
	ZuushKitchenID  *int       `gorm:"column:zuush_kitchen_id" json:"zuush_kitchen_id"`
	ZuushQty        *float32   `gorm:"column:zuush_qty" json:"zuush_qty"`
}

func (v *ViewCartSetFood) TableName() string {
	return "view_cart_set_food"
}

type ViewSetFoodSubMenu struct {
	FoodType   *string `gorm:"column:food_type" json:"food_type"`
	FoodTypeID *int    `gorm:"column:food_type_id" json:"food_type_id"`
	ID         *int    `gorm:"column:id" json:"id"`
	MenuID     *int    `gorm:"column:menu_id" json:"menu_id"`
}

func (v *ViewSetFoodSubMenu) TableName() string {
	return "view_set_food_sub_menu"
}

type ViewSetFoodSubMenuFoods struct {
	FoodID    *int `gorm:"column:food_id" json:"food_id"`
	ID        *int `gorm:"column:id" json:"id"`
	KitchenID *int `gorm:"column:kitchen_id" json:"kitchen_id"`
	SetQty    *int `gorm:"column:set_qty" json:"set_qty"`
	SubMenuID *int `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (v *ViewSetFoodSubMenuFoods) TableName() string {
	return "view_set_food_sub_menu_foods"
}
