package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type Friends1 struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Phone string `gorm:"column:phone" json:"phone"`
	Zurag string `gorm:"column:zurag" json:"zurag"`
}

//  TableName sets the insert table name for this struct type
func (f *Friends1) TableName() string {
	return "friends"
}
