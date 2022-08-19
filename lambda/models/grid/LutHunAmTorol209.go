package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHunAmTorol209 struct {
	HunAmTorol *string `gorm:"column:hun_am_torol" json:"hun_am_torol"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutHunAmTorolMainTable209 struct {
	HunAmTorol *string `gorm:"column:hun_am_torol" json:"hun_am_torol"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHunAmTorolMainTable209) TableName() string {
	return "lut_hun_am_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutHunAmTorol209) TableName() string {
	return "lut_hun_am_torol"
}

var LutHunAmTorol209Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Хүн амын төрөл ",
	Identity:  "id",
	DataTable: "lut_hun_am_torol",
	MainTable: "lut_hun_am_torol",
	DataModel: new(LutHunAmTorol209),
	Data:      new([]LutHunAmTorol209),
	MainModel: new(LutHunAmTorolMainTable209),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "hun_am_torol", Label: "Хүн амын төрөл "},
	},
	ColumnList: []string{"id", "hun_am_torol"},
	Filters: map[string]string{
		"id": "",

		"hun_am_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutHunAmTorol209,
}

func fillVirtualColumnsLutHunAmTorol209(rowsPre interface{}) interface{} {
	return rowsPre
}
