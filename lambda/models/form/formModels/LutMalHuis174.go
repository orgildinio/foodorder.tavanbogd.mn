package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalHuis174 struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalHuis *string `gorm:"column:mal_huis" json:"mal_huis"`
}

//  TableName sets the insert table name for this struct type
func (l *LutMalHuis174) TableName() string {
	return "lut_mal_huis"
}
