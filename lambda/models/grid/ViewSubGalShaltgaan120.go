package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubGalShaltgaan120 struct {
	ID           *int    `gorm:"column:id" json:"id"`
	ShaltgaanGal *string `gorm:"column:shaltgaan_gal" json:"shaltgaan_gal"`
	SubReasonGal *string `gorm:"column:sub_reason_gal" json:"sub_reason_gal"`
}

type LutSubReasonGalMainTable120 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanGalID *int    `gorm:"column:shaltgaan_gal_id" json:"shaltgaan_gal_id"`
	SubReasonGal   *string `gorm:"column:sub_reason_gal" json:"sub_reason_gal"`
}

func (l *LutSubReasonGalMainTable120) TableName() string {
	return "lut_sub_reason_gal"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubGalShaltgaan120) TableName() string {
	return "view_sub_gal_shaltgaan"
}

var ViewSubGalShaltgaan120Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Галын Үзүүлэлт",
	Identity:  "id",
	DataTable: "view_sub_gal_shaltgaan",
	MainTable: "lut_sub_reason_gal",
	DataModel: new(ViewSubGalShaltgaan120),
	Data:      new([]ViewSubGalShaltgaan120),
	MainModel: new(LutSubReasonGalMainTable120),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "shaltgaan_gal", Label: "Галын төрөл"},
		datagrid.Column{Model: "sub_reason_gal", Label: "Галын дэд төрөл"},
	},
	ColumnList: []string{"id", "shaltgaan_gal", "sub_reason_gal"},
	Filters: map[string]string{
		"id": "",

		"shaltgaan_gal_id": "",

		"shaltgaan_gal": "",

		"sub_reason_gal": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubGalShaltgaan120,
}

func fillVirtualColumnsViewSubGalShaltgaan120(rowsPre interface{}) interface{} {
	return rowsPre
}
