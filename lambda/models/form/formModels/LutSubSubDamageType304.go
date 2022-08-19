package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubSubDamageType304 struct {
	DamageSubTypeID  *int    `gorm:"column:damage_sub_type_id" json:"damage_sub_type_id"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubsubDamageType *string `gorm:"column:subsub_damage_type" json:"subsub_damage_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubSubDamageType304) TableName() string {
	return "lut_sub_sub_damage_type"
}
