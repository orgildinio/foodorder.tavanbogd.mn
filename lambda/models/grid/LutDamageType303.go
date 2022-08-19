package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDamageType303 struct {
	DamageType *string `gorm:"column:damage_type" json:"damage_type"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutDamageTypeMainTable303 struct {
	DamageType *string `gorm:"column:damage_type" json:"damage_type"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDamageTypeMainTable303) TableName() string {
	return "lut_damage_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutDamageType303) TableName() string {
	return "lut_damage_type"
}

var LutDamageType303Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Хохирлын төрөл ",
	Identity:  "id",
	DataTable: "lut_damage_type",
	MainTable: "lut_damage_type",
	DataModel: new(LutDamageType303),
	Data:      new([]LutDamageType303),
	MainModel: new(LutDamageTypeMainTable303),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "damage_type", Label: "Хохирлын төрөл "},
	},
	ColumnList: []string{"id", "damage_type"},
	Filters: map[string]string{
		"id": "",

		"damage_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutDamageType303,
}

func fillVirtualColumnsLutDamageType303(rowsPre interface{}) interface{} {
	return rowsPre
}
