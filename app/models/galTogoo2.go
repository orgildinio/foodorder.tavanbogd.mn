package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

type TblMenuGtHoer struct {
	BranchChefID   *int           `gorm:"column:branch_chef_id" json:"branch_chef_id"`
	CreatedAt      *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodDesertID   *int           `gorm:"column:food_desert_id" json:"food_desert_id"`
	FoodHoerID     *int           `gorm:"column:food_hoer_id" json:"food_hoer_id"`
	FoodSalatID    *int           `gorm:"column:food_salat_id" json:"food_salat_id"`
	FoodTimeTypeID *int           `gorm:"column:food_time_type_id" json:"food_time_type_id"`
	FoodUuhimID    *int           `gorm:"column:food_uuhim_id" json:"food_uuhim_id"`
	ID             int            `gorm:"column:id" json:"id"`
	MainMenuID     *int           `gorm:"column:main_menu_id" json:"main_menu_id"`
	SetDate        *DB.Date       `gorm:"column:set_date" json:"set_date"`
	SetName        *string        `gorm:"column:set_name" json:"set_name"`
	UpdatedAt      *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID         *int           `gorm:"column:user_id" json:"user_id"`
}

func (t *TblMenuGtHoer) TableName() string {
	return "tbl_menu_gt_hoer"
}

type SubMenuDesertGtHoer struct {
	FoodDesertID *int     `gorm:"column:food_desert_id" json:"food_desert_id"`
	FoodQuantity *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	ID           int      `gorm:"column:id" json:"id"`
	MenuFormID   *int     `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuDesertGtHoer) TableName() string {
	return "sub_menu_desert_gt_hoer"
}

type SubMenuHoerGtHoer struct {
	FoodHoerID   *int     `gorm:"column:food_hoer_id" json:"food_hoer_id"`
	FoodQuantity *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	ID           int      `gorm:"column:id" json:"id"`
	MenuFormID   *int     `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuHoerGtHoer) TableName() string {
	return "sub_menu_hoer_gt_hoer"
}

type SubMenuNegGtHoer struct {
	FoodNegID    *int     `gorm:"column:food_neg_id" json:"food_neg_id"`
	FoodQuantity *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	ID           int      `gorm:"column:id" json:"id"`
	MenuFormID   *int     `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuNegGtHoer) TableName() string {
	return "sub_menu_neg_gt_hoer"
}

type SubMenuSalatGtHoer struct {
	FoodQuantity *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	FoodSalatID  *int     `gorm:"column:food_salat_id" json:"food_salat_id"`
	ID           int      `gorm:"column:id" json:"id"`
	MenuFormID   *int     `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuSalatGtHoer) TableName() string {
	return "sub_menu_salat_gt_hoer"
}

type SubMenuUuhimGtHoer struct {
	FoodQuantity *float32 `gorm:"column:food_quantity" json:"food_quantity"`
	FoodUuhimID  *int     `gorm:"column:food_uuhim_id" json:"food_uuhim_id"`
	ID           int      `gorm:"column:id" json:"id"`
	MenuFormID   *int     `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuUuhimGtHoer) TableName() string {
	return "sub_menu_uuhim_gt_hoer"
}
