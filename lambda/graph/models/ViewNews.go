package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewNews struct {
	CreatedAt           *time.Time             `gorm:"column:created_at" json:"created_at"`
	DeletedAt           *time.Time             `gorm:"column:deleted_at" json:"deleted_at"`
	Delgerengui         *string                `gorm:"column:delgerengui" json:"delgerengui"`
	Garchig             *string                `gorm:"column:garchig" json:"garchig"`
	Handalt             *int                   `gorm:"column:handalt" json:"handalt"`
	ID                  *int                   `gorm:"column:id" json:"id"`
	Logo                *string                `gorm:"column:logo" json:"logo"`
	NewsType            *string                `gorm:"column:news_type" json:"news_type"`
	NewsTypeID          *int                   `gorm:"column:news_type_id" json:"news_type_id"`
	Share               *int                   `gorm:"column:share" json:"share"`
	UpdatedAt           *time.Time             `gorm:"column:updated_at" json:"updated_at"`
	UrsahEseh           *int                   `gorm:"column:ursah_eseh" json:"ursah_eseh"`
	Zurag               *string                `gorm:"column:zurag" json:"zurag"`
	SubNewsSocialTypies []*SubNewsSocialTypies `gorm:"-:all"`
}

//  TableName sets the insert table name for this struct type
func (v *ViewNews) TableName() string {
	return "view_news"
}
