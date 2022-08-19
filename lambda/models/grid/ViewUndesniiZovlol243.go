package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewUndesniiZovlol243 struct {
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Huraldaan            *string    `gorm:"column:huraldaan" json:"huraldaan"`
	ID                   *int       `gorm:"column:id" json:"id"`
	TailbarMnBureldehuun *string    `gorm:"column:tailbar_mn_bureldehuun" json:"tailbar_mn_bureldehuun"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type UndesniiZowlolMainTable243 struct {
	ButetsBureldhuun *string    `gorm:"column:butets_bureldhuun" json:"butets_bureldhuun"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID               int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *UndesniiZowlolMainTable243) TableName() string {
	return "undesnii_zowlol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewUndesniiZovlol243) TableName() string {
	return "view_undesnii_zovlol"
}

var ViewUndesniiZovlol243Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Үндэсний зөвлөл",
	Identity:  "id",
	DataTable: "view_undesnii_zovlol",
	MainTable: "undesnii_zowlol",
	DataModel: new(ViewUndesniiZovlol243),
	Data:      new([]ViewUndesniiZovlol243),
	MainModel: new(UndesniiZowlolMainTable243),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "huraldaan", Label: "Хуралдаан"},
		datagrid.Column{Model: "tailbar_mn_bureldehuun", Label: "Тайлбар"},
	},
	ColumnList: []string{"id", "huraldaan", "tailbar_mn_bureldehuun"},
	Filters: map[string]string{
		"id": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"huraldaan": "",

		"tailbar_mn_bureldehuun": "",
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
	FillVirtualColumns: fillVirtualColumnsViewUndesniiZovlol243,
}

func fillVirtualColumnsViewUndesniiZovlol243(rowsPre interface{}) interface{} {
	return rowsPre
}
