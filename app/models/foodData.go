package models

import (
	"gorm.io/gorm"
	"time"
)

type TblFood struct {
	CreatedAt       *time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FoodCalorie     *float32       `gorm:"column:food_calorie" json:"food_calorie"`
	FoodDetail      *string        `gorm:"column:food_detail" json:"food_detail"`
	FoodIamge       *string        `gorm:"column:food_iamge" json:"food_iamge"`
	FoodIngredients *string        `gorm:"column:food_ingredients" json:"food_ingredients"`
	FoodName        *string        `gorm:"column:food_name" json:"food_name"`
	FoodNumber      *float32       `gorm:"column:food_number" json:"food_number"`
	FoodPrice       *float32       `gorm:"column:food_price" json:"food_price"`
	FoodTypeID      int            `gorm:"column:food_type_id" json:"food_type_id"`
	ID              int            `gorm:"column:id" json:"id"`
	UpdatedAt       *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	UserID          *int           `gorm:"column:user_id" json:"user_id"`
}

func (t *TblFood) TableName() string {
	return "tbl_food"
}

type LutFoodType struct {
	FoodType *string `gorm:"column:food_type" json:"food_type"`
	ID       int     `gorm:"column:id" json:"id"`
}

func (l *LutFoodType) TableName() string {
	return "lut_food_type"
}

type FoodTypeRequest struct {
	Model struct {
		FoodTypeID   int `json:"food_type_id"`
		SubMenuFoods []struct {
			FoodID    interface{} `json:"food_id"`
			SubMenuID interface{} `json:"sub_menu_id"`
		} `json:"sub_menu_foods"`
	} `json:"model"`
	EditMode bool `json:"editMode"`
}

type Field struct {
	Field string                 `json:"field"`
	Value interface{}            `json:"value"`
	Props map[string]interface{} `json:"props"`
}
type Field2 struct {
	Field string                 `json:"field"`
	Value interface{}            `json:"value"`
	Props map[string]interface{} `json:"props"`
}
