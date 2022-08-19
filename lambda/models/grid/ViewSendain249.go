package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSendain249 struct {
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        *int       `gorm:"column:id" json:"id"`
	UaHuree   *string    `gorm:"column:ua_huree" json:"ua_huree"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SendainUaMainTable249 struct {
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UaHuree   *string    `gorm:"column:ua_huree" json:"ua_huree"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (s *SendainUaMainTable249) TableName() string {
	return "sendain_ua"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSendain249) TableName() string {
	return "view_sendain"
}

var ViewSendain249Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Сендайн үйл ажиллагаа",
	Identity:  "id",
	DataTable: "view_sendain",
	MainTable: "sendain_ua",
	DataModel: new(ViewSendain249),
	Data:      new([]ViewSendain249),
	MainModel: new(SendainUaMainTable249),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ua_huree", Label: "Үйл ажиллагааны хүрээ"},
		datagrid.Column{Model: "created_at", Label: "Огноо"},
	},
	ColumnList: []string{"id", "ua_huree", "created_at"},
	Filters: map[string]string{
		"id": "",

		"ua_huree": "",

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
	FillVirtualColumns: fillVirtualColumnsViewSendain249,
}

func fillVirtualColumnsViewSendain249(rowsPre interface{}) interface{} {
	return rowsPre
}
