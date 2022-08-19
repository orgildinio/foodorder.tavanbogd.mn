package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutUndsenTorol306 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndsenTorol *string `gorm:"column:undsen_torol" json:"undsen_torol"`
}

//  TableName sets the insert table name for this struct type
func (l *LutUndsenTorol306) TableName() string {
	return "lut_undsen_torol"
}
