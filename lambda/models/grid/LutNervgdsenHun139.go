package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNervgdsenHun139 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NervgdsenHun *string `gorm:"column:nervgdsen_hun" json:"nervgdsen_hun"`
}

type LutNervgdsenHunMainTable139 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NervgdsenHun *string `gorm:"column:nervgdsen_hun" json:"nervgdsen_hun"`
}

func (l *LutNervgdsenHunMainTable139) TableName() string {
	return "lut_nervgdsen_hun"
}

//  TableName sets the insert table name for this struct type
func (l *LutNervgdsenHun139) TableName() string {
	return "lut_nervgdsen_hun"
}

var LutNervgdsenHun139Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Нэрвэгдсэн хүн амын мэдээлэл ",
	Identity:  "id",
	DataTable: "lut_nervgdsen_hun",
	MainTable: "lut_nervgdsen_hun",
	DataModel: new(LutNervgdsenHun139),
	Data:      new([]LutNervgdsenHun139),
	MainModel: new(LutNervgdsenHunMainTable139),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "nervgdsen_hun", Label: "Нэрвэгдсэн хүн амын мэдээлэл"},
	},
	ColumnList: []string{"id", "nervgdsen_hun"},
	Filters: map[string]string{
		"id": "",

		"nervgdsen_hun": "",
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
	FillVirtualColumns: fillVirtualColumnsLutNervgdsenHun139,
}

func fillVirtualColumnsLutNervgdsenHun139(rowsPre interface{}) interface{} {
	return rowsPre
}
