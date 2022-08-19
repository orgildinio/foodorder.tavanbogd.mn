package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDedUndsenTorol308 struct {
	DedUndsenTorol *string `gorm:"column:ded_undsen_torol" json:"ded_undsen_torol"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndsenTorolID  *int    `gorm:"column:undsen_torol_id" json:"undsen_torol_id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDedUndsenTorol308) TableName() string {
	return "lut_ded_undsen_torol"
}
