package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDedUndsenTorol310 struct {
	DedUndsenTorol *string `gorm:"column:ded_undsen_torol" json:"ded_undsen_torol"`
	ID             *int    `gorm:"column:id" json:"id"`
	UndsenTorol    *string `gorm:"column:undsen_torol" json:"undsen_torol"`
}

type LutDedUndsenTorolMainTable310 struct {
	DedUndsenTorol *string `gorm:"column:ded_undsen_torol" json:"ded_undsen_torol"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndsenTorolID  *int    `gorm:"column:undsen_torol_id" json:"undsen_torol_id"`
}

func (l *LutDedUndsenTorolMainTable310) TableName() string {
	return "lut_ded_undsen_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDedUndsenTorol310) TableName() string {
	return "view_ded_undsen_torol"
}

var ViewDedUndsenTorol310Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Үндсэн дэд төрөл ",
	Identity:  "id",
	DataTable: "view_ded_undsen_torol",
	MainTable: "lut_ded_undsen_torol",
	DataModel: new(ViewDedUndsenTorol310),
	Data:      new([]ViewDedUndsenTorol310),
	MainModel: new(LutDedUndsenTorolMainTable310),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "undsen_torol", Label: "Үндсэн төрөл "},
		datagrid.Column{Model: "ded_undsen_torol", Label: "Үндсэн дэд төрөл "},
	},
	ColumnList: []string{"id", "undsen_torol", "ded_undsen_torol"},
	Filters: map[string]string{
		"id": "",

		"undsen_torol_id": "",

		"undsen_torol": "",

		"ded_undsen_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDedUndsenTorol310,
}

func fillVirtualColumnsViewDedUndsenTorol310(rowsPre interface{}) interface{} {
	return rowsPre
}
