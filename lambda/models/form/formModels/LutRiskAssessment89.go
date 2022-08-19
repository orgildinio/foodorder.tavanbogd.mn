package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutRiskAssessment89 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RiskAssessment *string `gorm:"column:risk_assessment" json:"risk_assessment"`
}

//  TableName sets the insert table name for this struct type
func (l *LutRiskAssessment89) TableName() string {
	return "lut_risk_assessment"
}
