package models

import "time"

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
	MenuID     *int `gorm:"column:menu_id" json:"menu_id"`
}

func (s *SubMenu) TableName() string {
	return "sub_menu"
}

type SubMenuFoods struct {
	FoodID    *int `gorm:"column:food_id" json:"food_id"`
	ID        int  `gorm:"column:id" json:"id"`
	SubMenuID *int `gorm:"column:sub_menu_id" json:"sub_menu_id"`
}

func (s *SubMenuFoods) TableName() string {
	return "sub_menu_foods"
}
