package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubDamageType302 struct {
	DamageType    *string `gorm:"column:damage_type" json:"damage_type"`
	ID            *int    `gorm:"column:id" json:"id"`
	SubDamageType *string `gorm:"column:sub_damage_type" json:"sub_damage_type"`
}

type LutSubDamageTypeMainTable302 struct {
	DamageTypeID  *int    `gorm:"column:damage_type_id" json:"damage_type_id"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubDamageType *string `gorm:"column:sub_damage_type" json:"sub_damage_type"`
}

func (l *LutSubDamageTypeMainTable302) TableName() string {
	return "lut_sub_damage_type"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubDamageType302) TableName() string {
	return "view_sub_damage_type"
}

var ViewSubDamageType302Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хохирлын дэд төрөл ",
	Identity:  "id",
	DataTable: "view_sub_damage_type",
	MainTable: "lut_sub_damage_type",
	DataModel: new(ViewSubDamageType302),
	Data:      new([]ViewSubDamageType302),
	MainModel: new(LutSubDamageTypeMainTable302),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "damage_type", Label: "Хохирлын төрөл "},
		datagrid.Column{Model: "sub_damage_type", Label: "Хохирлын дэд төрөл "},
	},
	ColumnList: []string{"id", "damage_type", "sub_damage_type"},
	Filters: map[string]string{
		"id": "",

		"damage_type_id": "",

		"damage_type": "",

		"sub_damage_type": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubDamageType302,
}

func fillVirtualColumnsViewSubDamageType302(rowsPre interface{}) interface{} {
	return rowsPre
}
