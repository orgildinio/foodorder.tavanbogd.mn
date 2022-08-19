package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubAnimals124 struct {
	AnimalsID  *int    `gorm:"column:animals_id" json:"animals_id"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubAnimals *string `gorm:"column:sub_animals" json:"sub_animals"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubAnimals124) TableName() string {
	return "lut_sub_animals"
}
