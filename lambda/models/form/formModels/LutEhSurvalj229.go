package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEhSurvalj229 struct {
	EhSurvalj *string `gorm:"column:eh_survalj" json:"eh_survalj"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutEhSurvalj229) TableName() string {
	return "lut_eh_survalj"
}
