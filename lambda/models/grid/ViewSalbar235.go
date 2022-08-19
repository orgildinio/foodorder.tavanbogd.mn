package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSalbar235 struct {
	Angi   *string `gorm:"column:angi" json:"angi"`
	ID     *int    `gorm:"column:id" json:"id"`
	Salbar *string `gorm:"column:salbar" json:"salbar"`
}

type ViewSalbarMainTable235 struct {
	Angi   *string `gorm:"column:angi" json:"angi"`
	AngiID *int    `gorm:"column:angi_id" json:"angi_id"`
	ID     *int    `gorm:"column:id" json:"id"`
	Salbar *string `gorm:"column:salbar" json:"salbar"`
}

func (v *ViewSalbarMainTable235) TableName() string {
	return "view_salbar"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSalbar235) TableName() string {
	return "view_salbar"
}

var ViewSalbar235Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Салбар ",
	Identity:  "id",
	DataTable: "view_salbar",
	MainTable: "view_salbar",
	DataModel: new(ViewSalbar235),
	Data:      new([]ViewSalbar235),
	MainModel: new(ViewSalbarMainTable235),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "angi", Label: "Анги"},
		datagrid.Column{Model: "salbar", Label: "Салбар "},
	},
	ColumnList: []string{"id", "angi", "salbar"},
	Filters: map[string]string{
		"id": "",

		"angi_id": "",

		"angi": "",

		"salbar": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSalbar235,
}

func fillVirtualColumnsViewSalbar235(rowsPre interface{}) interface{} {
	return rowsPre
}
