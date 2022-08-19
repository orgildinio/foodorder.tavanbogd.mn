package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewMergejilZovlol241 struct {
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	HuraldaanMn *string    `gorm:"column:huraldaan_mn" json:"huraldaan_mn"`
	ID          *int       `gorm:"column:id" json:"id"`
	Tailbar     *string    `gorm:"column:tailbar" json:"tailbar"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type MergejliinZowlolMainTable241 struct {
	BureldehuunMn *string    `gorm:"column:bureldehuun_mn" json:"bureldehuun_mn"`
	Conference    *string    `gorm:"column:conference" json:"conference"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	Delegation    *string    `gorm:"column:delegation" json:"delegation"`
	DeletedAt     *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Huraldaan     *string    `gorm:"column:huraldaan" json:"huraldaan"`
	ID            int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (m *MergejliinZowlolMainTable241) TableName() string {
	return "mergejliin_zowlol"
}

//  TableName sets the insert table name for this struct type
func (v *ViewMergejilZovlol241) TableName() string {
	return "view_mergejil_zovlol"
}

var ViewMergejilZovlol241Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Мэргэжлийн зөвлөл",
	Identity:  "id",
	DataTable: "view_mergejil_zovlol",
	MainTable: "mergejliin_zowlol",
	DataModel: new(ViewMergejilZovlol241),
	Data:      new([]ViewMergejilZovlol241),
	MainModel: new(MergejliinZowlolMainTable241),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "tailbar", Label: "Тайлбар"},
		datagrid.Column{Model: "huraldaan_mn", Label: "Хуралдаан"},
	},
	ColumnList: []string{"id", "tailbar", "huraldaan_mn"},
	Filters: map[string]string{
		"id": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"tailbar": "",

		"tailbar_mn": "",

		"huraldaan_mn": "",
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
	FillVirtualColumns: fillVirtualColumnsViewMergejilZovlol241,
}

func fillVirtualColumnsViewMergejilZovlol241(rowsPre interface{}) interface{} {
	return rowsPre
}
