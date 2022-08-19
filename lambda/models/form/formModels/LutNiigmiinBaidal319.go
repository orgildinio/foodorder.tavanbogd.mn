package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNiigmiinBaidal319 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NiigmiinBaidal *string `gorm:"column:niigmiin_baidal" json:"niigmiin_baidal"`
}

//  TableName sets the insert table name for this struct type
func (l *LutNiigmiinBaidal319) TableName() string {
	return "lut_niigmiin_baidal"
}
