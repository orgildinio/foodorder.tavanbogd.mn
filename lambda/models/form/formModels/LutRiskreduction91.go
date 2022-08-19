package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutRiskreduction91 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Riskreduction *string `gorm:"column:riskreduction" json:"riskreduction"`
}

//  TableName sets the insert table name for this struct type
func (l *LutRiskreduction91) TableName() string {
	return "lut_riskreduction"
}
