package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHeregtseeTorol166 struct {
	HeregtseeTorol *string `gorm:"column:heregtsee_torol" json:"heregtsee_torol"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutHeregtseeTorol166) TableName() string {
	return "lut_heregtsee_torol"
}
