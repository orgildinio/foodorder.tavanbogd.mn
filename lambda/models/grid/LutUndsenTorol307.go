package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutUndsenTorol307 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndsenTorol *string `gorm:"column:undsen_torol" json:"undsen_torol"`
}

type LutUndsenTorolMainTable307 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndsenTorol *string `gorm:"column:undsen_torol" json:"undsen_torol"`
}

func (l *LutUndsenTorolMainTable307) TableName() string {
	return "lut_undsen_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutUndsenTorol307) TableName() string {
	return "lut_undsen_torol"
}

var LutUndsenTorol307Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Үндсэн төрөл ",
	Identity:  "id",
	DataTable: "lut_undsen_torol",
	MainTable: "lut_undsen_torol",
	DataModel: new(LutUndsenTorol307),
	Data:      new([]LutUndsenTorol307),
	MainModel: new(LutUndsenTorolMainTable307),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "undsen_torol", Label: "Үндсэн төрөл "},
	},
	ColumnList: []string{"id", "undsen_torol"},
	Filters: map[string]string{
		"id": "",

		"undsen_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutUndsenTorol307,
}

func fillVirtualColumnsLutUndsenTorol307(rowsPre interface{}) interface{} {
	return rowsPre
}
