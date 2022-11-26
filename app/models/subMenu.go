package models

type SubMenuFoods struct {
    FoodID        *int `gorm:"column:food_id" json:"food_id"`
    ID            int  `gorm:"column:id" json:"id"`
    SubMenuFormID *int `gorm:"column:sub_menu_form_id" json:"sub_menu_form_id"`
}

func (s *SubMenuFoods) TableName() string {
    return "sub_menu_foods"
}

type SubMenu struct {
    FoodTypeID *int `gorm:"column:food_type_id" json:"food_type_id"`
    ID         int  `gorm:"column:id" json:"id"`
    MenuFormID *int `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenu) TableName() string {
    return "sub_menu"
}
