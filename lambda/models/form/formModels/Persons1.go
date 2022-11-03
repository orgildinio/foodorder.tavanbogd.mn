package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}
var _ = gorm.DB{}

type Persons1 struct {
	CreatedAt time.Time `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	FirstName string    `gorm:"column:FIRST_NAME" json:"FIRST_NAME"`
	LastName  string    `gorm:"column:LAST_NAME" json:"LAST_NAME"`
	PersonID  int64     `gorm:"column:PERSON_ID;primaryKey;autoIncrement" json:"PERSON_ID"`
	UpdatedAt time.Time `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
}

type PersonItemsPersons1 struct {
	ID              int64  `gorm:"column:ID;primaryKey;autoIncrement" json:"ID"`
	ItemDescription string `gorm:"column:ITEM_DESCRIPTION" json:"ITEM_DESCRIPTION"`
	ItemName        string `gorm:"column:ITEM_NAME" json:"ITEM_NAME"`
	PersonID        int64  `gorm:"column:PERSON_ID" json:"PERSON_ID"`
}

func (p *PersonItemsPersons1) TableName() string {
	return "PERSON_ITEMS"
}

// TableName sets the insert table name for this struct type
func (p *Persons1) TableName() string {
	return "PERSONS"
}
