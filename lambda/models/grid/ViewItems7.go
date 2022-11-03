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

type ViewItems7 struct {
	FirstName       string `gorm:"column:FIRST_NAME" json:"FIRST_NAME"`
	ID              int64  `gorm:"column:ID" json:"ID"`
	ItemDescription string `gorm:"column:ITEM_DESCRIPTION" json:"ITEM_DESCRIPTION"`
	ItemName        string `gorm:"column:ITEM_NAME" json:"ITEM_NAME"`
	LastName        string `gorm:"column:LAST_NAME" json:"LAST_NAME"`
}

type PersonItemsMainTable7 struct {
	ID              int64  `gorm:"column:ID;primaryKey;autoIncrement" json:"ID"`
	ItemDescription string `gorm:"column:ITEM_DESCRIPTION" json:"ITEM_DESCRIPTION"`
	ItemName        string `gorm:"column:ITEM_NAME" json:"ITEM_NAME"`
	PersonID        int64  `gorm:"column:PERSON_ID" json:"PERSON_ID"`
}

func (p *PersonItemsMainTable7) TableName() string {
	return "PERSON_ITEMS"
}

// TableName sets the insert table name for this struct type
func (v *ViewItems7) TableName() string {
	return "VIEW_ITEMS"
}

var ViewItems7Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Эд зүйл",
	Identity:  "ID",
	DataTable: "VIEW_ITEMS",
	MainTable: "PERSON_ITEMS",
	DataModel: new(ViewItems7),
	Data:      new([]ViewItems7),
	MainModel: new(PersonItemsMainTable7),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ITEM_NAME", Label: "Эд зүйл"},
		datagrid.Column{Model: "ITEM_DESCRIPTION", Label: "Тайлбар"},
		datagrid.Column{Model: "FIRST_NAME", Label: "Эзэн нэр"},
		datagrid.Column{Model: "LAST_NAME", Label: "Эзэн овог"},
	},
	ColumnList: []string{"ID", "ITEM_NAME", "ITEM_DESCRIPTION", "FIRST_NAME", "LAST_NAME"},
	Filters: map[string]string{
		"ID": "",

		"PERSON_ID": "",

		"ITEM_NAME": "",

		"ITEM_DESCRIPTION": "",

		"FIRST_NAME": "",

		"LAST_NAME": "",
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
	FillVirtualColumns: fillVirtualColumnsViewItems7,
}

func fillVirtualColumnsViewItems7(rowsPre interface{}) interface{} {
	return rowsPre
}
