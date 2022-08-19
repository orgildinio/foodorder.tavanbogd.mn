package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAnimals123 struct {
	Animals *string `gorm:"column:animals" json:"animals"`
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAnimals123) TableName() string {
	return "lut_animals"
}
