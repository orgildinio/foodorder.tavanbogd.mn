package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNDisasterReason99 struct {
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NDisasterReason *string `gorm:"column:n_disaster_reason" json:"n_disaster_reason"`
}

type LutNDisasterReasonMainTable99 struct {
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NDisasterReason *string `gorm:"column:n_disaster_reason" json:"n_disaster_reason"`
}

func (l *LutNDisasterReasonMainTable99) TableName() string {
	return "lut_n_disaster_reason"
}

//  TableName sets the insert table name for this struct type
func (l *LutNDisasterReason99) TableName() string {
	return "lut_n_disaster_reason"
}

var LutNDisasterReason99Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Аюулт үзэгдлийн шалтгаан",
	Identity:  "id",
	DataTable: "lut_n_disaster_reason",
	MainTable: "lut_n_disaster_reason",
	DataModel: new(LutNDisasterReason99),
	Data:      new([]LutNDisasterReason99),
	MainModel: new(LutNDisasterReasonMainTable99),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "n_disaster_reason", Label: "Аюулт үзэгдлийн шалтгаан"},
	},
	ColumnList: []string{"id", "n_disaster_reason"},
	Filters: map[string]string{
		"id": "",

		"n_disaster_reason": "",
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
	FillVirtualColumns: fillVirtualColumnsLutNDisasterReason99,
}

func fillVirtualColumnsLutNDisasterReason99(rowsPre interface{}) interface{} {
	return rowsPre
}
