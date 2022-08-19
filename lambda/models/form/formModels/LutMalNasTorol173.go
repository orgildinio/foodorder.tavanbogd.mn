package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalNasTorol173 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalNasTorol *string `gorm:"column:mal_nas_torol" json:"mal_nas_torol"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMalNasTorol173) TableName() string {
	return "lut_mal_nas_torol"
}
