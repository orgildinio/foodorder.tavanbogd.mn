package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutMalNasTorol190 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalNasTorol *string `gorm:"column:mal_nas_torol" json:"mal_nas_torol"`
}

type LutMalNasTorolMainTable190 struct {
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalNasTorol *string `gorm:"column:mal_nas_torol" json:"mal_nas_torol"`
}

func (l *LutMalNasTorolMainTable190) TableName() string {
	return "lut_mal_nas_torol"
}

//  TableName sets the insert table name for this struct type
func (l *LutMalNasTorol190) TableName() string {
	return "lut_mal_nas_torol"
}

var LutMalNasTorol190Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Малын нас",
	Identity:  "id",
	DataTable: "lut_mal_nas_torol",
	MainTable: "lut_mal_nas_torol",
	DataModel: new(LutMalNasTorol190),
	Data:      new([]LutMalNasTorol190),
	MainModel: new(LutMalNasTorolMainTable190),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "mal_nas_torol", Label: "Малын нас"},
	},
	ColumnList: []string{"id", "mal_nas_torol"},
	Filters: map[string]string{
		"id": "",

		"mal_nas_torol": "",
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
	FillVirtualColumns: fillVirtualColumnsLutMalNasTorol190,
}

func fillVirtualColumnsLutMalNasTorol190(rowsPre interface{}) interface{} {
	return rowsPre
}
