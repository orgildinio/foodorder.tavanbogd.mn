package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutTsol232 struct {
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Tsol *string `gorm:"column:tsol" json:"tsol"`
}

//  TableName sets the insert table name for this struct type
func (l *LutTsol232) TableName() string {
	return "lut_tsol"
}
