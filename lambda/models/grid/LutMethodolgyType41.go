package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMethodolgyType41 struct {
	Argachlal *string `gorm:"column:argachlal" json:"argachlal"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutMethodolgyTypeMainTable41 struct {
	Argachlal *string `gorm:"column:argachlal" json:"argachlal"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutMethodolgyTypeMainTable41) TableName() string {
	return "lut_methodolgy_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutMethodolgyType41) TableName() string {
	return "lut_methodolgy_type"
}

var LutMethodolgyType41Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Аргачлал",
	Identity:  "id",
	DataTable: "lut_methodolgy_type",
	MainTable: "lut_methodolgy_type",
	DataModel: new(LutMethodolgyType41),
	Data:      new([]LutMethodolgyType41),
	MainModel: new(LutMethodolgyTypeMainTable41),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "argachlal", Label: "Аргачлал"},
	},
	ColumnList: []string{"id", "argachlal"},
	Filters: map[string]string{
		"id": "",

		"argachlal": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMethodolgyType41,
}

func fillVirtualColumnsLutMethodolgyType41(rowsPre interface{}) interface{} {
	return rowsPre
}
