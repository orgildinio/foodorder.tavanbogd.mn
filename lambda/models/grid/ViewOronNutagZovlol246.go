package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewOronNutagZovlol246 struct {
	Aimagname *string    `gorm:"column:aimagname" json:"aimagname"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        *int       `gorm:"column:id" json:"id"`
	Sumname   *string    `gorm:"column:sumname" json:"sumname"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type OronNutgiinZowlolMainTable246 struct {
	AimagID   *int       `gorm:"column:aimag_id" json:"aimag_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SumID     *int       `gorm:"column:sum_id" json:"sum_id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (o *OronNutgiinZowlolMainTable246) TableName() string {
	return "oron_nutgiin_zowlol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewOronNutagZovlol246) TableName() string {
	return "view_oron_nutag_zovlol"
}

var ViewOronNutagZovlol246Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Орон нутгийн зөвлөл",
	Identity:  "id",
	DataTable: "view_oron_nutag_zovlol",
	MainTable: "oron_nutgiin_zowlol",
	DataModel: new(ViewOronNutagZovlol246),
	Data:      new([]ViewOronNutagZovlol246),
	MainModel: new(OronNutgiinZowlolMainTable246),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "aimagname", Label: "Аймаг / Нийслэл"},
		datagrid.Column{Model: "sumname", Label: "Дүүрэг"},
	},
	ColumnList: []string{"id", "aimagname", "sumname"},
	Filters: map[string]string{
		"id": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"aimag_id": "",

		"sum_id": "",

		"aimagname": "",

		"sumname": "",
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
	FillVirtualColumns: fillVirtualColumnsViewOronNutagZovlol246,
}

func fillVirtualColumnsViewOronNutagZovlol246(rowsPre interface{}) interface{} {
	return rowsPre
}
