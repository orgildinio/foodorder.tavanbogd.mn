package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewHeregtseeDedDedTorol185 struct {
	HeregtseeDedDedTorol *string `gorm:"column:heregtsee_ded__ded_torol" json:"heregtsee_ded__ded_torol"`
	HeregtseeDedTorol    *string `gorm:"column:heregtsee_ded_torol" json:"heregtsee_ded_torol"`
	ID                   *int    `gorm:"column:id" json:"id"`
}

type LutHeregtseeDedDedTorolMainTable185 struct {
	HeregtseeDedDedTorol *string `gorm:"column:heregtsee_ded__ded_torol" json:"heregtsee_ded__ded_torol"`
	HeregtseeDedTorolID  *int    `gorm:"column:heregtsee_ded_torol_id" json:"heregtsee_ded_torol_id"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHeregtseeDedDedTorolMainTable185) TableName() string {
	return "lut_heregtsee_ded_ded_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewHeregtseeDedDedTorol185) TableName() string {
	return "view_heregtsee_ded_ded_torol"
}

var ViewHeregtseeDedDedTorol185Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хэрэгцээний дэдийн дэд төрөл",
	Identity:  "id",
	DataTable: "view_heregtsee_ded_ded_torol",
	MainTable: "lut_heregtsee_ded_ded_torol",
	DataModel: new(ViewHeregtseeDedDedTorol185),
	Data:      new([]ViewHeregtseeDedDedTorol185),
	MainModel: new(LutHeregtseeDedDedTorolMainTable185),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "heregtsee_ded_torol", Label: "Хэрэгцээний дэд төрөл"},
		datagrid.Column{Model: "heregtsee_ded__ded_torol", Label: "Хэрэгцээний дэдийн дэд төрөл"},
	},
	ColumnList: []string{"id", "heregtsee_ded_torol", "heregtsee_ded__ded_torol"},
	Filters: map[string]string{
		"id": "",

		"heregtsee_ded_torol_id": "",

		"heregtsee_ded_torol": "",

		"heregtsee_ded__ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewHeregtseeDedDedTorol185,
}

func fillVirtualColumnsViewHeregtseeDedDedTorol185(rowsPre interface{}) interface{} {
	return rowsPre
}
