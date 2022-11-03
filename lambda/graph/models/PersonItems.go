package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}
var _ = gorm.DB{}

type PersonItems struct {
	ID              int64  `gorm:"column:ID;primaryKey;autoIncrement" json:"ID"`
	ItemDescription string `gorm:"column:ITEM_DESCRIPTION" json:"ITEM_DESCRIPTION"`
	ItemName        string `gorm:"column:ITEM_NAME" json:"ITEM_NAME"`
	PersonID        int64  `gorm:"column:PERSON_ID" json:"PERSON_ID"`
}

// TableName sets the insert table name for this struct type
func (p *PersonItems) TableName() string {
	return "PERSON_ITEMS"
}
