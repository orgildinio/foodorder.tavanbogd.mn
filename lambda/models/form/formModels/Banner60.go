package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type Banner60 struct {
	Banner       *string    `gorm:"column:banner" json:"banner"`
	BannerTypeID *int       `gorm:"column:banner_type_id" json:"banner_type_id"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID           int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsID       *int       `gorm:"column:news_id" json:"news_id"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (b *Banner60) TableName() string {
	return "banner"
}
