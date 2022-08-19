package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDedNervegdsenHun146 struct {
	DedNervegdsenHun *string `gorm:"column:ded_nervegdsen_hun" json:"ded_nervegdsen_hun"`
	ID               *int    `gorm:"column:id" json:"id"`
	NervgdsenHun     *string `gorm:"column:nervgdsen_hun" json:"nervgdsen_hun"`
}

type LutDedNervegdsenHunMainTable146 struct {
	DedNervegdsenHun *string `gorm:"column:ded_nervegdsen_hun" json:"ded_nervegdsen_hun"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NervegdsenHunID  *int    `gorm:"column:nervegdsen_hun_id" json:"nervegdsen_hun_id"`
}

func (l *LutDedNervegdsenHunMainTable146) TableName() string {
	return "lut_ded_nervegdsen_hun"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDedNervegdsenHun146) TableName() string {
	return "view_ded_nervegdsen_hun"
}

var ViewDedNervegdsenHun146Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Нэрвэгдсэн хүн амын дэд мэдээлэл",
	Identity:  "id",
	DataTable: "view_ded_nervegdsen_hun",
	MainTable: "lut_ded_nervegdsen_hun",
	DataModel: new(ViewDedNervegdsenHun146),
	Data:      new([]ViewDedNervegdsenHun146),
	MainModel: new(LutDedNervegdsenHunMainTable146),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "nervgdsen_hun", Label: "Нэрвэгдсэн хүн амын мэдээлэл"},
		datagrid.Column{Model: "ded_nervegdsen_hun", Label: "Нэрвэгдсэн хүн амын дэд мэдээлэл"},
	},
	ColumnList: []string{"id", "nervgdsen_hun", "ded_nervegdsen_hun"},
	Filters: map[string]string{
		"id": "",

		"nervegdsen_hun_id": "",

		"nervgdsen_hun": "",

		"ded_nervegdsen_hun": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDedNervegdsenHun146,
}

func fillVirtualColumnsViewDedNervegdsenHun146(rowsPre interface{}) interface{} {
	return rowsPre
}
