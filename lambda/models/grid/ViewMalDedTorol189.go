package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewMalDedTorol189 struct {
	ID          *int    `gorm:"column:id" json:"id"`
	MalDedTorol *string `gorm:"column:mal_ded_torol" json:"mal_ded_torol"`
	MalTorol    *string `gorm:"column:mal_torol" json:"mal_torol"`
}

type LutMalDedTorolMainTable189 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalDedTorol *string `gorm:"column:mal_ded_torol" json:"mal_ded_torol"`
	MalTorolID  *int    `gorm:"column:mal_torol_id" json:"mal_torol_id"`
}

func (l *LutMalDedTorolMainTable189) TableName() string {
	return "lut_mal_ded_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewMalDedTorol189) TableName() string {
	return "view_mal_ded_torol"
}

var ViewMalDedTorol189Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Малын дэд төрөл",
	Identity:  "id",
	DataTable: "view_mal_ded_torol",
	MainTable: "lut_mal_ded_torol",
	DataModel: new(ViewMalDedTorol189),
	Data:      new([]ViewMalDedTorol189),
	MainModel: new(LutMalDedTorolMainTable189),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "mal_torol", Label: "Малын төрөл"},
		datagrid.Column{Model: "mal_ded_torol", Label: "Малын дэд төрөл"},
	},
	ColumnList: []string{"id", "mal_torol", "mal_ded_torol"},
	Filters: map[string]string{
		"id": "",

		"mal_torol_id": "",

		"mal_torol": "",

		"mal_ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewMalDedTorol189,
}

func fillVirtualColumnsViewMalDedTorol189(rowsPre interface{}) interface{} {
	return rowsPre
}
