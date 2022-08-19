package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisLayers277 struct {
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      *string    `gorm:"column:name" json:"name"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type GisLayersMainTable277 struct {
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

func (g *GisLayersMainTable277) TableName() string {
	return "gis_layers"
}

//  TableName sets the insert table name for this struct type
func (g *GisLayers277) TableName() string {
	return "gis_layers"
}

var GisLayers277Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Давхарга",
	Identity:  "id",
	DataTable: "gis_layers",
	MainTable: "gis_layers",
	DataModel: new(GisLayers277),
	Data:      new([]GisLayers277),
	MainModel: new(GisLayersMainTable277),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "name", Label: "Нэр"},
	},
	ColumnList: []string{"id", "name"},
	Filters: map[string]string{
		"id": "",

		"category_id": "Select",

		"active": "",

		"show": "",

		"name": "",

		"menu_order": "",

		"layer_order": "",

		"layer_type": "",

		"layer_url": "",

		"info_template": "",

		"search_info": "",

		"created_at": "",

		"updated_at": "",

		"popup_fields": "",

		"search_fields": "",

		"style_field": "",

		"check_inluded": "",

		"export": "",
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
	FillVirtualColumns: fillVirtualColumnsGisLayers277,
}

func fillVirtualColumnsGisLayers277(rowsPre interface{}) interface{} {
	return rowsPre
}
