package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubAccDisasterType108 struct {
	AccDisasterTypeID  *int    `gorm:"column:acc_disaster_type_id" json:"acc_disaster_type_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubAccDisasterType *string `gorm:"column:sub_acc_disaster_type" json:"sub_acc_disaster_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubAccDisasterType108) TableName() string {
	return "lut_sub_acc_disaster_type"
}
