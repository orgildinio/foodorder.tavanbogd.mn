package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMeasure186 struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Measure *string `gorm:"column:measure" json:"measure"`
}

type LutMeasureMainTable186 struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Measure *string `gorm:"column:measure" json:"measure"`
}

func (l *LutMeasureMainTable186) TableName() string {
	return "lut_measure"
}

//  TableName sets the insert table name for this struct type
func (l *LutMeasure186) TableName() string {
	return "lut_measure"
}

var LutMeasure186Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хэмжих нэгж",
	Identity:  "id",
	DataTable: "lut_measure",
	MainTable: "lut_measure",
	DataModel: new(LutMeasure186),
	Data:      new([]LutMeasure186),
	MainModel: new(LutMeasureMainTable186),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "measure", Label: "Хэмжих нэгж"},
	},
	ColumnList: []string{"id", "measure"},
	Filters: map[string]string{
		"id": "",

		"measure": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMeasure186,
}

func fillVirtualColumnsLutMeasure186(rowsPre interface{}) interface{} {
	return rowsPre
}
