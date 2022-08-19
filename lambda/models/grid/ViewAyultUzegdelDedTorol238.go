package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewAyultUzegdelDedTorol238 struct {
	AyultUzegdelDedTorol *string `gorm:"column:ayult_uzegdel_ded_torol" json:"ayult_uzegdel_ded_torol"`
	AyultUzegdelTorol    *string `gorm:"column:ayult_uzegdel_torol" json:"ayult_uzegdel_torol"`
	ID                   *int    `gorm:"column:id" json:"id"`
}

type LutAyultUzegdelDedTorolMainTable238 struct {
	AyultUzegdelDedTorol *string `gorm:"column:ayult_uzegdel_ded_torol" json:"ayult_uzegdel_ded_torol"`
	AyultUzegdelTorolID  *int    `gorm:"column:ayult_uzegdel_torol_id" json:"ayult_uzegdel_torol_id"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAyultUzegdelDedTorolMainTable238) TableName() string {
	return "lut_ayult_uzegdel_ded_torol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewAyultUzegdelDedTorol238) TableName() string {
	return "view_ayult_uzegdel_ded_torol"
}

var ViewAyultUzegdelDedTorol238Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Аюулт үзэгдийн дэд төрөл",
	Identity:  "id",
	DataTable: "view_ayult_uzegdel_ded_torol",
	MainTable: "lut_ayult_uzegdel_ded_torol",
	DataModel: new(ViewAyultUzegdelDedTorol238),
	Data:      new([]ViewAyultUzegdelDedTorol238),
	MainModel: new(LutAyultUzegdelDedTorolMainTable238),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ayult_uzegdel_torol", Label: "Аюулт үзэгдийн төрөл"},
		datagrid.Column{Model: "ayult_uzegdel_ded_torol", Label: "Аюулт үзэгдийн дэд төрөл"},
	},
	ColumnList: []string{"id", "ayult_uzegdel_torol", "ayult_uzegdel_ded_torol"},
	Filters: map[string]string{
		"id": "",

		"ayult_uzegdel_torol_id": "",

		"ayult_uzegdel_torol": "",

		"ayult_uzegdel_ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewAyultUzegdelDedTorol238,
}

func fillVirtualColumnsViewAyultUzegdelDedTorol238(rowsPre interface{}) interface{} {
	return rowsPre
}
