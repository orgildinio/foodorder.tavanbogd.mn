package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type TblMenuGtNeg struct {
	CreatedAt   *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodTypeID  *int           `gorm:"column:food_type_id" json:"food_type_id"`
	GtBranch    *string        `gorm:"column:gt_branch" json:"gt_branch"`
	ID          int            `gorm:"column:id" json:"id"`
	IsOrder     *int           `gorm:"column:is_order" json:"is_order"`
	KitchenID   int            `gorm:"column:kitchen_id" json:"kitchen_id"`
	MainMenuID  int            `gorm:"column:main_menu_id" json:"main_menu_id"`
	OrderRuleID *int           `gorm:"column:order_rule_id" json:"order_rule_id"`
	SetDate     DB.Date        `gorm:"column:set_date" json:"set_date"`
	SetName     *string        `gorm:"column:set_name" json:"set_name"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID      *int           `gorm:"column:user_id" json:"user_id"`
}

func (t *TblMenuGtNeg) TableName() string {
	return "tbl_menu_gt_neg"
}

type SubMenuGtNeg struct {
	FoodTypeID *int `gorm:"column:food_type_id" json:"food_type_id"`
	ID         int  `gorm:"column:id" json:"id"`
	KitchenID  int  `gorm:"column:kitchen_id" json:"kitchen_id"`
	MenuID     int  `gorm:"column:menu_id" json:"menu_id"`
}

func (s *SubMenuGtNeg) TableName() string {
	return "sub_menu_gt_neg"
}

type SubMenuFoodsGtNeg struct {
	FoodID *int `gorm:"column:food_id" json:"food_id"`
	//FoodQuantity  *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	ID        int `gorm:"column:id" json:"id"`
	KitchenID int `gorm:"column:kitchen_id" json:"kitchen_id"`
	//OrderQuantity *float32 `gorm:"column:order_quantity" json:"order_quantity"`
	SubMenuID int `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (s *SubMenuFoodsGtNeg) TableName() string {
	return "sub_menu_foods_gt_neg"
}

type Kitchen struct {
	BranchID      *int    `gorm:"column:branch_id" json:"branch_id"`
	ID            int     `gorm:"column:id" json:"id"`
	KitcheAddress *string `gorm:"column:kitche_address" json:"kitche_address"`
	KitckenName   *string `gorm:"column:kitcken_name" json:"kitcken_name"`
	UserID        *int    `gorm:"column:user_id" json:"user_id"`
}

func (k *Kitchen) TableName() string {
	return "kitchen"
}
