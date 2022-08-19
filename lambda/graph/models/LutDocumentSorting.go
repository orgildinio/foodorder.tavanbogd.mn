package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDocumentSorting struct {
	DocumentSorting   *string `gorm:"column:document_sorting" json:"document_sorting"`
	DocumentSortingEn *string `gorm:"column:document_sorting_en" json:"document_sorting_en"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDocumentSorting) TableName() string {
	return "lut_document_sorting"
}
