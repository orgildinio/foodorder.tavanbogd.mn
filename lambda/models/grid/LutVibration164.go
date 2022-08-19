package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutVibration164 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Vibration *string `gorm:"column:vibration" json:"vibration"`
}

type LutVibrationMainTable164 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Vibration *string `gorm:"column:vibration" json:"vibration"`
}

func (l *LutVibrationMainTable164) TableName() string {
	return "lut_vibration"
}

//  TableName sets the insert table name for this struct type
func (l *LutVibration164) TableName() string {
	return "lut_vibration"
}

var LutVibration164Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Чичирхийллийн хэмжээ",
	Identity:  "id",
	DataTable: "lut_vibration",
	MainTable: "lut_vibration",
	DataModel: new(LutVibration164),
	Data:      new([]LutVibration164),
	MainModel: new(LutVibrationMainTable164),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "vibration", Label: "Чичирхийллийн хэмжээ"},
	},
	ColumnList: []string{"id", "vibration"},
	Filters: map[string]string{
		"id": "",

		"vibration": "",
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
	FillVirtualColumns: fillVirtualColumnsLutVibration164,
}

func fillVirtualColumnsLutVibration164(rowsPre interface{}) interface{} {
	return rowsPre
}
