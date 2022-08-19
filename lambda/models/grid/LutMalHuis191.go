package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalHuis191 struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalHuis *string `gorm:"column:mal_huis" json:"mal_huis"`
}

type LutMalHuisMainTable191 struct {
	ID      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalHuis *string `gorm:"column:mal_huis" json:"mal_huis"`
}

func (l *LutMalHuisMainTable191) TableName() string {
	return "lut_mal_huis"
}

//  TableName sets the insert table name for this struct type
func (l *LutMalHuis191) TableName() string {
	return "lut_mal_huis"
}

var LutMalHuis191Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Малын хүйс",
	Identity:  "id",
	DataTable: "lut_mal_huis",
	MainTable: "lut_mal_huis",
	DataModel: new(LutMalHuis191),
	Data:      new([]LutMalHuis191),
	MainModel: new(LutMalHuisMainTable191),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "mal_huis", Label: "Малын хүйс"},
	},
	ColumnList: []string{"id", "mal_huis"},
	Filters: map[string]string{
		"id": "",

		"mal_huis": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMalHuis191,
}

func fillVirtualColumnsLutMalHuis191(rowsPre interface{}) interface{} {
	return rowsPre
}
