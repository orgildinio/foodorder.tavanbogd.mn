package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDedDedAyulTorol328 struct {
	ID                     *int    `gorm:"column:id" json:"id"`
	LsubMaindisasterType   *string `gorm:"column:lsub_maindisaster_type" json:"lsub_maindisaster_type"`
	SubSubMaindisasterType *string `gorm:"column:sub_sub_maindisaster_type" json:"sub_sub_maindisaster_type"`
}

type LutSubSubMaindisasterTypeMainTable328 struct {
	ID                     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubMaindisasterTypeID  *int    `gorm:"column:sub_maindisaster_type_id" json:"sub_maindisaster_type_id"`
	SubSubMaindisasterType *string `gorm:"column:sub_sub_maindisaster_type" json:"sub_sub_maindisaster_type"`
}

func (l *LutSubSubMaindisasterTypeMainTable328) TableName() string {
	return "lut_sub_sub_maindisaster_type"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDedDedAyulTorol328) TableName() string {
	return "view_ded_ded_ayul_torol"
}

var ViewDedDedAyulTorol328Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Аюулын дэд дэд төрөл",
	Identity:  "id",
	DataTable: "view_ded_ded_ayul_torol",
	MainTable: "lut_sub_sub_maindisaster_type",
	DataModel: new(ViewDedDedAyulTorol328),
	Data:      new([]ViewDedDedAyulTorol328),
	MainModel: new(LutSubSubMaindisasterTypeMainTable328),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "lsub_maindisaster_type", Label: "Аюулын дэд төрөл"},
		datagrid.Column{Model: "sub_sub_maindisaster_type", Label: "Аюулын дэд дэд төрөл"},
	},
	ColumnList: []string{"id", "lsub_maindisaster_type", "sub_sub_maindisaster_type"},
	Filters: map[string]string{
		"id": "",

		"sub_maindisaster_type_id": "",

		"lsub_maindisaster_type": "",

		"sub_sub_maindisaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDedDedAyulTorol328,
}

func fillVirtualColumnsViewDedDedAyulTorol328(rowsPre interface{}) interface{} {
	return rowsPre
}
