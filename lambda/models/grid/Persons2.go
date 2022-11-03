package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}
var _ = gorm.DB{}

type Persons2 struct {
	CreatedAt time.Time      `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	DeletedAt gorm.DeletedAt `gorm:"column:DELETED_AT" json:"DELETED_AT"`
	FirstName string         `gorm:"column:FIRST_NAME" json:"FIRST_NAME"`
	LastName  string         `gorm:"column:LAST_NAME" json:"LAST_NAME"`
	PersonID  int64          `gorm:"column:PERSON_ID;primaryKey;autoIncrement" json:"PERSON_ID"`
	UpdatedAt time.Time      `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
}

type PersonsMainTable2 struct {
	CreatedAt time.Time      `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	DeletedAt gorm.DeletedAt `gorm:"column:DELETED_AT" json:"DELETED_AT"`
	FirstName string         `gorm:"column:FIRST_NAME" json:"FIRST_NAME"`
	LastName  string         `gorm:"column:LAST_NAME" json:"LAST_NAME"`
	PersonID  int64          `gorm:"column:PERSON_ID;primaryKey;autoIncrement" json:"PERSON_ID"`
	UpdatedAt time.Time      `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
}

func (p *PersonsMainTable2) TableName() string {
	return "PERSONS"
}

// TableName sets the insert table name for this struct type
func (p *Persons2) TableName() string {
	return "PERSONS"
}

var Persons2Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Хүн",
	Identity:  "PERSON_ID",
	DataTable: "PERSONS",
	MainTable: "PERSONS",
	DataModel: new(Persons2),
	Data:      new([]Persons2),
	MainModel: new(PersonsMainTable2),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "LAST_NAME", Label: "Овог"},
		datagrid.Column{Model: "FIRST_NAME", Label: "Нэр"},
		datagrid.Column{Model: "CREATED_AT", Label: "Бүргэсэн"},
		datagrid.Column{Model: "UPDATED_AT", Label: "Зассан"},
	},
	ColumnList: []string{"PERSON_ID", "LAST_NAME", "FIRST_NAME", "CREATED_AT", "UPDATED_AT", "DELETED_AT"},
	Filters: map[string]string{
		"PERSON_ID": "",

		"LAST_NAME": "",

		"FIRST_NAME": "",

		"CREATED_AT": "",

		"UPDATED_AT": "",

		"DELETED_AT": "",
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
	FillVirtualColumns: fillVirtualColumnsPersons2,
}

func fillVirtualColumnsPersons2(rowsPre interface{}) interface{} {
	return rowsPre
}
