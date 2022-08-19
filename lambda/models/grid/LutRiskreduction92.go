package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutRiskreduction92 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Riskreduction *string `gorm:"column:riskreduction" json:"riskreduction"`
}

type LutRiskreductionMainTable92 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Riskreduction *string `gorm:"column:riskreduction" json:"riskreduction"`
}

func (l *LutRiskreductionMainTable92) TableName() string {
	return "lut_riskreduction"
}

//  TableName sets the insert table name for this struct type
func (l *LutRiskreduction92) TableName() string {
	return "lut_riskreduction"
}

var LutRiskreduction92Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Гамшгийн эрсдэлийг бууруулах зөвлөл",
	Identity:  "id",
	DataTable: "lut_riskreduction",
	MainTable: "lut_riskreduction",
	DataModel: new(LutRiskreduction92),
	Data:      new([]LutRiskreduction92),
	MainModel: new(LutRiskreductionMainTable92),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "riskreduction", Label: "Гамшгийн эрсдэлийг бууруулах зөвлөл"},
	},
	ColumnList: []string{"id", "riskreduction"},
	Filters: map[string]string{
		"id": "",

		"riskreduction": "",
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
	FillVirtualColumns: fillVirtualColumnsLutRiskreduction92,
}

func fillVirtualColumnsLutRiskreduction92(rowsPre interface{}) interface{} {
	return rowsPre
}
