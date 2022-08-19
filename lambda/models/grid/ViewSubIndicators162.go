package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubIndicators162 struct {
	ID           *int    `gorm:"column:id" json:"id"`
	Indicator    *string `gorm:"column:indicator" json:"indicator"`
	SubIndicator *string `gorm:"column:sub_indicator" json:"sub_indicator"`
}

type LutSubIndicatorMainTable162 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IndicatorID  *int    `gorm:"column:indicator_id" json:"indicator_id"`
	SubIndicator *string `gorm:"column:sub_indicator" json:"sub_indicator"`
}

func (l *LutSubIndicatorMainTable162) TableName() string {
	return "lut_sub_indicator"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubIndicators162) TableName() string {
	return "view_sub_indicators"
}

var ViewSubIndicators162Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Дэд индикатор",
	Identity:  "id",
	DataTable: "view_sub_indicators",
	MainTable: "lut_sub_indicator",
	DataModel: new(ViewSubIndicators162),
	Data:      new([]ViewSubIndicators162),
	MainModel: new(LutSubIndicatorMainTable162),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "indicator", Label: "Индикатор"},
		datagrid.Column{Model: "sub_indicator", Label: "Дэд индикатор"},
	},
	ColumnList: []string{"id", "indicator", "sub_indicator"},
	Filters: map[string]string{
		"id": "",

		"indicator_id": "",

		"indicator": "",

		"sub_indicator": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubIndicators162,
}

func fillVirtualColumnsViewSubIndicators162(rowsPre interface{}) interface{} {
	return rowsPre
}
