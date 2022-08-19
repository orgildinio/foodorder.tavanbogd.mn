package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubReasonHunees113 struct {
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanHuneesID *int    `gorm:"column:shaltgaan_hunees_id" json:"shaltgaan_hunees_id"`
	SubReasonHunees   *string `gorm:"column:sub_reason_hunees" json:"sub_reason_hunees"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubReasonHunees113) TableName() string {
	return "lut_sub_reason_hunees"
}
