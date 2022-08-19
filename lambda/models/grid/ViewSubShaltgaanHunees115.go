package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubShaltgaanHunees115 struct {
	ID                 *int    `gorm:"column:id" json:"id"`
	LutShaltgaanHunees *string `gorm:"column:lut_shaltgaan_hunees" json:"lut_shaltgaan_hunees"`
	SubReasonHunees    *string `gorm:"column:sub_reason_hunees" json:"sub_reason_hunees"`
}

type LutSubReasonHuneesMainTable115 struct {
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanHuneesID *int    `gorm:"column:shaltgaan_hunees_id" json:"shaltgaan_hunees_id"`
	SubReasonHunees   *string `gorm:"column:sub_reason_hunees" json:"sub_reason_hunees"`
}

func (l *LutSubReasonHuneesMainTable115) TableName() string {
	return "lut_sub_reason_hunees"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubShaltgaanHunees115) TableName() string {
	return "view_sub_shaltgaan_hunees"
}

var ViewSubShaltgaanHunees115Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Дэд төрөл /Хүнээс/",
	Identity:  "id",
	DataTable: "view_sub_shaltgaan_hunees",
	MainTable: "lut_sub_reason_hunees",
	DataModel: new(ViewSubShaltgaanHunees115),
	Data:      new([]ViewSubShaltgaanHunees115),
	MainModel: new(LutSubReasonHuneesMainTable115),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "lut_shaltgaan_hunees", Label: "Төрөл"},
		datagrid.Column{Model: "sub_reason_hunees", Label: "Дэд төрөл"},
	},
	ColumnList: []string{"id", "lut_shaltgaan_hunees", "sub_reason_hunees"},
	Filters: map[string]string{
		"id": "",

		"shaltgaan_hunees_id": "",

		"lut_shaltgaan_hunees": "",

		"sub_reason_hunees": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubShaltgaanHunees115,
}

func fillVirtualColumnsViewSubShaltgaanHunees115(rowsPre interface{}) interface{} {
	return rowsPre
}
