package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDetailedEvalutionType24 struct {
	DetailedEvalutionType *string `gorm:"column:detailed_evalution_type" json:"detailed_evalution_type"`
	ID                    int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutDetailedEvalutionTypeMainTable24 struct {
	DetailedEvalutionType *string `gorm:"column:detailed_evalution_type" json:"detailed_evalution_type"`
	ID                    int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutDetailedEvalutionTypeMainTable24) TableName() string {
	return "lut_detailed_evalution_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutDetailedEvalutionType24) TableName() string {
	return "lut_detailed_evalution_type"
}

var LutDetailedEvalutionType24Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Ерөнхий үнэлгээний төрөл ",
	Identity:  "id",
	DataTable: "lut_detailed_evalution_type",
	MainTable: "lut_detailed_evalution_type",
	DataModel: new(LutDetailedEvalutionType24),
	Data:      new([]LutDetailedEvalutionType24),
	MainModel: new(LutDetailedEvalutionTypeMainTable24),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "detailed_evalution_type", Label: "Ерөнхий үнэлгээний төрөл "},
	},
	ColumnList: []string{"id", "detailed_evalution_type"},
	Filters: map[string]string{
		"id": "",

		"detailed_evalution_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutDetailedEvalutionType24,
}

func fillVirtualColumnsLutDetailedEvalutionType24(rowsPre interface{}) interface{} {
	return rowsPre
}
