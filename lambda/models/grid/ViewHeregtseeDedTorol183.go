package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewHeregtseeDedTorol183 struct {
	HeregtseeDedTorol *string `gorm:"column:heregtsee_ded_torol" json:"heregtsee_ded_torol"`
	HeregtseeTorol    *string `gorm:"column:heregtsee_torol" json:"heregtsee_torol"`
	ID                *int    `gorm:"column:id" json:"id"`
}

type LutHeregtseeDedTorolMainTable183 struct {
	HeregtseeDedTorol *string `gorm:"column:heregtsee_ded_torol" json:"heregtsee_ded_torol"`
	HeregtseeTorolID  *int    `gorm:"column:heregtsee_torol_id" json:"heregtsee_torol_id"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHeregtseeDedTorolMainTable183) TableName() string {
	return "lut_heregtsee_ded_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewHeregtseeDedTorol183) TableName() string {
	return "view_heregtsee_ded_torol"
}

var ViewHeregtseeDedTorol183Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хэрэгцээний дэд төрөл",
	Identity:  "id",
	DataTable: "view_heregtsee_ded_torol",
	MainTable: "lut_heregtsee_ded_torol",
	DataModel: new(ViewHeregtseeDedTorol183),
	Data:      new([]ViewHeregtseeDedTorol183),
	MainModel: new(LutHeregtseeDedTorolMainTable183),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "heregtsee_torol", Label: "Хэрэгцээний төрөл"},
		datagrid.Column{Model: "heregtsee_ded_torol", Label: "Хэрэгцээний дэд төрөл"},
	},
	ColumnList: []string{"id", "heregtsee_torol", "heregtsee_ded_torol"},
	Filters: map[string]string{
		"id": "",

		"heregtsee_torol_id": "",

		"heregtsee_torol": "",

		"heregtsee_ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewHeregtseeDedTorol183,
}

func fillVirtualColumnsViewHeregtseeDedTorol183(rowsPre interface{}) interface{} {
	return rowsPre
}
