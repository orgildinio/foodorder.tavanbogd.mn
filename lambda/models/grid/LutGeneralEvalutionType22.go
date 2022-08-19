package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutGeneralEvalutionType22 struct {
	GeneralEvalutionType *string `gorm:"column:general_evalution_type" json:"general_evalution_type"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutGeneralEvalutionTypeMainTable22 struct {
	GeneralEvalutionType *string `gorm:"column:general_evalution_type" json:"general_evalution_type"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutGeneralEvalutionTypeMainTable22) TableName() string {
	return "lut_general_evalution_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutGeneralEvalutionType22) TableName() string {
	return "lut_general_evalution_type"
}

var LutGeneralEvalutionType22Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Ерөнхий үнэлгээний төрөл ",
	Identity:  "id",
	DataTable: "lut_general_evalution_type",
	MainTable: "lut_general_evalution_type",
	DataModel: new(LutGeneralEvalutionType22),
	Data:      new([]LutGeneralEvalutionType22),
	MainModel: new(LutGeneralEvalutionTypeMainTable22),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "general_evalution_type", Label: "Ерөнхий үнэлгээний төрөл "},
	},
	ColumnList: []string{"id", "general_evalution_type"},
	Filters: map[string]string{
		"id": "",

		"general_evalution_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutGeneralEvalutionType22,
}

func fillVirtualColumnsLutGeneralEvalutionType22(rowsPre interface{}) interface{} {
	return rowsPre
}
