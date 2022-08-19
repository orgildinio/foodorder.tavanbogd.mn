package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDisasterType18 struct {
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutDisasterTypeMainTable18 struct {
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDisasterTypeMainTable18) TableName() string {
	return "lut_disaster_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutDisasterType18) TableName() string {
	return "lut_disaster_type"
}

var LutDisasterType18Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Гамшгийн төрөл",
	Identity:  "id",
	DataTable: "lut_disaster_type",
	MainTable: "lut_disaster_type",
	DataModel: new(LutDisasterType18),
	Data:      new([]LutDisasterType18),
	MainModel: new(LutDisasterTypeMainTable18),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "disaster_type", Label: "Гамшгийн төрөл"},
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
	FillVirtualColumns: fillVirtualColumnsLutDisasterType18,
}

func fillVirtualColumnsLutDisasterType18(rowsPre interface{}) interface{} {
	return rowsPre
}
