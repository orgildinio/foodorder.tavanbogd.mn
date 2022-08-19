package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAyultUzegdelTorol239 struct {
	AyultUzegdelTorol *string `gorm:"column:ayult_uzegdel_torol" json:"ayult_uzegdel_torol"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAyultUzegdelTorolMainTable239 struct {
	AyultUzegdelTorol *string `gorm:"column:ayult_uzegdel_torol" json:"ayult_uzegdel_torol"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAyultUzegdelTorolMainTable239) TableName() string {
	return "lut_ayult_uzegdel_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutAyultUzegdelTorol239) TableName() string {
	return "lut_ayult_uzegdel_torol"
}

var LutAyultUzegdelTorol239Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Аюулт үзэгдийн төрөл",
	Identity:  "id",
	DataTable: "lut_ayult_uzegdel_torol",
	MainTable: "lut_ayult_uzegdel_torol",
	DataModel: new(LutAyultUzegdelTorol239),
	Data:      new([]LutAyultUzegdelTorol239),
	MainModel: new(LutAyultUzegdelTorolMainTable239),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ayult_uzegdel_torol", Label: "Аюулт үзэгдийн төрөл"},
	},
	ColumnList: []string{"id", "ayult_uzegdel_torol"},
	Filters: map[string]string{
		"id": "",

		"ayult_uzegdel_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutAyultUzegdelTorol239,
}

func fillVirtualColumnsLutAyultUzegdelTorol239(rowsPre interface{}) interface{} {
	return rowsPre
}
