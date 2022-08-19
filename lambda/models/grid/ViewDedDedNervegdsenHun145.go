package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDedDedNervegdsenHun145 struct {
	DedDedTorol      *string `gorm:"column:ded_ded_torol" json:"ded_ded_torol"`
	DedNervegdsenHun *string `gorm:"column:ded_nervegdsen_hun" json:"ded_nervegdsen_hun"`
	ID               *int    `gorm:"column:id" json:"id"`
}

type LutSubSubNervgdsenHunMainTable145 struct {
	DedDedTorol *string `gorm:"column:ded_ded_torol" json:"ded_ded_torol"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NDedTorolID *int    `gorm:"column:n_ded_torol_id" json:"n_ded_torol_id"`
}

func (l *LutSubSubNervgdsenHunMainTable145) TableName() string {
	return "lut_sub_sub_nervgdsen_hun"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDedDedNervegdsenHun145) TableName() string {
	return "view_ded_ded_nervegdsen_hun"
}

var ViewDedDedNervegdsenHun145Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Нэрвэгдсэн хүн амын дэд дэд мэдээлэл",
	Identity:  "id",
	DataTable: "view_ded_ded_nervegdsen_hun",
	MainTable: "lut_sub_sub_nervgdsen_hun",
	DataModel: new(ViewDedDedNervegdsenHun145),
	Data:      new([]ViewDedDedNervegdsenHun145),
	MainModel: new(LutSubSubNervgdsenHunMainTable145),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ded_nervegdsen_hun", Label: "Нэрвэгдсэн хүн амын дэд мэдээлэл"},
		datagrid.Column{Model: "ded_ded_torol", Label: "Нэрвэгдсэн хүн амын дэд дэд мэдээлэл"},
	},
	ColumnList: []string{"id", "ded_nervegdsen_hun", "ded_ded_torol"},
	Filters: map[string]string{
		"id": "",

		"n_ded_torol_id": "",

		"ded_nervegdsen_hun": "",

		"ded_ded_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDedDedNervegdsenHun145,
}

func fillVirtualColumnsViewDedDedNervegdsenHun145(rowsPre interface{}) interface{} {
	return rowsPre
}
