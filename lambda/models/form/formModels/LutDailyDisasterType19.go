package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDailyDisasterType19 struct {
	DailyDisasterType *string `gorm:"column:daily_disaster_type" json:"daily_disaster_type"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDailyDisasterType19) TableName() string {
	return "lut_daily_disaster_type"
}
