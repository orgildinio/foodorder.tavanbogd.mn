package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAlbanTushaal181 struct {
	AlbanTushaal *string `gorm:"column:alban_tushaal" json:"alban_tushaal"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAlbanTushaalMainTable181 struct {
	AlbanTushaal *string `gorm:"column:alban_tushaal" json:"alban_tushaal"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAlbanTushaalMainTable181) TableName() string {
	return "lut_alban_tushaal"
}

//  TableName sets the insert table name for this struct type
func (l *LutAlbanTushaal181) TableName() string {
	return "lut_alban_tushaal"
}

var LutAlbanTushaal181Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Албан тушаал",
	Identity:  "id",
	DataTable: "lut_alban_tushaal",
	MainTable: "lut_alban_tushaal",
	DataModel: new(LutAlbanTushaal181),
	Data:      new([]LutAlbanTushaal181),
	MainModel: new(LutAlbanTushaalMainTable181),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "alban_tushaal", Label: "Албан тушаал"},
	},
	ColumnList: []string{"id", "alban_tushaal"},
	Filters: map[string]string{
		"id": "",

		"alban_tushaal": "",
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
	FillVirtualColumns: fillVirtualColumnsLutAlbanTushaal181,
}

func fillVirtualColumnsLutAlbanTushaal181(rowsPre interface{}) interface{} {
	return rowsPre
}
