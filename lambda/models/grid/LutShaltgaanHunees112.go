package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutShaltgaanHunees112 struct {
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LutShaltgaanHunees *string `gorm:"column:lut_shaltgaan_hunees" json:"lut_shaltgaan_hunees"`
}

type LutShaltgaanHuneesMainTable112 struct {
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LutShaltgaanHunees *string `gorm:"column:lut_shaltgaan_hunees" json:"lut_shaltgaan_hunees"`
}

func (l *LutShaltgaanHuneesMainTable112) TableName() string {
	return "lut_shaltgaan_hunees"
}

//  TableName sets the insert table name for this struct type
func (l *LutShaltgaanHunees112) TableName() string {
	return "lut_shaltgaan_hunees"
}

var LutShaltgaanHunees112Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Төрөл /Хүнээс/",
	Identity:  "id",
	DataTable: "lut_shaltgaan_hunees",
	MainTable: "lut_shaltgaan_hunees",
	DataModel: new(LutShaltgaanHunees112),
	Data:      new([]LutShaltgaanHunees112),
	MainModel: new(LutShaltgaanHuneesMainTable112),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "lut_shaltgaan_hunees", Label: "Төрөл /Хүнээс/"},
	},
	ColumnList: []string{"id", "lut_shaltgaan_hunees"},
	Filters: map[string]string{
		"id": "",

		"lut_shaltgaan_hunees": "",
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
	FillVirtualColumns: fillVirtualColumnsLutShaltgaanHunees112,
}

func fillVirtualColumnsLutShaltgaanHunees112(rowsPre interface{}) interface{} {
	return rowsPre
}
