package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAnimals126 struct {
	Animals *string `gorm:"column:animals" json:"animals"`
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutAnimalsMainTable126 struct {
	Animals *string `gorm:"column:animals" json:"animals"`
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutAnimalsMainTable126) TableName() string {
	return "lut_animals"
}

//  TableName sets the insert table name for this struct type
func (l *LutAnimals126) TableName() string {
	return "lut_animals"
}

var LutAnimals126Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Мал амьтны төрөл",
	Identity:  "id",
	DataTable: "lut_animals",
	MainTable: "lut_animals",
	DataModel: new(LutAnimals126),
	Data:      new([]LutAnimals126),
	MainModel: new(LutAnimalsMainTable126),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "animals", Label: "Мал амьтны төрөл"},
	},
	ColumnList: []string{"id", "animals"},
	Filters: map[string]string{
		"id": "",

		"animals": "",
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
	FillVirtualColumns: fillVirtualColumnsLutAnimals126,
}

func fillVirtualColumnsLutAnimals126(rowsPre interface{}) interface{} {
	return rowsPre
}
