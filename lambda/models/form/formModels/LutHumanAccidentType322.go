package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHumanAccidentType322 struct {
	AAccidentType *string `gorm:"column:a_accident_type" json:"a_accident_type"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHumanAccidentType322) TableName() string {
	return "lut_human_accident_type"
}
