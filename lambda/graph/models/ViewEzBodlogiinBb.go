package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewEzBodlogiinBb struct {
	BarimtNerEn          *string    `gorm:"column:barimt_ner_en" json:"barimt_ner_en"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DocumentSorting      *string    `gorm:"column:document_sorting" json:"document_sorting"`
	DocumentSortingEn    *string    `gorm:"column:document_sorting_en" json:"document_sorting_en"`
	DocumentSortingID    *int       `gorm:"column:document_sorting_id" json:"document_sorting_id"`
	EzBarimtBichgiinNer  *string    `gorm:"column:ez_barimt_bichgiin_ner" json:"ez_barimt_bichgiin_ner"`
	FileEn               *string    `gorm:"column:file_en" json:"file_en"`
	Hawsralt             *string    `gorm:"column:hawsralt" json:"hawsralt"`
	ID                   *int       `gorm:"column:id" json:"id"`
	SubDocumentSortingID *int       `gorm:"column:sub_document_sorting_id" json:"sub_document_sorting_id"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (v *ViewEzBodlogiinBb) TableName() string {
	return "view_ez_bodlogiin_bb"
}
