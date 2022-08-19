package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type GisCategory275 struct {
	Active    *int       `gorm:"column:active" json:"active"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      *string    `gorm:"column:name" json:"name"`
	Show      *int       `gorm:"column:show" json:"show"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type GisCategoryMainTable275 struct {
	Active     *int       `gorm:"column:active" json:"active"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"created_at"`
	Icon       *string    `gorm:"column:icon" json:"icon"`
	ID         int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LayerOrder *int       `gorm:"column:layer_order" json:"layer_order"`
	MenuOrder  *int       `gorm:"column:menu_order" json:"menu_order"`
	Name       *string    `gorm:"column:name" json:"name"`
	Show       *int       `gorm:"column:show" json:"show"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (g *GisCategoryMainTable275) TableName() string {
	return "gis_category"
}

//  TableName sets the insert table name for this struct type
func (g *GisCategory275) TableName() string {
	return "gis_category"
}

var GisCategory275Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Давхарга бүлэг",
	Identity:  "id",
	DataTable: "gis_category",
	MainTable: "gis_category",
	DataModel: new(GisCategory275),
	Data:      new([]GisCategory275),
	MainModel: new(GisCategoryMainTable275),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "name", Label: "Нэр"},
		datagrid.Column{Model: "active", Label: ""},
		datagrid.Column{Model: "show", Label: ""},
	},
	ColumnList: []string{"id", "name", "active", "show"},
	Filters: map[string]string{
		"id": "",

		"name": "",

		"icon": "",

		"menu_order": "",

		"layer_order": "",

		"created_at": "",

		"updated_at": "",

		"active": "",

		"show": "",
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
	FillVirtualColumns: fillVirtualColumnsGisCategory275,
}

func fillVirtualColumnsGisCategory275(rowsPre interface{}) interface{} {
	return rowsPre
}
