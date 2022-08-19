package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutLevelType161 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LevelType *string `gorm:"column:level_type" json:"level_type"`
}

type LutLevelTypeMainTable161 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LevelType *string `gorm:"column:level_type" json:"level_type"`
}

func (l *LutLevelTypeMainTable161) TableName() string {
	return "lut_level_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutLevelType161) TableName() string {
	return "lut_level_type"
}

var LutLevelType161Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Эрсдлийн түвшин",
	Identity:  "id",
	DataTable: "lut_level_type",
	MainTable: "lut_level_type",
	DataModel: new(LutLevelType161),
	Data:      new([]LutLevelType161),
	MainModel: new(LutLevelTypeMainTable161),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "level_type", Label: "Эрсдлийн түвшин"},
	},
	ColumnList: []string{"id", "level_type"},
	Filters: map[string]string{
		"id": "",

		"level_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutLevelType161,
}

func fillVirtualColumnsLutLevelType161(rowsPre interface{}) interface{} {
	return rowsPre
}
