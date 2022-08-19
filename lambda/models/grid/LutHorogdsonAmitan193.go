package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHorogdsonAmitan193 struct {
	HorogdsonAmitan *string `gorm:"column:horogdson_amitan" json:"horogdson_amitan"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutHorogdsonAmitanMainTable193 struct {
	HorogdsonAmitan *string `gorm:"column:horogdson_amitan" json:"horogdson_amitan"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHorogdsonAmitanMainTable193) TableName() string {
	return "lut_horogdson_amitan"
}

//  TableName sets the insert table name for this struct type
func (l *LutHorogdsonAmitan193) TableName() string {
	return "lut_horogdson_amitan"
}

var LutHorogdsonAmitan193Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хорогдсон тэжээвэр амьтан",
	Identity:  "id",
	DataTable: "lut_horogdson_amitan",
	MainTable: "lut_horogdson_amitan",
	DataModel: new(LutHorogdsonAmitan193),
	Data:      new([]LutHorogdsonAmitan193),
	MainModel: new(LutHorogdsonAmitanMainTable193),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "horogdson_amitan", Label: "Хорогдсон тэжээвэр амьтан"},
	},
	ColumnList: []string{"id", "horogdson_amitan"},
	Filters: map[string]string{
		"id": "",

		"horogdson_amitan": "",
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
	FillVirtualColumns: fillVirtualColumnsLutHorogdsonAmitan193,
}

func fillVirtualColumnsLutHorogdsonAmitan193(rowsPre interface{}) interface{} {
	return rowsPre
}
