package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEvalutionSector11 struct {
	EvalutionSector *string `gorm:"column:evalution_sector" json:"evalution_sector"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutEvalutionSector11) TableName() string {
	return "lut_evalution_sector"
}
