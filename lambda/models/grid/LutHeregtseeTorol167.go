package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHeregtseeTorol167 struct {
	HeregtseeTorol *string `gorm:"column:heregtsee_torol" json:"heregtsee_torol"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutHeregtseeTorolMainTable167 struct {
	HeregtseeTorol *string `gorm:"column:heregtsee_torol" json:"heregtsee_torol"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHeregtseeTorolMainTable167) TableName() string {
	return "lut_heregtsee_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutHeregtseeTorol167) TableName() string {
	return "lut_heregtsee_torol"
}

var LutHeregtseeTorol167Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хэрэгцээний төрөл",
	Identity:  "id",
	DataTable: "lut_heregtsee_torol",
	MainTable: "lut_heregtsee_torol",
	DataModel: new(LutHeregtseeTorol167),
	Data:      new([]LutHeregtseeTorol167),
	MainModel: new(LutHeregtseeTorolMainTable167),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "heregtsee_torol", Label: "Хэрэгцээний төрөл"},
	},
	ColumnList: []string{"id", "heregtsee_torol"},
	Filters: map[string]string{
		"id": "",

		"heregtsee_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutHeregtseeTorol167,
}

func fillVirtualColumnsLutHeregtseeTorol167(rowsPre interface{}) interface{} {
	return rowsPre
}
