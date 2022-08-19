package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutBiologyDisasterType134 struct {
	BiologyDisasterType *string `gorm:"column:biology_disaster_type" json:"biology_disaster_type"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutBiologyDisasterType134) TableName() string {
	return "lut_biology_disaster_type"
}
