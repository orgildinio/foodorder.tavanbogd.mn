package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAvsanArgaHemjee175 struct {
	AvsanArgaHemjee *string `gorm:"column:avsan_arga_hemjee" json:"avsan_arga_hemjee"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAvsanArgaHemjee175) TableName() string {
	return "lut_avsan_arga_hemjee"
}
