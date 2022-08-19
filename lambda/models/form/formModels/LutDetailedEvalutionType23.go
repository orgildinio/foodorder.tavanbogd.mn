package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDetailedEvalutionType23 struct {
	DetailedEvalutionType *string `gorm:"column:detailed_evalution_type" json:"detailed_evalution_type"`
	ID                    int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDetailedEvalutionType23) TableName() string {
	return "lut_detailed_evalution_type"
}
