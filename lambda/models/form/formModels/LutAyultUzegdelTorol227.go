package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAyultUzegdelTorol227 struct {
	AyultUzegdelTorol *string `gorm:"column:ayult_uzegdel_torol" json:"ayult_uzegdel_torol"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAyultUzegdelTorol227) TableName() string {
	return "lut_ayult_uzegdel_torol"
}
