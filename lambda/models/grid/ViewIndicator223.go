package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewIndicator223 struct {
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           *int    `gorm:"column:id" json:"id"`
	Indicator    *string `gorm:"column:indicator" json:"indicator"`
}

type LutIndicatorMainTable223 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Indicator   *string `gorm:"column:indicator" json:"indicator"`
	IndicatorID *int    `gorm:"column:indicator_id" json:"indicator_id"`
}

func (l *LutIndicatorMainTable223) TableName() string {
	return "lut_indicator"
}

//  TableName sets the insert table name for this struct type
func (v *ViewIndicator223) TableName() string {
	return "view_indicator"
}

var ViewIndicator223Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Индикатор",
	Identity:  "id",
	DataTable: "view_indicator",
	MainTable: "lut_indicator",
	DataModel: new(ViewIndicator223),
	Data:      new([]ViewIndicator223),
	MainModel: new(LutIndicatorMainTable223),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "disaster_type", Label: "Гамшгийн эрсдлийн үзүүлэлтийн төрөл"},
		datagrid.Column{Model: "indicator", Label: "Индикатор"},
	},
	ColumnList: []string{"id", "disaster_type", "indicator"},
	Filters: map[string]string{
		"id": "",

		"disaster_type": "",

		"indicator": "",

		"indicator_id": "",
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
	FillVirtualColumns: fillVirtualColumnsViewIndicator223,
}

func fillVirtualColumnsViewIndicator223(rowsPre interface{}) interface{} {
	return rowsPre
}
