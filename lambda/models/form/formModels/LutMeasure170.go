package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMeasure170 struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Measure *string `gorm:"column:measure" json:"measure"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMeasure170) TableName() string {
	return "lut_measure"
}
