package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type HolbooBarih58 struct {
	BNer      *string    `gorm:"column:b_ner" json:"b_ner"`
	Bairshil  *string    `gorm:"column:bairshil" json:"bairshil"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email     *string    `gorm:"column:email" json:"email"`
	Hayag     *string    `gorm:"column:hayag" json:"hayag"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas      *string    `gorm:"column:utas" json:"utas"`
	Web       *string    `gorm:"column:web" json:"web"`
}

type SubSocialHolbooBarih58 struct {
	HolbooBarihID *int    `gorm:"column:holboo_barih_id" json:"holboo_barih_id"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Link          *string `gorm:"column:link" json:"link"`
	SocialTypeID  *int    `gorm:"column:social_type_id" json:"social_type_id"`
}

func (s *SubSocialHolbooBarih58) TableName() string {
	return "sub_social"
}

//  TableName sets the insert table name for this struct type
func (h *HolbooBarih58) TableName() string {
	return "holboo_barih"
}
