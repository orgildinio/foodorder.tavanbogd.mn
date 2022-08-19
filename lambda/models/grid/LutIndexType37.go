package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutIndexType37 struct {
	ErsdeliinIndex *string `gorm:"column:ersdeliin_index" json:"ersdeliin_index"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutIndexTypeMainTable37 struct {
	ErsdeliinIndex *string `gorm:"column:ersdeliin_index" json:"ersdeliin_index"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutIndexTypeMainTable37) TableName() string {
	return "lut_index_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutIndexType37) TableName() string {
	return "lut_index_type"
}

var LutIndexType37Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Эрсдлийн индекс",
	Identity:  "id",
	DataTable: "lut_index_type",
	MainTable: "lut_index_type",
	DataModel: new(LutIndexType37),
	Data:      new([]LutIndexType37),
	MainModel: new(LutIndexTypeMainTable37),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ersdeliin_index", Label: "Эрсдлийн индекс"},
	},
	ColumnList: []string{"id", "ersdeliin_index"},
	Filters: map[string]string{
		"id": "",

		"ersdeliin_index": "",
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
	FillVirtualColumns: fillVirtualColumnsLutIndexType37,
}

func fillVirtualColumnsLutIndexType37(rowsPre interface{}) interface{} {
	return rowsPre
}
