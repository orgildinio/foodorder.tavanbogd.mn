package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDamageType299 struct {
	DamageType *string `gorm:"column:damage_type" json:"damage_type"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDamageType299) TableName() string {
	return "lut_damage_type"
}
