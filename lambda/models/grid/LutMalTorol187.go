package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalTorol187 struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalTorol *string `gorm:"column:mal_torol" json:"mal_torol"`
}

type LutMalTorolMainTable187 struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalTorol *string `gorm:"column:mal_torol" json:"mal_torol"`
}

func (l *LutMalTorolMainTable187) TableName() string {
	return "lut_mal_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutMalTorol187) TableName() string {
	return "lut_mal_torol"
}

var LutMalTorol187Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Малын төрөл",
	Identity:  "id",
	DataTable: "lut_mal_torol",
	MainTable: "lut_mal_torol",
	DataModel: new(LutMalTorol187),
	Data:      new([]LutMalTorol187),
	MainModel: new(LutMalTorolMainTable187),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "mal_torol", Label: "Малын төрөл"},
	},
	ColumnList: []string{"id", "mal_torol"},
	Filters: map[string]string{
		"id": "",

		"mal_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMalTorol187,
}

func fillVirtualColumnsLutMalTorol187(rowsPre interface{}) interface{} {
	return rowsPre
}
