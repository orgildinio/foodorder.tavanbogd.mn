package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAlbanTushaal165 struct {
	AlbanTushaal *string `gorm:"column:alban_tushaal" json:"alban_tushaal"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAlbanTushaal165) TableName() string {
	return "lut_alban_tushaal"
}
