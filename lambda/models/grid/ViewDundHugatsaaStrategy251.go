package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewDundHugatsaaStrategy251 struct {
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID             *int       `gorm:"column:id" json:"id"`
	MongolStrategy *string    `gorm:"column:mongol_strategy" json:"mongol_strategy"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type DundHugtsaaStrategyMainTable251 struct {
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID             int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MongolStrategy *string    `gorm:"column:mongol_strategy" json:"mongol_strategy"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (d *DundHugtsaaStrategyMainTable251) TableName() string {
	return "dund_hugtsaa_strategy"
}

//  TableName sets the insert table name for this struct type
func (v *ViewDundHugatsaaStrategy251) TableName() string {
	return "view_dund_hugatsaa_strategy"
}

var ViewDundHugatsaaStrategy251Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Дунд хугацааны стратеги",
	Identity:  "id",
	DataTable: "view_dund_hugatsaa_strategy",
	MainTable: "dund_hugtsaa_strategy",
	DataModel: new(ViewDundHugatsaaStrategy251),
	Data:      new([]ViewDundHugatsaaStrategy251),
	MainModel: new(DundHugtsaaStrategyMainTable251),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "mongol_strategy", Label: "Монгол улсад хэрэгжүүлэх дунд хугацааны стратеги"},
	},
	ColumnList: []string{"id", "mongol_strategy"},
	Filters: map[string]string{
		"id": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"mongol_strategy": "",
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
	FillVirtualColumns: fillVirtualColumnsViewDundHugatsaaStrategy251,
}

func fillVirtualColumnsViewDundHugatsaaStrategy251(rowsPre interface{}) interface{} {
	return rowsPre
}
