package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNewsType struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsType *string `gorm:"column:news_type" json:"news_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutNewsType) TableName() string {
	return "lut_news_type"
}
