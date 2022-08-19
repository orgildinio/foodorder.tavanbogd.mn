package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutTejeeverAmitniiTurul194 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TejeeverAmitniiTurul *string `gorm:"column:tejeever_amitnii_turul" json:"tejeever_amitnii_turul"`
}

type LutTejeeverAmitniiTurulMainTable194 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TejeeverAmitniiTurul *string `gorm:"column:tejeever_amitnii_turul" json:"tejeever_amitnii_turul"`
}

func (l *LutTejeeverAmitniiTurulMainTable194) TableName() string {
	return "lut_tejeever_amitnii_turul"
}

//  TableName sets the insert table name for this struct type
func (l *LutTejeeverAmitniiTurul194) TableName() string {
	return "lut_tejeever_amitnii_turul"
}

var LutTejeeverAmitniiTurul194Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Тэжээвэр амьтны төрөл",
	Identity:  "id",
	DataTable: "lut_tejeever_amitnii_turul",
	MainTable: "lut_tejeever_amitnii_turul",
	DataModel: new(LutTejeeverAmitniiTurul194),
	Data:      new([]LutTejeeverAmitniiTurul194),
	MainModel: new(LutTejeeverAmitniiTurulMainTable194),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "tejeever_amitnii_turul", Label: "Тэжээвэр амьтны төрөл"},
	},
	ColumnList: []string{"id", "tejeever_amitnii_turul"},
	Filters: map[string]string{
		"id": "",

		"tejeever_amitnii_turul": "",
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
	FillVirtualColumns: fillVirtualColumnsLutTejeeverAmitniiTurul194,
}

func fillVirtualColumnsLutTejeeverAmitniiTurul194(rowsPre interface{}) interface{} {
	return rowsPre
}
