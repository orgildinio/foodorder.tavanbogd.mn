package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMAccidentType324 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MAccidentType *string `gorm:"column:m_accident_type" json:"m_accident_type"`
}

type LutMAccidentTypeMainTable324 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MAccidentType *string `gorm:"column:m_accident_type" json:"m_accident_type"`
}

func (l *LutMAccidentTypeMainTable324) TableName() string {
	return "lut_m_accident_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutMAccidentType324) TableName() string {
	return "lut_m_accident_type"
}

var LutMAccidentType324Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Гамшиг, аюулт үзэгдэл, осолд өртсөн байдал /M/",
	Identity:  "id",
	DataTable: "lut_m_accident_type",
	MainTable: "lut_m_accident_type",
	DataModel: new(LutMAccidentType324),
	Data:      new([]LutMAccidentType324),
	MainModel: new(LutMAccidentTypeMainTable324),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "m_accident_type", Label: "Гамшиг, аюулт үзэгдэл, осолд өртсөн байдал"},
	},
	ColumnList: []string{"id", "m_accident_type"},
	Filters: map[string]string{
		"id": "",

		"m_accident_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMAccidentType324,
}

func fillVirtualColumnsLutMAccidentType324(rowsPre interface{}) interface{} {
	return rowsPre
}
