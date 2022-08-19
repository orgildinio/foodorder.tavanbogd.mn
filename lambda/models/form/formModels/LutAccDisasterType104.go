package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAccDisasterType104 struct {
	Description  *string `gorm:"column:description" json:"description"`
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAccDisasterType104) TableName() string {
	return "lut_acc_disaster_type"
}
