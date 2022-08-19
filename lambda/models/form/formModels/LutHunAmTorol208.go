package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHunAmTorol208 struct {
	HunAmTorol *string `gorm:"column:hun_am_torol" json:"hun_am_torol"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHunAmTorol208) TableName() string {
	return "lut_hun_am_torol"
}
