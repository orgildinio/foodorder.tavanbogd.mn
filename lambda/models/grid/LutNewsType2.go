package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutNewsType2 struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsType *string `gorm:"column:news_type" json:"news_type"`
}

type LutNewsTypeMainTable2 struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsType *string `gorm:"column:news_type" json:"news_type"`
}

func (l *LutNewsTypeMainTable2) TableName() string {
	return "lut_news_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutNewsType2) TableName() string {
	return "lut_news_type"
}

var LutNewsType2Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Мэдээний төрөл",
	Identity:  "id",
	DataTable: "lut_news_type",
	MainTable: "lut_news_type",
	DataModel: new(LutNewsType2),
	Data:      new([]LutNewsType2),
	MainModel: new(LutNewsTypeMainTable2),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "news_type", Label: "Мэдээний төрөл"},
	},
	ColumnList: []string{"id", "news_type"},
	Filters: map[string]string{
		"id": "",

		"news_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutNewsType2,
}

func fillVirtualColumnsLutNewsType2(rowsPre interface{}) interface{} {
	return rowsPre
}
