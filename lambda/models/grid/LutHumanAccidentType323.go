package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHumanAccidentType323 struct {
	AAccidentType *string `gorm:"column:a_accident_type" json:"a_accident_type"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutHumanAccidentTypeMainTable323 struct {
	AAccidentType *string `gorm:"column:a_accident_type" json:"a_accident_type"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHumanAccidentTypeMainTable323) TableName() string {
	return "lut_human_accident_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutHumanAccidentType323) TableName() string {
	return "lut_human_accident_type"
}

var LutHumanAccidentType323Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Гамшиг, аюулт үзэгдэл, осолд өртсөн байдал /H/",
	Identity:  "id",
	DataTable: "lut_human_accident_type",
	MainTable: "lut_human_accident_type",
	DataModel: new(LutHumanAccidentType323),
	Data:      new([]LutHumanAccidentType323),
	MainModel: new(LutHumanAccidentTypeMainTable323),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "a_accident_type", Label: "Гамшиг, аюулт үзэгдэл, осолд өртсөн байдал"},
	},
	ColumnList: []string{"id", "a_accident_type"},
	Filters: map[string]string{
		"id": "",

		"a_accident_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutHumanAccidentType323,
}

func fillVirtualColumnsLutHumanAccidentType323(rowsPre interface{}) interface{} {
	return rowsPre
}
