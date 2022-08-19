package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutActivityType16 struct {
	ActivityType *string `gorm:"column:activity_type" json:"activity_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutActivityTypeMainTable16 struct {
	ActivityType *string `gorm:"column:activity_type" json:"activity_type"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutActivityTypeMainTable16) TableName() string {
	return "lut_activity_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutActivityType16) TableName() string {
	return "lut_activity_type"
}

var LutActivityType16Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Үйл ажиллагааны чиглэл ",
	Identity:  "id",
	DataTable: "lut_activity_type",
	MainTable: "lut_activity_type",
	DataModel: new(LutActivityType16),
	Data:      new([]LutActivityType16),
	MainModel: new(LutActivityTypeMainTable16),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "activity_type", Label: "Үйл ажиллагааны чиглэл "},
	},
	ColumnList: []string{"id", "activity_type"},
	Filters: map[string]string{
		"id": "",

		"activity_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutActivityType16,
}

func fillVirtualColumnsLutActivityType16(rowsPre interface{}) interface{} {
	return rowsPre
}
