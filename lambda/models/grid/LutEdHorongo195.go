package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEdHorongo195 struct {
	EdHorongo *string `gorm:"column:ed_horongo" json:"ed_horongo"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutEdHorongoMainTable195 struct {
	EdHorongo *string `gorm:"column:ed_horongo" json:"ed_horongo"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutEdHorongoMainTable195) TableName() string {
	return "lut_ed_horongo"
}

//  TableName sets the insert table name for this struct type
func (l *LutEdHorongo195) TableName() string {
	return "lut_ed_horongo"
}

var LutEdHorongo195Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Эд хөрөнгө",
	Identity:  "id",
	DataTable: "lut_ed_horongo",
	MainTable: "lut_ed_horongo",
	DataModel: new(LutEdHorongo195),
	Data:      new([]LutEdHorongo195),
	MainModel: new(LutEdHorongoMainTable195),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ed_horongo", Label: "Эд хөрөнгө"},
	},
	ColumnList: []string{"id", "ed_horongo"},
	Filters: map[string]string{
		"id": "",

		"ed_horongo": "",
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
	FillVirtualColumns: fillVirtualColumnsLutEdHorongo195,
}

func fillVirtualColumnsLutEdHorongo195(rowsPre interface{}) interface{} {
	return rowsPre
}
