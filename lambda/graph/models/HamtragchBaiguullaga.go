package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type HamtragchBaiguullaga struct {
	BNer      *string    `gorm:"column:b_ner" json:"b_ner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Link      *string    `gorm:"column:link" json:"link"`
	Logo      *string    `gorm:"column:logo" json:"logo"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (h *HamtragchBaiguullaga) TableName() string {
	return "hamtragch_baiguullaga"
}
