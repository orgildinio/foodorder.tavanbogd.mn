package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAngi234 struct {
	Angi *string `gorm:"column:angi" json:"angi"`
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAngiMainTable234 struct {
	Angi *string `gorm:"column:angi" json:"angi"`
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAngiMainTable234) TableName() string {
	return "lut_angi"
}

//  TableName sets the insert table name for this struct type
func (l *LutAngi234) TableName() string {
	return "lut_angi"
}

var LutAngi234Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Анги",
	Identity:  "id",
	DataTable: "lut_angi",
	MainTable: "lut_angi",
	DataModel: new(LutAngi234),
	Data:      new([]LutAngi234),
	MainModel: new(LutAngiMainTable234),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "angi", Label: "Анги"},
	},
	ColumnList: []string{"id", "angi"},
	Filters: map[string]string{
		"id": "",

		"angi": "",
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
	FillVirtualColumns: fillVirtualColumnsLutAngi234,
}

func fillVirtualColumnsLutAngi234(rowsPre interface{}) interface{} {
	return rowsPre
}
