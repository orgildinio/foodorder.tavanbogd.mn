package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHorogdsonAmitan176 struct {
	HorogdsonAmitan *string `gorm:"column:horogdson_amitan" json:"horogdson_amitan"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHorogdsonAmitan176) TableName() string {
	return "lut_horogdson_amitan"
}
