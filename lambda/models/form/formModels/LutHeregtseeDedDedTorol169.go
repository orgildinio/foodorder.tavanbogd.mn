package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHeregtseeDedDedTorol169 struct {
	HeregtseeDedDedTorol *string `gorm:"column:heregtsee_ded__ded_torol" json:"heregtsee_ded__ded_torol"`
	HeregtseeDedTorolID  *int    `gorm:"column:heregtsee_ded_torol_id" json:"heregtsee_ded_torol_id"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHeregtseeDedDedTorol169) TableName() string {
	return "lut_heregtsee_ded_ded_torol"
}
