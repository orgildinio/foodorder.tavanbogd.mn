package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubAnimals127 struct {
	Animals    *string `gorm:"column:animals" json:"animals"`
	ID         *int    `gorm:"column:id" json:"id"`
	SubAnimals *string `gorm:"column:sub_animals" json:"sub_animals"`
}

type LutSubAnimalsMainTable127 struct {
	AnimalsID  *int    `gorm:"column:animals_id" json:"animals_id"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubAnimals *string `gorm:"column:sub_animals" json:"sub_animals"`
}

func (l *LutSubAnimalsMainTable127) TableName() string {
	return "lut_sub_animals"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubAnimals127) TableName() string {
	return "view_sub_animals"
}

var ViewSubAnimals127Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Мал амьтны дэд.төрөл",
	Identity:  "id",
	DataTable: "view_sub_animals",
	MainTable: "lut_sub_animals",
	DataModel: new(ViewSubAnimals127),
	Data:      new([]ViewSubAnimals127),
	MainModel: new(LutSubAnimalsMainTable127),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "animals", Label: "Мал амьтны.төрөл"},
		datagrid.Column{Model: "sub_animals", Label: "Мал амьтны дэд.төрөл"},
	},
	ColumnList: []string{"id", "animals", "sub_animals"},
	Filters: map[string]string{
		"id": "",

		"animals_id": "",

		"animals": "",

		"sub_animals": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubAnimals127,
}

func fillVirtualColumnsViewSubAnimals127(rowsPre interface{}) interface{} {
	return rowsPre
}
