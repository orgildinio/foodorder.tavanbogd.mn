package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubIndicator29 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IndicatorID  *int    `gorm:"column:indicator_id" json:"indicator_id"`
	SubIndicator *string `gorm:"column:sub_indicator" json:"sub_indicator"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubIndicator29) TableName() string {
	return "lut_sub_indicator"
}
