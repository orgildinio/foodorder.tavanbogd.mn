package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAccDisasterType105 struct {
	Description  *string `gorm:"column:description" json:"description"`
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAccDisasterTypeMainTable105 struct {
	Description  *string `gorm:"column:description" json:"description"`
	DisasterType *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAccDisasterTypeMainTable105) TableName() string {
	return "lut_acc_disaster_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutAccDisasterType105) TableName() string {
	return "lut_acc_disaster_type"
}

var LutAccDisasterType105Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Ослын төрөл",
	Identity:  "id",
	DataTable: "lut_acc_disaster_type",
	MainTable: "lut_acc_disaster_type",
	DataModel: new(LutAccDisasterType105),
	Data:      new([]LutAccDisasterType105),
	MainModel: new(LutAccDisasterTypeMainTable105),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "disaster_type", Label: "Ослын төрөл"},
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
	FillVirtualColumns: fillVirtualColumnsLutAccDisasterType105,
}

func fillVirtualColumnsLutAccDisasterType105(rowsPre interface{}) interface{} {
	return rowsPre
}
