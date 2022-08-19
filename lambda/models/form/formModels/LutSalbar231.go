package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSalbar231 struct {
	AngiID *int    `gorm:"column:angi_id" json:"angi_id"`
	ID     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Salbar *string `gorm:"column:salbar" json:"salbar"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSalbar231) TableName() string {
	return "lut_salbar"
}
