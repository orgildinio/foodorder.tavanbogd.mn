package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisConnection270 struct {
	Connection *string `gorm:"column:connection" json:"connection"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Layer      *string `gorm:"column:layer" json:"layer"`
	LocalForm  *int    `gorm:"column:local_form" json:"local_form"`
	LocalGrid  *int    `gorm:"column:local_grid" json:"local_grid"`
	Title      *string `gorm:"column:title" json:"title"`
}

//  TableName sets the insert table name for this struct type
func (g *GisConnection270) TableName() string {
	return "gis_connection"
}
