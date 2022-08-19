package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutTsol236 struct {
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Tsol *string `gorm:"column:tsol" json:"tsol"`
}

type LutTsolMainTable236 struct {
	ID   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Tsol *string `gorm:"column:tsol" json:"tsol"`
}

func (l *LutTsolMainTable236) TableName() string {
	return "lut_tsol"
}

//  TableName sets the insert table name for this struct type
func (l *LutTsol236) TableName() string {
	return "lut_tsol"
}

var LutTsol236Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Цол",
	Identity:  "id",
	DataTable: "lut_tsol",
	MainTable: "lut_tsol",
	DataModel: new(LutTsol236),
	Data:      new([]LutTsol236),
	MainModel: new(LutTsolMainTable236),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "tsol", Label: ""},
	},
	ColumnList: []string{"id", "tsol"},
	Filters: map[string]string{
		"id": "",

		"tsol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutTsol236,
}

func fillVirtualColumnsLutTsol236(rowsPre interface{}) interface{} {
	return rowsPre
}
