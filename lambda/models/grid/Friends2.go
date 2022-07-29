package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type Friends2 struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Phone string `gorm:"column:phone" json:"phone"`
	Zurag string `gorm:"column:zurag" json:"zurag"`
}

type FriendsMainTable2 struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Phone string `gorm:"column:phone" json:"phone"`
	Zurag string `gorm:"column:zurag" json:"zurag"`
}

func (f *FriendsMainTable2) TableName() string {
	return "friends"
}

//  TableName sets the insert table name for this struct type
func (f *Friends2) TableName() string {
	return "friends"
}

var Friends2Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "friends",
	Identity:  "id",
	DataTable: "friends",
	MainTable: "friends",
	DataModel: new(Friends2),
	Data:      new([]Friends2),
	MainModel: new(FriendsMainTable2),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "name", Label: ""},
		datagrid.Column{Model: "phone", Label: ""},
	},
	ColumnList: []string{"id", "name", "phone"},
	Filters: map[string]string{
		"id": "",

		"name": "",

		"phone": "",
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
	FillVirtualColumns: fillVirtualColumnsFriends2,
}

func fillVirtualColumnsFriends2(rowsPre interface{}) interface{} {
	return rowsPre
}
