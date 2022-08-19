package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEvalutionSector12 struct {
	EvalutionSector *string `gorm:"column:evalution_sector" json:"evalution_sector"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutEvalutionSectorMainTable12 struct {
	EvalutionSector *string `gorm:"column:evalution_sector" json:"evalution_sector"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutEvalutionSectorMainTable12) TableName() string {
	return "lut_evalution_sector"
}

//  TableName sets the insert table name for this struct type
func (l *LutEvalutionSector12) TableName() string {
	return "lut_evalution_sector"
}

var LutEvalutionSector12Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Үнэлгээ хийх салбар ",
	Identity:  "id",
	DataTable: "lut_evalution_sector",
	MainTable: "lut_evalution_sector",
	DataModel: new(LutEvalutionSector12),
	Data:      new([]LutEvalutionSector12),
	MainModel: new(LutEvalutionSectorMainTable12),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "evalution_sector", Label: "Үнэлгээ хийх салбар "},
	},
	ColumnList: []string{"id", "evalution_sector"},
	Filters: map[string]string{
		"id": "",

		"evalution_sector": "",
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
	FillVirtualColumns: fillVirtualColumnsLutEvalutionSector12,
}

func fillVirtualColumnsLutEvalutionSector12(rowsPre interface{}) interface{} {
	return rowsPre
}
