package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type HamtragchBaiguullaga57 struct {
	BNer      *string    `gorm:"column:b_ner" json:"b_ner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Link      *string    `gorm:"column:link" json:"link"`
	Logo      *string    `gorm:"column:logo" json:"logo"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type HamtragchBaiguullagaMainTable57 struct {
	BNer      *string    `gorm:"column:b_ner" json:"b_ner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Link      *string    `gorm:"column:link" json:"link"`
	Logo      *string    `gorm:"column:logo" json:"logo"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (h *HamtragchBaiguullagaMainTable57) TableName() string {
	return "hamtragch_baiguullaga"
}

//  TableName sets the insert table name for this struct type
func (h *HamtragchBaiguullaga57) TableName() string {
	return "hamtragch_baiguullaga"
}

var HamtragchBaiguullaga57Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Хамтрагч байгууллагууд",
	Identity:  "id",
	DataTable: "hamtragch_baiguullaga",
	MainTable: "hamtragch_baiguullaga",
	DataModel: new(HamtragchBaiguullaga57),
	Data:      new([]HamtragchBaiguullaga57),
	MainModel: new(HamtragchBaiguullagaMainTable57),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "b_ner", Label: "Хамтрагч байгууллагын нэр"},
		datagrid.Column{Model: "logo", Label: "Лого"},
		datagrid.Column{Model: "link", Label: "Линк"},
	},
	ColumnList: []string{"id", "b_ner", "logo", "link"},
	Filters: map[string]string{
		"id": "",

		"b_ner": "",

		"logo": "",

		"link": "",

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
	FillVirtualColumns: fillVirtualColumnsHamtragchBaiguullaga57,
}

func fillVirtualColumnsHamtragchBaiguullaga57(rowsPre interface{}) interface{} {
	return rowsPre
}
