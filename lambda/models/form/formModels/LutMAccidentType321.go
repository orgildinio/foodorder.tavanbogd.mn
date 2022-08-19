package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMAccidentType321 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MAccidentType *string `gorm:"column:m_accident_type" json:"m_accident_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMAccidentType321) TableName() string {
	return "lut_m_accident_type"
}
