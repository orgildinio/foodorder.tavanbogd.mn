package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutInfoSource136 struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	InfoSource *string `gorm:"column:info_source" json:"info_source"`
}

//  TableName sets the insert table name for this struct type
func (l *LutInfoSource136) TableName() string {
	return "lut_info_source"
}
