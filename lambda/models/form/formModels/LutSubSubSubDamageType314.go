package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubSubSubDamageType314 struct {
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubSubDamageTypeID  *int    `gorm:"column:sub_sub_damage_type_id" json:"sub_sub_damage_type_id"`
	SubSubSubDamageType *string `gorm:"column:sub_sub_sub_damage_type" json:"sub_sub_sub_damage_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubSubSubDamageType314) TableName() string {
	return "lut_sub_sub_sub_damage_type"
}
