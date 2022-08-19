package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewAvsanDedArgaHemjee313 struct {
	AvsanArgaDedHemjee *string `gorm:"column:avsan_arga_ded_hemjee" json:"avsan_arga_ded_hemjee"`
	AvsanArgaHemjee    *string `gorm:"column:avsan_arga_hemjee" json:"avsan_arga_hemjee"`
	ID                 *int    `gorm:"column:id" json:"id"`
}

type LutAvsanArgaDedHemjeeMainTable313 struct {
	AvsanArgaDedHemjee *string `gorm:"column:avsan_arga_ded_hemjee" json:"avsan_arga_ded_hemjee"`
	AvsanArgaHemjeeID  *int    `gorm:"column:avsan_arga_hemjee_id" json:"avsan_arga_hemjee_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAvsanArgaDedHemjeeMainTable313) TableName() string {
	return "lut_avsan_arga_ded_hemjee"
}

//  TableName sets the insert table name for this struct type
func (v *ViewAvsanDedArgaHemjee313) TableName() string {
	return "view_avsan_ded_arga_hemjee"
}

var ViewAvsanDedArgaHemjee313Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Малд авсан дэд арга хэмжээ",
	Identity:  "id",
	DataTable: "view_avsan_ded_arga_hemjee",
	MainTable: "lut_avsan_arga_ded_hemjee",
	DataModel: new(ViewAvsanDedArgaHemjee313),
	Data:      new([]ViewAvsanDedArgaHemjee313),
	MainModel: new(LutAvsanArgaDedHemjeeMainTable313),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "avsan_arga_hemjee", Label: "Малд авсан арга хэмжээ"},
		datagrid.Column{Model: "avsan_arga_ded_hemjee", Label: "Малд авсан дэд арга хэмжээ"},
	},
	ColumnList: []string{"id", "avsan_arga_hemjee", "avsan_arga_ded_hemjee"},
	Filters: map[string]string{
		"id": "",

		"avsan_arga_hemjee_id": "",

		"avsan_arga_hemjee": "",

		"avsan_arga_ded_hemjee": "",
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
	FillVirtualColumns: fillVirtualColumnsViewAvsanDedArgaHemjee313,
}

func fillVirtualColumnsViewAvsanDedArgaHemjee313(rowsPre interface{}) interface{} {
	return rowsPre
}
