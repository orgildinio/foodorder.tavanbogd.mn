package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAjillasanBaidalGal220 struct {
	Ajillasan_baidalGal *string `gorm:"column:ajillasan baidal_gal" json:"ajillasan baidal_gal"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAjillasanBaidalGalMainTable220 struct {
	Ajillasan_baidalGal *string `gorm:"column:ajillasan baidal_gal" json:"ajillasan baidal_gal"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAjillasanBaidalGalMainTable220) TableName() string {
	return "lut_ajillasan_baidal_gal"
}

//  TableName sets the insert table name for this struct type
func (l *LutAjillasanBaidalGal220) TableName() string {
	return "lut_ajillasan_baidal_gal"
}

var LutAjillasanBaidalGal220Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Ажилсан байдал",
	Identity:  "id",
	DataTable: "lut_ajillasan_baidal_gal",
	MainTable: "lut_ajillasan_baidal_gal",
	DataModel: new(LutAjillasanBaidalGal220),
	Data:      new([]LutAjillasanBaidalGal220),
	MainModel: new(LutAjillasanBaidalGalMainTable220),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ajillasan baidal_gal", Label: "Ажилсан байдал"},
	},
	ColumnList: []string{"id", "ajillasan baidal_gal"},
	Filters: map[string]string{
		"id": "",

		"ajillasan baidal_gal": "",
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
	FillVirtualColumns: fillVirtualColumnsLutAjillasanBaidalGal220,
}

func fillVirtualColumnsLutAjillasanBaidalGal220(rowsPre interface{}) interface{} {
	return rowsPre
}
