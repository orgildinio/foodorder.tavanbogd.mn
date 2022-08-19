package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutRiskAssessment90 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RiskAssessment *string `gorm:"column:risk_assessment" json:"risk_assessment"`
}

type LutRiskAssessmentMainTable90 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RiskAssessment *string `gorm:"column:risk_assessment" json:"risk_assessment"`
}

func (l *LutRiskAssessmentMainTable90) TableName() string {
	return "lut_risk_assessment"
}

//  TableName sets the insert table name for this struct type
func (l *LutRiskAssessment90) TableName() string {
	return "lut_risk_assessment"
}

var LutRiskAssessment90Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Мэргэжлийн зөвлөлийн бүтэц",
	Identity:  "id",
	DataTable: "lut_risk_assessment",
	MainTable: "lut_risk_assessment",
	DataModel: new(LutRiskAssessment90),
	Data:      new([]LutRiskAssessment90),
	MainModel: new(LutRiskAssessmentMainTable90),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "risk_assessment", Label: "Мэргэжлийн зөвлөлийн бүтэц"},
	},
	ColumnList: []string{"id", "risk_assessment"},
	Filters: map[string]string{
		"id": "",

		"risk_assessment": "",
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
	FillVirtualColumns: fillVirtualColumnsLutRiskAssessment90,
}

func fillVirtualColumnsLutRiskAssessment90(rowsPre interface{}) interface{} {
	return rowsPre
}
