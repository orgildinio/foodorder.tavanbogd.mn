package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type EzBodlogiinBb42 struct {
	BarimtNerEn         *string    `gorm:"column:barimt_ner_en" json:"barimt_ner_en"`
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`
	DocumentSortingID   *int       `gorm:"column:document_sorting_id" json:"document_sorting_id"`
	EzBarimtBichgiinNer *string    `gorm:"column:ez_barimt_bichgiin_ner" json:"ez_barimt_bichgiin_ner"`
	FileEn              *string    `gorm:"column:file_en" json:"file_en"`
	Hawsralt            *string    `gorm:"column:hawsralt" json:"hawsralt"`
	ID                  int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubDocumentSorting  *int       `gorm:"column:sub_document_sorting" json:"sub_document_sorting"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (e *EzBodlogiinBb42) TableName() string {
	return "ez_bodlogiin_bb"
}
