package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHeregtseeDedTorol168 struct {
	HeregtseeDedTorol *string `gorm:"column:heregtsee_ded_torol" json:"heregtsee_ded_torol"`
	HeregtseeTorolID  *int    `gorm:"column:heregtsee_torol_id" json:"heregtsee_torol_id"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHeregtseeDedTorol168) TableName() string {
	return "lut_heregtsee_ded_torol"
}
