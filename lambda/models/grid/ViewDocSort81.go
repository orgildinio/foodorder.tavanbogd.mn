package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDocSort81 struct {
	DocumentSorting    *string `gorm:"column:document_sorting" json:"document_sorting"`
	ID                 *int    `gorm:"column:id" json:"id"`
	SubDocumentSorting *string `gorm:"column:sub_document_sorting" json:"sub_document_sorting"`
}

type LutDocumentSortingMainTable81 struct {
	DocumentSorting   *string `gorm:"column:document_sorting" json:"document_sorting"`
	DocumentSortingEn *string `gorm:"column:document_sorting_en" json:"document_sorting_en"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDocumentSortingMainTable81) TableName() string {
	return "lut_document_sorting"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDocSort81) TableName() string {
	return "view_doc_sort"
}

var ViewDocSort81Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Эрх зүй бодлогын баримт бичиг ангилал",
	Identity:  "id",
	DataTable: "view_doc_sort",
	MainTable: "lut_document_sorting",
	DataModel: new(ViewDocSort81),
	Data:      new([]ViewDocSort81),
	MainModel: new(LutDocumentSortingMainTable81),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "document_sorting", Label: "Ангилал"},
		datagrid.Column{Model: "sub_document_sorting", Label: "Дэд ангилал"},
	},
	ColumnList: []string{"id", "document_sorting", "sub_document_sorting"},
	Filters: map[string]string{
		"id": "",

		"document_sorting": "",

		"sub_document_sorting": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDocSort81,
}

func fillVirtualColumnsViewDocSort81(rowsPre interface{}) interface{} {
	return rowsPre
}
