package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMaindisasterType84 struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MaindisasterType *string `gorm:"column:maindisaster_type" json:"maindisaster_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMaindisasterType84) TableName() string {
	return "lut_maindisaster_type"
}
