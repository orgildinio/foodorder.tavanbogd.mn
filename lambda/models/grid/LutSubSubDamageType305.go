package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubSubDamageType305 struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubsubDamageType *string `gorm:"column:subsub_damage_type" json:"subsub_damage_type"`
}

type LutSubSubDamageTypeMainTable305 struct {
	DamageSubTypeID  *int    `gorm:"column:damage_sub_type_id" json:"damage_sub_type_id"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubsubDamageType *string `gorm:"column:subsub_damage_type" json:"subsub_damage_type"`
}

func (l *LutSubSubDamageTypeMainTable305) TableName() string {
	return "lut_sub_sub_damage_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutSubSubDamageType305) TableName() string {
	return "lut_sub_sub_damage_type"
}

var LutSubSubDamageType305Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хохирлын дэд дэд төрөл ",
	Identity:  "id",
	DataTable: "lut_sub_sub_damage_type",
	MainTable: "lut_sub_sub_damage_type",
	DataModel: new(LutSubSubDamageType305),
	Data:      new([]LutSubSubDamageType305),
	MainModel: new(LutSubSubDamageTypeMainTable305),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "subsub_damage_type", Label: "Хохирлын дэд дэд төрөл "},
	},
	ColumnList: []string{"id", "subsub_damage_type"},
	Filters: map[string]string{
		"id": "",

		"damage_sub_type_id": "",

		"subsub_damage_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutSubSubDamageType305,
}

func fillVirtualColumnsLutSubSubDamageType305(rowsPre interface{}) interface{} {
	return rowsPre
}
