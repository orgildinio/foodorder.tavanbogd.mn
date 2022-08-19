package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutInfoSource137 struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	InfoSource *string `gorm:"column:info_source" json:"info_source"`
}

type LutInfoSourceMainTable137 struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	InfoSource *string `gorm:"column:info_source" json:"info_source"`
}

func (l *LutInfoSourceMainTable137) TableName() string {
	return "lut_info_source"
}

//  TableName sets the insert table name for this struct type
func (l *LutInfoSource137) TableName() string {
	return "lut_info_source"
}

var LutInfoSource137Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Мэдээллийг хүлээн авсан суваг",
	Identity:  "id",
	DataTable: "lut_info_source",
	MainTable: "lut_info_source",
	DataModel: new(LutInfoSource137),
	Data:      new([]LutInfoSource137),
	MainModel: new(LutInfoSourceMainTable137),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "info_source", Label: "L Мэдээллийг хүлээн авсан суваг"},
	},
	ColumnList: []string{"id", "info_source"},
	Filters: map[string]string{
		"id": "",

		"info_source": "",
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
	FillVirtualColumns: fillVirtualColumnsLutInfoSource137,
}

func fillVirtualColumnsLutInfoSource137(rowsPre interface{}) interface{} {
	return rowsPre
}
