package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDailyDisasterType20 struct {
	DailyDisasterType *string `gorm:"column:daily_disaster_type" json:"daily_disaster_type"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutDailyDisasterTypeMainTable20 struct {
	DailyDisasterType *string `gorm:"column:daily_disaster_type" json:"daily_disaster_type"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDailyDisasterTypeMainTable20) TableName() string {
	return "lut_daily_disaster_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutDailyDisasterType20) TableName() string {
	return "lut_daily_disaster_type"
}

var LutDailyDisasterType20Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Гамшгийн төрөл",
	Identity:  "id",
	DataTable: "lut_daily_disaster_type",
	MainTable: "lut_daily_disaster_type",
	DataModel: new(LutDailyDisasterType20),
	Data:      new([]LutDailyDisasterType20),
	MainModel: new(LutDailyDisasterTypeMainTable20),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "daily_disaster_type", Label: "Гамшгийн төрөл"},
	},
	ColumnList: []string{"id", "daily_disaster_type"},
	Filters: map[string]string{
		"id": "",

		"daily_disaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutDailyDisasterType20,
}

func fillVirtualColumnsLutDailyDisasterType20(rowsPre interface{}) interface{} {
	return rowsPre
}
