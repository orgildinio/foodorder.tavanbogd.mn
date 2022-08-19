package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubMaindisasterType86 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LsubMaindisasterType *string `gorm:"column:lsub_maindisaster_type" json:"lsub_maindisaster_type"`
	MaindisasterTypeID   *int    `gorm:"column:maindisaster_type_id" json:"maindisaster_type_id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubMaindisasterType86) TableName() string {
	return "lut_sub_maindisaster_type"
}
