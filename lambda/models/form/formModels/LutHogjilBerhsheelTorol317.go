package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHogjilBerhsheelTorol317 struct {
	HogjilBerhsheelTorol *string `gorm:"column:hogjil_berhsheel_torol" json:"hogjil_berhsheel_torol"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHogjilBerhsheelTorol317) TableName() string {
	return "lut_hogjil_berhsheel_torol"
}
