package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type HolbooBarih59 struct {
	BNer      *string    `gorm:"column:b_ner" json:"b_ner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email     *string    `gorm:"column:email" json:"email"`
	Hayag     *string    `gorm:"column:hayag" json:"hayag"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas      *string    `gorm:"column:utas" json:"utas"`
	Web       *string    `gorm:"column:web" json:"web"`
}

type HolbooBarihMainTable59 struct {
	BNer      *string    `gorm:"column:b_ner" json:"b_ner"`
	Bairshil  *string    `gorm:"column:bairshil" json:"bairshil"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email     *string    `gorm:"column:email" json:"email"`
	Hayag     *string    `gorm:"column:hayag" json:"hayag"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas      *string    `gorm:"column:utas" json:"utas"`
	Web       *string    `gorm:"column:web" json:"web"`
}

func (h *HolbooBarihMainTable59) TableName() string {
	return "holboo_barih"
}

//  TableName sets the insert table name for this struct type
func (h *HolbooBarih59) TableName() string {
	return "holboo_barih"
}

var HolbooBarih59Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Холбоо барих",
	Identity:  "id",
	DataTable: "holboo_barih",
	MainTable: "holboo_barih",
	DataModel: new(HolbooBarih59),
	Data:      new([]HolbooBarih59),
	MainModel: new(HolbooBarihMainTable59),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "b_ner", Label: "Байгууллагын нэр"},
		datagrid.Column{Model: "hayag", Label: "Хаяг"},
		datagrid.Column{Model: "web", Label: "Цахим хаяг"},
		datagrid.Column{Model: "email", Label: "Цахим шуудан"},
		datagrid.Column{Model: "utas", Label: "Утасны дугаар"},
	},
	ColumnList: []string{"id", "b_ner", "hayag", "web", "email", "utas"},
	Filters: map[string]string{
		"id": "",

		"b_ner": "",

		"hayag": "",

		"bairshil": "",

		"web": "",

		"email": "",

		"utas": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",
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
	FillVirtualColumns: fillVirtualColumnsHolbooBarih59,
}

func fillVirtualColumnsHolbooBarih59(rowsPre interface{}) interface{} {
	return rowsPre
}
