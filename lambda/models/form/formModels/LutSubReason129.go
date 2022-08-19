package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubReason129 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ReasonID  *int    `gorm:"column:reason_id" json:"reason_id"`
	SubReason *string `gorm:"column:sub_reason" json:"sub_reason"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubReason129) TableName() string {
	return "lut_sub_reason"
}
