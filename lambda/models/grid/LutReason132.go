package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutReason132 struct {
	ID     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Reason *string `gorm:"column:reason" json:"reason"`
}

type LutReasonMainTable132 struct {
	ID     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Reason *string `gorm:"column:reason" json:"reason"`
}

func (l *LutReasonMainTable132) TableName() string {
	return "lut_reason"
}

//  TableName sets the insert table name for this struct type
func (l *LutReason132) TableName() string {
	return "lut_reason"
}

var LutReason132Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Халдварласан үзүүлэлт",
	Identity:  "id",
	DataTable: "lut_reason",
	MainTable: "lut_reason",
	DataModel: new(LutReason132),
	Data:      new([]LutReason132),
	MainModel: new(LutReasonMainTable132),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "reason", Label: "Халдварласан үзүүлэлт"},
	},
	ColumnList: []string{"id", "reason"},
	Filters: map[string]string{
		"id": "",

		"reason": "",
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
	FillVirtualColumns: fillVirtualColumnsLutReason132,
}

func fillVirtualColumnsLutReason132(rowsPre interface{}) interface{} {
	return rowsPre
}
