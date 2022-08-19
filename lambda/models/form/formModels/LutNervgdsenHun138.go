package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNervgdsenHun138 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NervgdsenHun *string `gorm:"column:nervgdsen_hun" json:"nervgdsen_hun"`
}

//  TableName sets the insert table name for this struct type
func (l *LutNervgdsenHun138) TableName() string {
	return "lut_nervgdsen_hun"
}
