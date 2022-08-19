package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisLegends278 struct {
	BorderColor *string `gorm:"column:border_color" json:"border_color"`
	ElementType *string `gorm:"column:element_type" json:"element_type"`
	FillColor   *string `gorm:"column:fill_color" json:"fill_color"`
	Icon        *string `gorm:"column:icon" json:"icon"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LayerID     *int    `gorm:"column:layer_id" json:"layer_id"`
	LineType    *string `gorm:"column:line_type" json:"line_type"`
	StyleType   *string `gorm:"column:style_type" json:"style_type"`
	StyleValue  *string `gorm:"column:style_value" json:"style_value"`
	Title       *string `gorm:"column:title" json:"title"`
}

//  TableName sets the insert table name for this struct type
func (g *GisLegends278) TableName() string {
	return "gis_legends"
}
