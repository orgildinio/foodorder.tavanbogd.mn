package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutReason128 struct {
	ID     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Reason *string `gorm:"column:reason" json:"reason"`
}

//  TableName sets the insert table name for this struct type
func (l *LutReason128) TableName() string {
	return "lut_reason"
}
