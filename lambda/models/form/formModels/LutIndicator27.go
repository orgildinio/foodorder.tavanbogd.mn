package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutIndicator27 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Indicator   *string `gorm:"column:indicator" json:"indicator"`
	IndicatorID *int    `gorm:"column:indicator_id" json:"indicator_id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutIndicator27) TableName() string {
	return "lut_indicator"
}
