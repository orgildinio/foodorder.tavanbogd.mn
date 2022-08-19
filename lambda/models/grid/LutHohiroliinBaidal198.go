package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHohiroliinBaidal198 struct {
	HohiroliinBaidal *string `gorm:"column:hohiroliin_baidal" json:"hohiroliin_baidal"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutHohiroliinBaidalMainTable198 struct {
	HohiroliinBaidal *string `gorm:"column:hohiroliin_baidal" json:"hohiroliin_baidal"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHohiroliinBaidalMainTable198) TableName() string {
	return "lut_hohiroliin_baidal"
}

//  TableName sets the insert table name for this struct type
func (l *LutHohiroliinBaidal198) TableName() string {
	return "lut_hohiroliin_baidal"
}

var LutHohiroliinBaidal198Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хохирлын байдал",
	Identity:  "id",
	DataTable: "lut_hohiroliin_baidal",
	MainTable: "lut_hohiroliin_baidal",
	DataModel: new(LutHohiroliinBaidal198),
	Data:      new([]LutHohiroliinBaidal198),
	MainModel: new(LutHohiroliinBaidalMainTable198),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "hohiroliin_baidal", Label: "Хохирлын байдал"},
	},
	ColumnList: []string{"id", "hohiroliin_baidal"},
	Filters: map[string]string{
		"id": "",

		"hohiroliin_baidal": "",
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
	FillVirtualColumns: fillVirtualColumnsLutHohiroliinBaidal198,
}

func fillVirtualColumnsLutHohiroliinBaidal198(rowsPre interface{}) interface{} {
	return rowsPre
}
