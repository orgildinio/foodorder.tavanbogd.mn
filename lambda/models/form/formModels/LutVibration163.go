package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutVibration163 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Vibration *string `gorm:"column:vibration" json:"vibration"`
}

//  TableName sets the insert table name for this struct type
func (l *LutVibration163) TableName() string {
	return "lut_vibration"
}
