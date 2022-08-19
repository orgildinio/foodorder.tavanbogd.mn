package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubReason131 struct {
	ID        *int    `gorm:"column:id" json:"id"`
	Reason    *string `gorm:"column:reason" json:"reason"`
	ReasonID  *int    `gorm:"column:reason_id" json:"reason_id"`
	SubReason *string `gorm:"column:sub_reason" json:"sub_reason"`
}

type LutSubReasonMainTable131 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ReasonID  *int    `gorm:"column:reason_id" json:"reason_id"`
	SubReason *string `gorm:"column:sub_reason" json:"sub_reason"`
}

func (l *LutSubReasonMainTable131) TableName() string {
	return "lut_sub_reason"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubReason131) TableName() string {
	return "view_sub_reason"
}

var ViewSubReason131Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Халдварласан дэд үзүүлэлт",
	Identity:  "id",
	DataTable: "view_sub_reason",
	MainTable: "lut_sub_reason",
	DataModel: new(ViewSubReason131),
	Data:      new([]ViewSubReason131),
	MainModel: new(LutSubReasonMainTable131),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "reason_id", Label: ""},
		datagrid.Column{Model: "reason", Label: "Халдварласан үзүүлэлт"},
		datagrid.Column{Model: "sub_reason", Label: "Халдварласан дэд үзүүлэлт"},
	},
	ColumnList: []string{"id", "reason_id", "reason", "sub_reason"},
	Filters: map[string]string{
		"id": "",

		"reason_id": "",

		"reason": "",

		"sub_reason": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubReason131,
}

func fillVirtualColumnsViewSubReason131(rowsPre interface{}) interface{} {
	return rowsPre
}
