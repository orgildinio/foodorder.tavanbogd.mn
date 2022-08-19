package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDedDedEdHorongo218 struct {
	DedDedHorongo *string `gorm:"column:ded_ded_horongo" json:"ded_ded_horongo"`
	DedEdHorongo  *string `gorm:"column:ded_ed_horongo" json:"ded_ed_horongo"`
	ID            *int    `gorm:"column:id" json:"id"`
}

type LutDedDedEdHorongoMainTable218 struct {
	DedDedHorongo  *string `gorm:"column:ded_ded_horongo" json:"ded_ded_horongo"`
	DedEdHorongoID *int    `gorm:"column:ded_ed_horongo_id" json:"ded_ed_horongo_id"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDedDedEdHorongoMainTable218) TableName() string {
	return "lut_ded_ded_ed_horongo"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDedDedEdHorongo218) TableName() string {
	return "view_ded_ded_ed_horongo"
}

var ViewDedDedEdHorongo218Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Эд хөрөнгийн дэд дэд төрөл",
	Identity:  "id",
	DataTable: "view_ded_ded_ed_horongo",
	MainTable: "lut_ded_ded_ed_horongo",
	DataModel: new(ViewDedDedEdHorongo218),
	Data:      new([]ViewDedDedEdHorongo218),
	MainModel: new(LutDedDedEdHorongoMainTable218),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ded_ed_horongo", Label: "Эд хөрөнгийн дэд төрөл"},
		datagrid.Column{Model: "ded_ded_horongo", Label: "Эд хөрөнгийн дэд дэд төрөл"},
	},
	ColumnList: []string{"id", "ded_ed_horongo", "ded_ded_horongo"},
	Filters: map[string]string{
		"id": "",

		"ded_ed_horongo_id": "",

		"ded_ed_horongo": "",

		"ded_ded_horongo": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDedDedEdHorongo218,
}

func fillVirtualColumnsViewDedDedEdHorongo218(rowsPre interface{}) interface{} {
	return rowsPre
}
