package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubSubMaindisasterType326 struct {
	ID                     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubMaindisasterTypeID  *int    `gorm:"column:sub_maindisaster_type_id" json:"sub_maindisaster_type_id"`
	SubSubMaindisasterType *string `gorm:"column:sub_sub_maindisaster_type" json:"sub_sub_maindisaster_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubSubMaindisasterType326) TableName() string {
	return "lut_sub_sub_maindisaster_type"
}
