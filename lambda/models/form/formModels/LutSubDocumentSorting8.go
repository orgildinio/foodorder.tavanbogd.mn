package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubDocumentSorting8 struct {
	DocumentSortingID  *int    `gorm:"column:document_sorting_id" json:"document_sorting_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubDocumentSorting *string `gorm:"column:sub_document_sorting" json:"sub_document_sorting"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubDocumentSorting8) TableName() string {
	return "lut_sub_document_sorting"
}
