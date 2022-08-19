package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubSubNervgdsenHun144 struct {
	DedDedTorol *string `gorm:"column:ded_ded_torol" json:"ded_ded_torol"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NDedTorolID *int    `gorm:"column:n_ded_torol_id" json:"n_ded_torol_id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubSubNervgdsenHun144) TableName() string {
	return "lut_sub_sub_nervgdsen_hun"
}
