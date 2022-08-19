package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutFire107 struct {
	Fire *string `gorm:"column:fire" json:"fire"`
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutFireMainTable107 struct {
	Fire *string `gorm:"column:fire" json:"fire"`
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutFireMainTable107) TableName() string {
	return "lut_fire"
}

//  TableName sets the insert table name for this struct type
func (l *LutFire107) TableName() string {
	return "lut_fire"
}

var LutFire107Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Гал түймрийг унтраахад ашигласан арга",
	Identity:  "id",
	DataTable: "lut_fire",
	MainTable: "lut_fire",
	DataModel: new(LutFire107),
	Data:      new([]LutFire107),
	MainModel: new(LutFireMainTable107),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "fire", Label: "Гал түймрийг унтраахад ашигласан арга"},
	},
	ColumnList: []string{"id", "fire"},
	Filters: map[string]string{
		"id": "",

		"fire": "",
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
	FillVirtualColumns: fillVirtualColumnsLutFire107,
}

func fillVirtualColumnsLutFire107(rowsPre interface{}) interface{} {
	return rowsPre
}
