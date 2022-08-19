package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewHunAmDedDedTorol215 struct {
	HunAmDedDedTorol *string `gorm:"column:hun_am_ded_ded_torol" json:"hun_am_ded_ded_torol"`
	HunAmDedTorol    *string `gorm:"column:hun_am_ded_torol" json:"hun_am_ded_torol"`
	ID               *int    `gorm:"column:id" json:"id"`
}

type LutHunAmDedDedTorolMainTable215 struct {
	HunAmDedDedTorol *string `gorm:"column:hun_am_ded_ded_torol" json:"hun_am_ded_ded_torol"`
	HunAmDedTorolID  *int    `gorm:"column:hun_am_ded_torol_id" json:"hun_am_ded_torol_id"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHunAmDedDedTorolMainTable215) TableName() string {
	return "lut_hun_am_ded_ded_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewHunAmDedDedTorol215) TableName() string {
	return "view_hun_am_ded_ded_torol"
}

var ViewHunAmDedDedTorol215Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хүн амын дэд дэд төрөл",
	Identity:  "id",
	DataTable: "view_hun_am_ded_ded_torol",
	MainTable: "lut_hun_am_ded_ded_torol",
	DataModel: new(ViewHunAmDedDedTorol215),
	Data:      new([]ViewHunAmDedDedTorol215),
	MainModel: new(LutHunAmDedDedTorolMainTable215),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "hun_am_ded_torol", Label: "Хүн амын дэд төрөл"},
		datagrid.Column{Model: "hun_am_ded_ded_torol", Label: "Хүн амын дэд дэд төрөл"},
	},
	ColumnList: []string{"id", "hun_am_ded_torol", "hun_am_ded_ded_torol"},
	Filters: map[string]string{
		"id": "",

		"hun_am_ded_torol_id": "",

		"hun_am_ded_torol": "",

		"hun_am_ded_ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewHunAmDedDedTorol215,
}

func fillVirtualColumnsViewHunAmDedDedTorol215(rowsPre interface{}) interface{} {
	return rowsPre
}
