package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewBanner62 struct {
	Banner    *string    `gorm:"column:banner" json:"banner"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Garchig   *string    `gorm:"column:garchig" json:"garchig"`
	ID        *int       `gorm:"column:id" json:"id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type BannerMainTable62 struct {
	Banner       *string    `gorm:"column:banner" json:"banner"`
	BannerTypeID *int       `gorm:"column:banner_type_id" json:"banner_type_id"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID           int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsID       *int       `gorm:"column:news_id" json:"news_id"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (b *BannerMainTable62) TableName() string {
	return "banner"
}

//  TableName sets the insert table name for this struct type
func (v *ViewBanner62) TableName() string {
	return "view_banner"
}

var ViewBanner62Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Баннер",
	Identity:  "id",
	DataTable: "view_banner",
	MainTable: "banner",
	DataModel: new(ViewBanner62),
	Data:      new([]ViewBanner62),
	MainModel: new(BannerMainTable62),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "banner", Label: "Баннер"},
		datagrid.Column{Model: "garchig", Label: "Гарчиг"},
	},
	ColumnList: []string{"id", "banner", "garchig"},
	Filters: map[string]string{
		"id": "",

		"banner_type_id": "",

		"banner": "",

		"garchig": "",

		"news_id": "",

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
	FillVirtualColumns: fillVirtualColumnsViewBanner62,
}

func fillVirtualColumnsViewBanner62(rowsPre interface{}) interface{} {
	return rowsPre
}
