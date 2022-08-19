package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubDamageType300 struct {
	DamageTypeID  *int    `gorm:"column:damage_type_id" json:"damage_type_id"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubDamageType *string `gorm:"column:sub_damage_type" json:"sub_damage_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubDamageType300) TableName() string {
	return "lut_sub_damage_type"
}
