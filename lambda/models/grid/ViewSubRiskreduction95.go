package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubRiskreduction95 struct {
	ID               *int    `gorm:"column:id" json:"id"`
	Riskreduction    *string `gorm:"column:riskreduction" json:"riskreduction"`
	SubRiskreduction *string `gorm:"column:sub_riskreduction" json:"sub_riskreduction"`
}

type LutSubRiskreductionMainTable95 struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RiskreductionID  *int    `gorm:"column:riskreduction_id" json:"riskreduction_id"`
	SubRiskreduction *string `gorm:"column:sub_riskreduction" json:"sub_riskreduction"`
}

func (l *LutSubRiskreductionMainTable95) TableName() string {
	return "lut_sub_riskreduction"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubRiskreduction95) TableName() string {
	return "view_sub_riskreduction"
}

var ViewSubRiskreduction95Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Гамшгийн эрсдэлийг бууруулах дэд зөвлөл",
	Identity:  "id",
	DataTable: "view_sub_riskreduction",
	MainTable: "lut_sub_riskreduction",
	DataModel: new(ViewSubRiskreduction95),
	Data:      new([]ViewSubRiskreduction95),
	MainModel: new(LutSubRiskreductionMainTable95),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "riskreduction", Label: "L Гамшгийн эрсдэлийг бууруулах зөвлөл"},
		datagrid.Column{Model: "sub_riskreduction", Label: "L Гамшгийн эрсдэлийг бууруулах дэд зөвлөл"},
	},
	ColumnList: []string{"id", "riskreduction", "sub_riskreduction"},
	Filters: map[string]string{
		"id": "",

		"riskreduction_id": "",

		"riskreduction": "",

		"sub_riskreduction": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubRiskreduction95,
}

func fillVirtualColumnsViewSubRiskreduction95(rowsPre interface{}) interface{} {
	return rowsPre
}
