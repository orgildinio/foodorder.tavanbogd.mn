package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubSubSubDamageType316 struct {
	ID                  *int    `gorm:"column:id" json:"id"`
	SubSubSubDamageType *string `gorm:"column:sub_sub_sub_damage_type" json:"sub_sub_sub_damage_type"`
	SubsubDamageType    *string `gorm:"column:subsub_damage_type" json:"subsub_damage_type"`
}

type LutSubSubSubDamageTypeMainTable316 struct {
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubSubDamageTypeID  *int    `gorm:"column:sub_sub_damage_type_id" json:"sub_sub_damage_type_id"`
	SubSubSubDamageType *string `gorm:"column:sub_sub_sub_damage_type" json:"sub_sub_sub_damage_type"`
}

func (l *LutSubSubSubDamageTypeMainTable316) TableName() string {
	return "lut_sub_sub_sub_damage_type"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubSubSubDamageType316) TableName() string {
	return "view_sub_sub_sub_damage_type"
}

var ViewSubSubSubDamageType316Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хохирлын дэд дэд дэд төрөл ",
	Identity:  "id",
	DataTable: "view_sub_sub_sub_damage_type",
	MainTable: "lut_sub_sub_sub_damage_type",
	DataModel: new(ViewSubSubSubDamageType316),
	Data:      new([]ViewSubSubSubDamageType316),
	MainModel: new(LutSubSubSubDamageTypeMainTable316),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "subsub_damage_type", Label: "Хохирлын дэд дэд төрөл "},
		datagrid.Column{Model: "sub_sub_sub_damage_type", Label: "Хохирлын дэд дэд дэд төрөл "},
	},
	ColumnList: []string{"id", "subsub_damage_type", "sub_sub_sub_damage_type"},
	Filters: map[string]string{
		"id": "",

		"sub_sub_damage_type_id": "",

		"subsub_damage_type": "",

		"sub_sub_sub_damage_type": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubSubSubDamageType316,
}

func fillVirtualColumnsViewSubSubSubDamageType316(rowsPre interface{}) interface{} {
	return rowsPre
}
