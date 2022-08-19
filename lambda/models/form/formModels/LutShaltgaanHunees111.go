package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutShaltgaanHunees111 struct {
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LutShaltgaanHunees *string `gorm:"column:lut_shaltgaan_hunees" json:"lut_shaltgaan_hunees"`
}

//  TableName sets the insert table name for this struct type
func (l *LutShaltgaanHunees111) TableName() string {
	return "lut_shaltgaan_hunees"
}
