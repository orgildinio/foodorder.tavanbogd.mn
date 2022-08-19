package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutFire106 struct {
	Fire *string `gorm:"column:fire" json:"fire"`
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutFire106) TableName() string {
	return "lut_fire"
}
