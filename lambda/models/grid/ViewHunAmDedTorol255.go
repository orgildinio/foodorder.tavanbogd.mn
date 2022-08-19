package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewHunAmDedTorol255 struct {
	HunAmDedTorol *string `gorm:"column:hun_am_ded_torol" json:"hun_am_ded_torol"`
	HunAmTorol    *string `gorm:"column:hun_am_torol" json:"hun_am_torol"`
	ID            *int    `gorm:"column:id" json:"id"`
}

type LutHunAmDedTorolMainTable255 struct {
	HunAmDedTorol *string `gorm:"column:hun_am_ded_torol" json:"hun_am_ded_torol"`
	HunAmTorolID  *int    `gorm:"column:hun_am_torol_id" json:"hun_am_torol_id"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHunAmDedTorolMainTable255) TableName() string {
	return "lut_hun_am_ded_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewHunAmDedTorol255) TableName() string {
	return "view_hun_am_ded_torol"
}

var ViewHunAmDedTorol255Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хүн амын дэд төрөл",
	Identity:  "id",
	DataTable: "view_hun_am_ded_torol",
	MainTable: "lut_hun_am_ded_torol",
	DataModel: new(ViewHunAmDedTorol255),
	Data:      new([]ViewHunAmDedTorol255),
	MainModel: new(LutHunAmDedTorolMainTable255),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "hun_am_torol", Label: "Хүн амын төрөл"},
		datagrid.Column{Model: "hun_am_ded_torol", Label: "Хүн амын дэд төрөл"},
	},
	ColumnList: []string{"id", "hun_am_torol", "hun_am_ded_torol"},
	Filters: map[string]string{
		"id": "",

		"hun_am_torol_id": "",

		"hun_am_torol": "",

		"hun_am_ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewHunAmDedTorol255,
}

func fillVirtualColumnsViewHunAmDedTorol255(rowsPre interface{}) interface{} {
	return rowsPre
}
