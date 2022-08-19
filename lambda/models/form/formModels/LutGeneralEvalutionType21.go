package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutGeneralEvalutionType21 struct {
	GeneralEvalutionType *string `gorm:"column:general_evalution_type" json:"general_evalution_type"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutGeneralEvalutionType21) TableName() string {
	return "lut_general_evalution_type"
}
