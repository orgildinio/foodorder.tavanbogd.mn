package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAvsanArgaHemjee192 struct {
	AvsanArgaHemjee *string `gorm:"column:avsan_arga_hemjee" json:"avsan_arga_hemjee"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAvsanArgaHemjeeMainTable192 struct {
	AvsanArgaHemjee *string `gorm:"column:avsan_arga_hemjee" json:"avsan_arga_hemjee"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAvsanArgaHemjeeMainTable192) TableName() string {
	return "lut_avsan_arga_hemjee"
}

//  TableName sets the insert table name for this struct type
func (l *LutAvsanArgaHemjee192) TableName() string {
	return "lut_avsan_arga_hemjee"
}

var LutAvsanArgaHemjee192Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Малд авсан арга хэмжээ",
	Identity:  "id",
	DataTable: "lut_avsan_arga_hemjee",
	MainTable: "lut_avsan_arga_hemjee",
	DataModel: new(LutAvsanArgaHemjee192),
	Data:      new([]LutAvsanArgaHemjee192),
	MainModel: new(LutAvsanArgaHemjeeMainTable192),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "avsan_arga_hemjee", Label: "Малд авсан арга хэмжээ"},
	},
	ColumnList: []string{"id", "avsan_arga_hemjee"},
	Filters: map[string]string{
		"id": "",

		"avsan_arga_hemjee": "",
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
	FillVirtualColumns: fillVirtualColumnsLutAvsanArgaHemjee192,
}

func fillVirtualColumnsLutAvsanArgaHemjee192(rowsPre interface{}) interface{} {
	return rowsPre
}
