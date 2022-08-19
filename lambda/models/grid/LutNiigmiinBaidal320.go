package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNiigmiinBaidal320 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NiigmiinBaidal *string `gorm:"column:niigmiin_baidal" json:"niigmiin_baidal"`
}

type LutNiigmiinBaidalMainTable320 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NiigmiinBaidal *string `gorm:"column:niigmiin_baidal" json:"niigmiin_baidal"`
}

func (l *LutNiigmiinBaidalMainTable320) TableName() string {
	return "lut_niigmiin_baidal"
}

//  TableName sets the insert table name for this struct type
func (l *LutNiigmiinBaidal320) TableName() string {
	return "lut_niigmiin_baidal"
}

var LutNiigmiinBaidal320Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Нийгмийн байдал",
	Identity:  "id",
	DataTable: "lut_niigmiin_baidal",
	MainTable: "lut_niigmiin_baidal",
	DataModel: new(LutNiigmiinBaidal320),
	Data:      new([]LutNiigmiinBaidal320),
	MainModel: new(LutNiigmiinBaidalMainTable320),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "niigmiin_baidal", Label: "Нийгмийн байдал"},
	},
	ColumnList: []string{"id", "niigmiin_baidal"},
	Filters: map[string]string{
		"id": "",

		"niigmiin_baidal": "",
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
	FillVirtualColumns: fillVirtualColumnsLutNiigmiinBaidal320,
}

func fillVirtualColumnsLutNiigmiinBaidal320(rowsPre interface{}) interface{} {
	return rowsPre
}
