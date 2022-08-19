package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutLevelType159 struct{}

type LutLevelTypeMainTable159 struct{}

func (l *LutLevelTypeMainTable159) TableName() string {
	return "lut_level_type "
}

//  TableName sets the insert table name for this struct type
func (l *LutLevelType159) TableName() string {
	return "lut_level_type "
}

var LutLevelType159Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Эрсдлийн түвшин",
	Identity:  "id",
	DataTable: "lut_level_type ",
	MainTable: "lut_level_type ",
	DataModel: new(LutLevelType159),
	Data:      new([]LutLevelType159),
	MainModel: new(LutLevelTypeMainTable159),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ersdeliin_level", Label: "Эрсдлийн түвшин"},
	},
	ColumnList: []string{"id", "ersdeliin_level"},
	Filters: map[string]string{
		"id": "",

		"ersdeliin_level": "",
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
	FillVirtualColumns: fillVirtualColumnsLutLevelType159,
}

func fillVirtualColumnsLutLevelType159(rowsPre interface{}) interface{} {
	return rowsPre
}
