package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutActivityType15 struct {
	ActivityType *string `gorm:"column:activity_type" json:"activity_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutActivityType15) TableName() string {
	return "lut_activity_type"
}
