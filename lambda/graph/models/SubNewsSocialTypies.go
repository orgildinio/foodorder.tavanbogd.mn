package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubNewsSocialTypies struct {
	ID           int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsID       *int `gorm:"column:news_id" json:"news_id"`
	SocialTypeID *int `gorm:"column:social_type_id" json:"social_type_id"`
}

//  TableName sets the insert table name for this struct type
func (s *SubNewsSocialTypies) TableName() string {
	return "sub_news_social_typies"
}
