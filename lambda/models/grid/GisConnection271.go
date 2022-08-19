package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisConnection271 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LocalGrid *int    `gorm:"column:local_grid" json:"local_grid"`
	Title     *string `gorm:"column:title" json:"title"`
}

type GisConnectionMainTable271 struct {
	Connection *string `gorm:"column:connection" json:"connection"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Layer      *string `gorm:"column:layer" json:"layer"`
	LocalForm  *int    `gorm:"column:local_form" json:"local_form"`
	LocalGrid  *int    `gorm:"column:local_grid" json:"local_grid"`
	Title      *string `gorm:"column:title" json:"title"`
}

func (g *GisConnectionMainTable271) TableName() string {
	return "gis_connection"
}

//  TableName sets the insert table name for this struct type
func (g *GisConnection271) TableName() string {
	return "gis_connection"
}

var GisConnection271Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Газрын зургийн сервис холболт",
	Identity:  "id",
	DataTable: "gis_connection",
	MainTable: "gis_connection",
	DataModel: new(GisConnection271),
	Data:      new([]GisConnection271),
	MainModel: new(GisConnectionMainTable271),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "title", Label: "Нэр"},
		datagrid.Column{Model: "local_grid", Label: ""},
	},
	ColumnList: []string{"id", "title", "local_grid"},
	Filters: map[string]string{
		"id": "",

		"title": "",

		"connection": "",

		"layer": "",

		"local_form": "",

		"local_grid": "",
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
	FillVirtualColumns: fillVirtualColumnsGisConnection271,
}

func fillVirtualColumnsGisConnection271(rowsPre interface{}) interface{} {
	return rowsPre
}
