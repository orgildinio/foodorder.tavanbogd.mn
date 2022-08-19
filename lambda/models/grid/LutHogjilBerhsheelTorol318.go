package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutHogjilBerhsheelTorol318 struct {
	HogjilBerhsheelTorol *string `gorm:"column:hogjil_berhsheel_torol" json:"hogjil_berhsheel_torol"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutHogjilBerhsheelTorolMainTable318 struct {
	HogjilBerhsheelTorol *string `gorm:"column:hogjil_berhsheel_torol" json:"hogjil_berhsheel_torol"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutHogjilBerhsheelTorolMainTable318) TableName() string {
	return "lut_hogjil_berhsheel_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutHogjilBerhsheelTorol318) TableName() string {
	return "lut_hogjil_berhsheel_torol"
}

var LutHogjilBerhsheelTorol318Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Хөгжлийин бэрхшээлийн төрөл",
	Identity:  "id",
	DataTable: "lut_hogjil_berhsheel_torol",
	MainTable: "lut_hogjil_berhsheel_torol",
	DataModel: new(LutHogjilBerhsheelTorol318),
	Data:      new([]LutHogjilBerhsheelTorol318),
	MainModel: new(LutHogjilBerhsheelTorolMainTable318),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "hogjil_berhsheel_torol", Label: "Хөгжлийин бэрхшээлийн төрөл"},
	},
	ColumnList: []string{"id", "hogjil_berhsheel_torol"},
	Filters: map[string]string{
		"id": "",

		"hogjil_berhsheel_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutHogjilBerhsheelTorol318,
}

func fillVirtualColumnsLutHogjilBerhsheelTorol318(rowsPre interface{}) interface{} {
	return rowsPre
}
