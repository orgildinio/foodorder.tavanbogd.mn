package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHunAmDedTorol253 struct {
	HunAmDedTorol *string `gorm:"column:hun_am_ded_torol" json:"hun_am_ded_torol"`
	HunAmTorolID  *int    `gorm:"column:hun_am_torol_id" json:"hun_am_torol_id"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHunAmDedTorol253) TableName() string {
	return "lut_hun_am_ded_torol"
}
