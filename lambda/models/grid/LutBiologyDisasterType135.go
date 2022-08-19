package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutBiologyDisasterType135 struct {
	BiologyDisasterType *string `gorm:"column:biology_disaster_type" json:"biology_disaster_type"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutBiologyDisasterTypeMainTable135 struct {
	BiologyDisasterType *string `gorm:"column:biology_disaster_type" json:"biology_disaster_type"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutBiologyDisasterTypeMainTable135) TableName() string {
	return "lut_biology_disaster_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutBiologyDisasterType135) TableName() string {
	return "lut_biology_disaster_type"
}

var LutBiologyDisasterType135Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Биологийн гаралтай аюулт үзэгдлийн төрөл",
	Identity:  "id",
	DataTable: "lut_biology_disaster_type",
	MainTable: "lut_biology_disaster_type",
	DataModel: new(LutBiologyDisasterType135),
	Data:      new([]LutBiologyDisasterType135),
	MainModel: new(LutBiologyDisasterTypeMainTable135),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "biology_disaster_type", Label: "Биологийн гаралтай аюулт үзэгдлийн төрөл"},
	},
	ColumnList: []string{"id", "biology_disaster_type"},
	Filters: map[string]string{
		"id": "",

		"biology_disaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutBiologyDisasterType135,
}

func fillVirtualColumnsLutBiologyDisasterType135(rowsPre interface{}) interface{} {
	return rowsPre
}
