package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisLayers276 struct {
	Active       *int       `gorm:"column:active" json:"active"`
	CategoryID   *int       `gorm:"column:category_id" json:"category_id"`
	CheckInluded *int       `gorm:"column:check_inluded" json:"check_inluded"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	Export       *int       `gorm:"column:export" json:"export"`
	ID           int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	InfoTemplate *string    `gorm:"column:info_template" json:"info_template"`
	LayerOrder   *int       `gorm:"column:layer_order" json:"layer_order"`
	LayerType    *string    `gorm:"column:layer_type" json:"layer_type"`
	LayerURL     *string    `gorm:"column:layer_url" json:"layer_url"`
	MenuOrder    *int       `gorm:"column:menu_order" json:"menu_order"`
	Name         *string    `gorm:"column:name" json:"name"`
	PopupFields  *string    `gorm:"column:popup_fields" json:"popup_fields"`
	SearchFields *string    `gorm:"column:search_fields" json:"search_fields"`
	SearchInfo   *string    `gorm:"column:search_info" json:"search_info"`
	Show         *int       `gorm:"column:show" json:"show"`
	StyleField   *string    `gorm:"column:style_field" json:"style_field"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (g *GisLayers276) TableName() string {
	return "gis_layers"
}
