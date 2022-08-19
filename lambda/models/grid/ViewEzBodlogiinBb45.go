package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewEzBodlogiinBb45 struct {
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt           *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DocumentSorting     *string    `gorm:"column:document_sorting" json:"document_sorting"`
	DocumentSortingEn   *string    `gorm:"column:document_sorting_en" json:"document_sorting_en"`
	EzBarimtBichgiinNer *string    `gorm:"column:ez_barimt_bichgiin_ner" json:"ez_barimt_bichgiin_ner"`
	ID                  *int       `gorm:"column:id" json:"id"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type EzBodlogiinBbMainTable45 struct {
	BarimtNerEn         *string    `gorm:"column:barimt_ner_en" json:"barimt_ner_en"`
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt           *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DocumentSortingID   *int       `gorm:"column:document_sorting_id" json:"document_sorting_id"`
	EzBarimtBichgiinNer *string    `gorm:"column:ez_barimt_bichgiin_ner" json:"ez_barimt_bichgiin_ner"`
	FileEn              *string    `gorm:"column:file_en" json:"file_en"`
	Hawsralt            *string    `gorm:"column:hawsralt" json:"hawsralt"`
	ID                  int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubDocumentSorting  *int       `gorm:"column:sub_document_sorting" json:"sub_document_sorting"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *EzBodlogiinBbMainTable45) TableName() string {
	return "ez_bodlogiin_bb"
}

//  TableName sets the insert table name for this struct type
func (v *ViewEzBodlogiinBb45) TableName() string {
	return "view_ez_bodlogiin_bb"
}

var ViewEzBodlogiinBb45Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Эрх зүй бодлогын баримт бичиг",
	Identity:  "id",
	DataTable: "view_ez_bodlogiin_bb",
	MainTable: "ez_bodlogiin_bb",
	DataModel: new(ViewEzBodlogiinBb45),
	Data:      new([]ViewEzBodlogiinBb45),
	MainModel: new(EzBodlogiinBbMainTable45),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "document_sorting", Label: "Ангилал"},
		datagrid.Column{Model: "ez_barimt_bichgiin_ner", Label: "Эрх зүйн баримт бичгийн нэр"},
		datagrid.Column{Model: "document_sorting_en", Label: "Legal and policy name"},
	},
	ColumnList: []string{"id", "document_sorting", "ez_barimt_bichgiin_ner", "document_sorting_en"},
	Filters: map[string]string{
		"id": "",

		"document_sorting": "",

		"ez_barimt_bichgiin_ner": "",

		"document_sorting_en": "",

		"sub_document_sorting_id": "",

		"document_sorting_id": "",

		"hawsralt": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"barimt_ner_en": "",

		"file_en": "",
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
	FillVirtualColumns: fillVirtualColumnsViewEzBodlogiinBb45,
}

func fillVirtualColumnsViewEzBodlogiinBb45(rowsPre interface{}) interface{} {
	return rowsPre
}
