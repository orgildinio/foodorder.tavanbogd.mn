package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalDedTorol172 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalDedTorol *string `gorm:"column:mal_ded_torol" json:"mal_ded_torol"`
	MalTorolID  *int    `gorm:"column:mal_torol_id" json:"mal_torol_id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMalDedTorol172) TableName() string {
	return "lut_mal_ded_torol"
}
