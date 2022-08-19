package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAyultUzegdelDedTorol228 struct {
	AyultUzegdelDedTorol *string `gorm:"column:ayult_uzegdel_ded_torol" json:"ayult_uzegdel_ded_torol"`
	AyultUzegdelTorolID  *int    `gorm:"column:ayult_uzegdel_torol_id" json:"ayult_uzegdel_torol_id"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAyultUzegdelDedTorol228) TableName() string {
	return "lut_ayult_uzegdel_ded_torol"
}
