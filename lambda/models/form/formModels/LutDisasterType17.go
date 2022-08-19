package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDisasterType17 struct {
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDisasterType17) TableName() string {
	return "lut_disaster_type"
}
