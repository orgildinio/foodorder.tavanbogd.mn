package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalTorol171 struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalTorol *string `gorm:"column:mal_torol" json:"mal_torol"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMalTorol171) TableName() string {
	return "lut_mal_torol"
}
