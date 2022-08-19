package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutConfirmStatus14 struct {
	ConfirmStatus *string `gorm:"column:confirm_status" json:"confirm_status"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutConfirmStatusMainTable14 struct {
	ConfirmStatus *string `gorm:"column:confirm_status" json:"confirm_status"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutConfirmStatusMainTable14) TableName() string {
	return "lut_confirm_status"
}

//  TableName sets the insert table name for this struct type
func (l *LutConfirmStatus14) TableName() string {
	return "lut_confirm_status"
}

var LutConfirmStatus14Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Баталгаажуулалтын төлөв",
	Identity:  "id",
	DataTable: "lut_confirm_status",
	MainTable: "lut_confirm_status",
	DataModel: new(LutConfirmStatus14),
	Data:      new([]LutConfirmStatus14),
	MainModel: new(LutConfirmStatusMainTable14),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "confirm_status", Label: "Баталгаажуулалтын төлөв"},
	},
	ColumnList: []string{"id", "confirm_status"},
	Filters: map[string]string{
		"id": "",

		"confirm_status": "",
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
	FillVirtualColumns: fillVirtualColumnsLutConfirmStatus14,
}

func fillVirtualColumnsLutConfirmStatus14(rowsPre interface{}) interface{} {
	return rowsPre
}
