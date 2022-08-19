package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutShaltgaanGal117 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanGal *string `gorm:"column:shaltgaan_gal" json:"shaltgaan_gal"`
}

type LutShaltgaanGalMainTable117 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanGal *string `gorm:"column:shaltgaan_gal" json:"shaltgaan_gal"`
}

func (l *LutShaltgaanGalMainTable117) TableName() string {
	return "lut_shaltgaan_gal"
}

//  TableName sets the insert table name for this struct type
func (l *LutShaltgaanGal117) TableName() string {
	return "lut_shaltgaan_gal"
}

var LutShaltgaanGal117Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Галын төрөл",
	Identity:  "id",
	DataTable: "lut_shaltgaan_gal",
	MainTable: "lut_shaltgaan_gal",
	DataModel: new(LutShaltgaanGal117),
	Data:      new([]LutShaltgaanGal117),
	MainModel: new(LutShaltgaanGalMainTable117),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "shaltgaan_gal", Label: "Галын төрөл"},
	},
	ColumnList: []string{"id", "shaltgaan_gal"},
	Filters: map[string]string{
		"id": "",

		"shaltgaan_gal": "",
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
	FillVirtualColumns: fillVirtualColumnsLutShaltgaanGal117,
}

func fillVirtualColumnsLutShaltgaanGal117(rowsPre interface{}) interface{} {
	return rowsPre
}
