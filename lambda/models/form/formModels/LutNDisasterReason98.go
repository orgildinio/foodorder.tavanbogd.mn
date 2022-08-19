package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNDisasterReason98 struct {
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NDisasterReason *string `gorm:"column:n_disaster_reason" json:"n_disaster_reason"`
}

//  TableName sets the insert table name for this struct type
func (l *LutNDisasterReason98) TableName() string {
	return "lut_n_disaster_reason"
}
