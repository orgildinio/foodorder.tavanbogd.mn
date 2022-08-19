package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisBaseMaps272 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Image     *string `gorm:"column:image" json:"image"`
	IsDynamic *int    `gorm:"column:is_dynamic" json:"is_dynamic"`
	LayerName *string `gorm:"column:layerName" json:"layerName"`
	MaxZoom   *int    `gorm:"column:maxZoom" json:"maxZoom"`
	MinZoom   *int    `gorm:"column:minZoom" json:"minZoom"`
	Show      *int    `gorm:"column:show" json:"show"`
	URL       *string `gorm:"column:url" json:"url"`
}

//  TableName sets the insert table name for this struct type
func (g *GisBaseMaps272) TableName() string {
	return "gis_base_maps"
}
