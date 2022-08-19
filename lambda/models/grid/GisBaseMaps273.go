package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisBaseMaps273 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LayerName *string `gorm:"column:layerName" json:"layerName"`
}

type GisBaseMapsMainTable273 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Image     *string `gorm:"column:image" json:"image"`
	IsDynamic *int    `gorm:"column:is_dynamic" json:"is_dynamic"`
	LayerName *string `gorm:"column:layerName" json:"layerName"`
	MaxZoom   *int    `gorm:"column:maxZoom" json:"maxZoom"`
	MinZoom   *int    `gorm:"column:minZoom" json:"minZoom"`
	Show      *int    `gorm:"column:show" json:"show"`
	URL       *string `gorm:"column:url" json:"url"`
}

func (g *GisBaseMapsMainTable273) TableName() string {
	return "gis_base_maps"
}

//  TableName sets the insert table name for this struct type
func (g *GisBaseMaps273) TableName() string {
	return "gis_base_maps"
}

var GisBaseMaps273Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Суурь зураг",
	Identity:  "id",
	DataTable: "gis_base_maps",
	MainTable: "gis_base_maps",
	DataModel: new(GisBaseMaps273),
	Data:      new([]GisBaseMaps273),
	MainModel: new(GisBaseMapsMainTable273),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "layerName", Label: "Нэр"},
	},
	ColumnList: []string{"id", "layerName"},
	Filters: map[string]string{
		"id": "",

		"layerName": "",

		"url": "",

		"minZoom": "",

		"maxZoom": "",

		"show": "",

		"is_dynamic": "",

		"image": "",
	},
	Relations:   []models.GridRelation{},
	Condition:   "",
	Aggergation: "",
	BeforeFetch: nil,

	AfterFetch: nil,

	BeforeDelete: nil,

	AfterDelete: nil,

	BeforePrint:        nil,
	TriggerNameSpace:   "",
	FillVirtualColumns: fillVirtualColumnsGisBaseMaps273,
}

func fillVirtualColumnsGisBaseMaps273(rowsPre interface{}) interface{} {
	return rowsPre
}
