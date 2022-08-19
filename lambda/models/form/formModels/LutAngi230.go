package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAngi230 struct {
	Angi *string `gorm:"column:angi" json:"angi"`
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAngi230) TableName() string {
	return "lut_angi"
}
