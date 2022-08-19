package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMaindisasterType85 struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MaindisasterType *string `gorm:"column:maindisaster_type" json:"maindisaster_type"`
}

type LutMaindisasterTypeMainTable85 struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MaindisasterType *string `gorm:"column:maindisaster_type" json:"maindisaster_type"`
}

func (l *LutMaindisasterTypeMainTable85) TableName() string {
	return "lut_maindisaster_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutMaindisasterType85) TableName() string {
	return "lut_maindisaster_type"
}

var LutMaindisasterType85Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Аюулын төрөл",
	Identity:  "id",
	DataTable: "lut_maindisaster_type",
	MainTable: "lut_maindisaster_type",
	DataModel: new(LutMaindisasterType85),
	Data:      new([]LutMaindisasterType85),
	MainModel: new(LutMaindisasterTypeMainTable85),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "maindisaster_type", Label: "Аюулын төрөл"},
	},
	ColumnList: []string{"id", "maindisaster_type"},
	Filters: map[string]string{
		"id": "",

		"maindisaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMaindisasterType85,
}

func fillVirtualColumnsLutMaindisasterType85(rowsPre interface{}) interface{} {
	return rowsPre
}
