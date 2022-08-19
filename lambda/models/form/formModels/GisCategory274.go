package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisCategory274 struct {
	Active     *int       `gorm:"column:active" json:"active"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	Icon       *string    `gorm:"column:icon" json:"icon"`
	ID         int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LayerOrder *int       `gorm:"column:layer_order" json:"layer_order"`
	MenuOrder  *int       `gorm:"column:menu_order" json:"menu_order"`
	Name       *string    `gorm:"column:name" json:"name"`
	Show       *int       `gorm:"column:show" json:"show"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (g *GisCategory274) TableName() string {
	return "gis_category"
}
