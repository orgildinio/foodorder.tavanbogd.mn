package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHohiroliinBaidal180 struct {
	HohiroliinBaidal *string `gorm:"column:hohiroliin_baidal" json:"hohiroliin_baidal"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHohiroliinBaidal180) TableName() string {
	return "lut_hohiroliin_baidal"
}
