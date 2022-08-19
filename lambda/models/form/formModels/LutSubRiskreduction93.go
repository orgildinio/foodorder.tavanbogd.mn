package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubRiskreduction93 struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RiskreductionID  *int    `gorm:"column:riskreduction_id" json:"riskreduction_id"`
	SubRiskreduction *string `gorm:"column:sub_riskreduction" json:"sub_riskreduction"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubRiskreduction93) TableName() string {
	return "lut_sub_riskreduction"
}
