package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAvsanArgaDedHemjee311 struct {
	AvsanArgaDedHemjee *string `gorm:"column:avsan_arga_ded_hemjee" json:"avsan_arga_ded_hemjee"`
	AvsanArgaHemjeeID  *int    `gorm:"column:avsan_arga_hemjee_id" json:"avsan_arga_hemjee_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAvsanArgaDedHemjee311) TableName() string {
	return "lut_avsan_arga_ded_hemjee"
}
