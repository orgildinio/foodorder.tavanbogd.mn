package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutUilAjilgaaniiChiglelAan298 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UilAjilgaaniiChiglel *string `gorm:"column:uil_ajilgaanii_chiglel" json:"uil_ajilgaanii_chiglel"`
}

type LutUilAjilgaaniiChiglelAanMainTable298 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UilAjilgaaniiChiglel *string `gorm:"column:uil_ajilgaanii_chiglel" json:"uil_ajilgaanii_chiglel"`
}

func (l *LutUilAjilgaaniiChiglelAanMainTable298) TableName() string {
	return "lut_uil_ajilgaanii_chiglel_aan"
}

//  TableName sets the insert table name for this struct type
func (l *LutUilAjilgaaniiChiglelAan298) TableName() string {
	return "lut_uil_ajilgaanii_chiglel_aan"
}

var LutUilAjilgaaniiChiglelAan298Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L ТЗ БҮХИЙ ААН үйл ажилгааны төрөл",
	Identity:  "id",
	DataTable: "lut_uil_ajilgaanii_chiglel_aan",
	MainTable: "lut_uil_ajilgaanii_chiglel_aan",
	DataModel: new(LutUilAjilgaaniiChiglelAan298),
	Data:      new([]LutUilAjilgaaniiChiglelAan298),
	MainModel: new(LutUilAjilgaaniiChiglelAanMainTable298),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "uil_ajilgaanii_chiglel", Label: "Үйл ажилгааны төрөл"},
	},
	ColumnList: []string{"id", "uil_ajilgaanii_chiglel"},
	Filters: map[string]string{
		"id": "",

		"uil_ajilgaanii_chiglel": "",
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
	FillVirtualColumns: fillVirtualColumnsLutUilAjilgaaniiChiglelAan298,
}

func fillVirtualColumnsLutUilAjilgaaniiChiglelAan298(rowsPre interface{}) interface{} {
	return rowsPre
}
