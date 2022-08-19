package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type Users69 struct {
	AimagID        *int       `gorm:"column:aimag_id" json:"aimag_id"`
	AlbanTushaalID *int       `gorm:"column:alban_tushaal_id" json:"alban_tushaal_id"`
	AngiID         *int       `gorm:"column:angi_id" json:"angi_id"`
	BID            *int       `gorm:"column:b_id" json:"b_id"`
	BagID          *int       `gorm:"column:bag_id" json:"bag_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email          string     `gorm:"column:email" json:"email"`
	ID             int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Login          string     `gorm:"column:login" json:"login"`
	Password       string     `gorm:"column:password" json:"password"`
	RegisterNumber string     `gorm:"column:register_number" json:"register_number"`
	Role           *int       `gorm:"column:role" json:"role"`
	SalbariID      *int       `gorm:"column:salbari_id" json:"salbari_id"`
	SumID          *int       `gorm:"column:sum_id" json:"sum_id"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserTypeID     *int       `gorm:"column:user_type_id" json:"user_type_id"`
}

//  TableName sets the insert table name for this struct type
func (u *Users69) TableName() string {
	return "users"
}
