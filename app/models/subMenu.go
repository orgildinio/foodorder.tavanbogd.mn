package models

type SubMenuNeg struct {
	FoodNegID  *int `gorm:"column:food_neg_id" json:"food_neg_id"`
	ID         int  `gorm:"column:id" json:"id"`
	MenuFormID *int `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuNeg) TableName() string {
	return "sub_menu_neg"
}

type SubMenuHoer struct {
	FoodHoerID *int `gorm:"column:food_hoer_id" json:"food_hoer_id"`
	ID         int  `gorm:"column:id" json:"id"`
	MenuFormID *int `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuHoer) TableName() string {
	return "sub_menu_hoer"
}

type SubMenuUuhim struct {
	FoodUuhimID *int `gorm:"column:food_uuhim_id" json:"food_uuhim_id"`
	ID          int  `gorm:"column:id" json:"id"`
	MenuFormID  *int `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuUuhim) TableName() string {
	return "sub_menu_uuhim"
}

type SubMenuDesert struct {
	FoodDesertID *int `gorm:"column:food_desert_id" json:"food_desert_id"`
	ID           int  `gorm:"column:id" json:"id"`
	MenuFormID   *int `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuDesert) TableName() string {
	return "sub_menu_desert"
}

type SubMenuSalat struct {
	FoodSalatID *int `gorm:"column:food_salat_id" json:"food_salat_id"`
	ID          int  `gorm:"column:id" json:"id"`
	MenuFormID  *int `gorm:"column:menu_form_id" json:"menu_form_id"`
}

func (s *SubMenuSalat) TableName() string {
	return "sub_menu_salat"
}
