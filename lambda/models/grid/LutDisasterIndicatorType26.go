package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDisasterIndicatorType26 struct {
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutDisasterIndicatorTypeMainTable26 struct {
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDisasterIndicatorTypeMainTable26) TableName() string {
	return "lut_disaster_indicator_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutDisasterIndicatorType26) TableName() string {
	return "lut_disaster_indicator_type"
}

var LutDisasterIndicatorType26Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Гамшгийн эрсдлийн үзүүлэлтийн төрөл ",
	Identity:  "id",
	DataTable: "lut_disaster_indicator_type",
	MainTable: "lut_disaster_indicator_type",
	DataModel: new(LutDisasterIndicatorType26),
	Data:      new([]LutDisasterIndicatorType26),
	MainModel: new(LutDisasterIndicatorTypeMainTable26),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "disaster_type", Label: "Гамшгийн эрсдлийн үзүүлэлтийн төрөл "},
	},
	ColumnList: []string{"id", "disaster_type"},
	Filters: map[string]string{
		"id": "",

		"disaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutDisasterIndicatorType26,
}

func fillVirtualColumnsLutDisasterIndicatorType26(rowsPre interface{}) interface{} {
	return rowsPre
}
