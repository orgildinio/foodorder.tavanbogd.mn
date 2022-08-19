package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubMaindisasterType88 struct {
	ID                   *int    `gorm:"column:id" json:"id"`
	LsubMaindisasterType *string `gorm:"column:lsub_maindisaster_type" json:"lsub_maindisaster_type"`
	MaindisasterType     *string `gorm:"column:maindisaster_type" json:"maindisaster_type"`
}

type LutSubMaindisasterTypeMainTable88 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LsubMaindisasterType *string `gorm:"column:lsub_maindisaster_type" json:"lsub_maindisaster_type"`
	MaindisasterTypeID   *int    `gorm:"column:maindisaster_type_id" json:"maindisaster_type_id"`
}

func (l *LutSubMaindisasterTypeMainTable88) TableName() string {
	return "lut_sub_maindisaster_type"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubMaindisasterType88) TableName() string {
	return "view_sub_maindisaster_type"
}

var ViewSubMaindisasterType88Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Аюулын дэд төрөл",
	Identity:  "id",
	DataTable: "view_sub_maindisaster_type",
	MainTable: "lut_sub_maindisaster_type",
	DataModel: new(ViewSubMaindisasterType88),
	Data:      new([]ViewSubMaindisasterType88),
	MainModel: new(LutSubMaindisasterTypeMainTable88),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "maindisaster_type", Label: "Аюулын төрөл"},
		datagrid.Column{Model: "lsub_maindisaster_type", Label: "Аюулын дэд төрөл"},
	},
	ColumnList: []string{"id", "maindisaster_type", "lsub_maindisaster_type"},
	Filters: map[string]string{
		"id": "",

		"maindisaster_type_id": "",

		"maindisaster_type": "",

		"lsub_maindisaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubMaindisasterType88,
}

func fillVirtualColumnsViewSubMaindisasterType88(rowsPre interface{}) interface{} {
	return rowsPre
}
