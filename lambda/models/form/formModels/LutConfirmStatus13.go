package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutConfirmStatus13 struct {
	ConfirmStatus *string `gorm:"column:confirm_status" json:"confirm_status"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutConfirmStatus13) TableName() string {
	return "lut_confirm_status"
}
