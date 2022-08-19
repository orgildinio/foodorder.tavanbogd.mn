package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMethodolgyType40 struct {
	Argachlal *string `gorm:"column:argachlal" json:"argachlal"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMethodolgyType40) TableName() string {
	return "lut_methodolgy_type"
}
