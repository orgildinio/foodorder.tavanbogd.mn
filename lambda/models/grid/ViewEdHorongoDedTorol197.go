package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewEdHorongoDedTorol197 struct {
	DedEdHorongo *string `gorm:"column:ded_ed_horongo" json:"ded_ed_horongo"`
	EdHorongo    *string `gorm:"column:ed_horongo" json:"ed_horongo"`
	ID           *int    `gorm:"column:id" json:"id"`
}

type LutDedEdHorongoMainTable197 struct {
	DedEdHorongo     *string `gorm:"column:ded_ed_horongo" json:"ded_ed_horongo"`
	EdHorongoTorolID *int    `gorm:"column:ed_horongo_torol_id" json:"ed_horongo_torol_id"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDedEdHorongoMainTable197) TableName() string {
	return "lut_ded_ed_horongo"
}

//  TableName sets the insert table name for this struct type
func (v *ViewEdHorongoDedTorol197) TableName() string {
	return "view_ed_horongo_ded_torol"
}

var ViewEdHorongoDedTorol197Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Эд хөрөнгийн дэд төрөл",
	Identity:  "id",
	DataTable: "view_ed_horongo_ded_torol",
	MainTable: "lut_ded_ed_horongo",
	DataModel: new(ViewEdHorongoDedTorol197),
	Data:      new([]ViewEdHorongoDedTorol197),
	MainModel: new(LutDedEdHorongoMainTable197),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ed_horongo", Label: "Эд хөрөнгийн төрөл"},
		datagrid.Column{Model: "ded_ed_horongo", Label: "Эд хөрөнгийн дэд төрөл"},
	},
	ColumnList: []string{"id", "ed_horongo", "ded_ed_horongo"},
	Filters: map[string]string{
		"id": "",

		"ed_horongo_torol_id": "",

		"ed_horongo": "",

		"ded_ed_horongo": "",
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
	FillVirtualColumns: fillVirtualColumnsViewEdHorongoDedTorol197,
}

func fillVirtualColumnsViewEdHorongoDedTorol197(rowsPre interface{}) interface{} {
	return rowsPre
}
