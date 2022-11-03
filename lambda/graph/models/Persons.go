package models

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}
var _ = gorm.DB{}

type Persons struct {
	CreatedAt   time.Time      `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	DeletedAt   gorm.DeletedAt `gorm:"column:DELETED_AT" json:"DELETED_AT"`
	FirstName   string         `gorm:"column:FIRST_NAME" json:"FIRST_NAME"`
	LastName    string         `gorm:"column:LAST_NAME" json:"LAST_NAME"`
	PersonID    int64          `gorm:"column:PERSON_ID;primaryKey;autoIncrement" json:"PERSON_ID"`
	UpdatedAt   time.Time      `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
	PersonItems []*PersonItems `gorm:"-:all"`
}

// TableName sets the insert table name for this struct type
func (p *Persons) TableName() string {
	return "PERSONS"
}
