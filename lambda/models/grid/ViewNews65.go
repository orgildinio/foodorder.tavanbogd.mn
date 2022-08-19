package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewNews65 struct {
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Garchig   *string    `gorm:"column:garchig" json:"garchig"`
	ID        *int       `gorm:"column:id" json:"id"`
	NewsType  *string    `gorm:"column:news_type" json:"news_type"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UrsahEseh *int       `gorm:"column:ursah_eseh" json:"ursah_eseh"`
	Zurag     *string    `gorm:"column:zurag" json:"zurag"`
}

type NewsMainTable65 struct {
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Delgerengui *string    `gorm:"column:delgerengui" json:"delgerengui"`
	Garchig     *string    `gorm:"column:garchig" json:"garchig"`
	Handalt     *int       `gorm:"column:handalt" json:"handalt"`
	ID          int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Logo        *string    `gorm:"column:logo" json:"logo"`
	NewsTypeID  *int       `gorm:"column:news_type_id" json:"news_type_id"`
	Share       *int       `gorm:"column:share" json:"share"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UrsahEseh   *int       `gorm:"column:ursah_eseh" json:"ursah_eseh"`
	Zurag       *string    `gorm:"column:zurag" json:"zurag"`
}

func (n *NewsMainTable65) TableName() string {
	return "news"
}

//  TableName sets the insert table name for this struct type
func (v *ViewNews65) TableName() string {
	return "view_news"
}

var ViewNews65Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Мэдээ мэдээлэл",
	Identity:  "id",
	DataTable: "view_news",
	MainTable: "news",
	DataModel: new(ViewNews65),
	Data:      new([]ViewNews65),
	MainModel: new(NewsMainTable65),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "news_type", Label: "Aнгилал"},
		datagrid.Column{Model: "garchig", Label: "Гарчиг"},
		datagrid.Column{Model: "zurag", Label: "Зураг"},
		datagrid.Column{Model: "ursah_eseh", Label: "Урсах эсэх"},
	},
	ColumnList: []string{"id", "news_type", "garchig", "zurag", "ursah_eseh"},
	Filters: map[string]string{
		"id": "",

		"news_type": "Text",

		"news_type_id": "",

		"garchig": "",

		"delgerengui": "",

		"zurag": "Text",

		"ursah_eseh": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"handalt": "",

		"logo": "",

		"share": "",
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
	FillVirtualColumns: fillVirtualColumnsViewNews65,
}

func fillVirtualColumnsViewNews65(rowsPre interface{}) interface{} {
	return rowsPre
}
