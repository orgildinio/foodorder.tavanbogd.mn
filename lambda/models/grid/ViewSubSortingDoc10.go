package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubSortingDoc10 struct {
	DocumentSorting    *string `gorm:"column:document_sorting" json:"document_sorting"`
	ID                 *int    `gorm:"column:id" json:"id"`
	SubDocumentSorting *string `gorm:"column:sub_document_sorting" json:"sub_document_sorting"`
}

type LutSubDocumentSortingMainTable10 struct {
	DocumentSortingID  *int    `gorm:"column:document_sorting_id" json:"document_sorting_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubDocumentSorting *string `gorm:"column:sub_document_sorting" json:"sub_document_sorting"`
}

func (l *LutSubDocumentSortingMainTable10) TableName() string {
	return "lut_sub_document_sorting"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubSortingDoc10) TableName() string {
	return "view_sub_sorting_doc"
}

var ViewSubSortingDoc10Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Эрх зүй бодлогын баримт бичиг дэд ангилал ",
	Identity:  "id",
	DataTable: "view_sub_sorting_doc",
	MainTable: "lut_sub_document_sorting",
	DataModel: new(ViewSubSortingDoc10),
	Data:      new([]ViewSubSortingDoc10),
	MainModel: new(LutSubDocumentSortingMainTable10),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "sub_document_sorting", Label: "Эрх зүй бодлогын баримт бичиг ангилал "},
		datagrid.Column{Model: "document_sorting", Label: "Эрх зүй бодлогын баримт бичиг дэд ангилал "},
	},
	ColumnList: []string{"id", "sub_document_sorting", "document_sorting"},
	Filters: map[string]string{
		"id": "",

		"sub_document_sorting": "",

		"document_sorting": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubSortingDoc10,
}

func fillVirtualColumnsViewSubSortingDoc10(rowsPre interface{}) interface{} {
	return rowsPre
}
