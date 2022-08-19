package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutOrgSource97 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgSource *string `gorm:"column:org_source" json:"org_source"`
}

type LutOrgSourceMainTable97 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgSource *string `gorm:"column:org_source" json:"org_source"`
}

func (l *LutOrgSourceMainTable97) TableName() string {
	return "lut_org_source"
}

//  TableName sets the insert table name for this struct type
func (l *LutOrgSource97) TableName() string {
	return "lut_org_source"
}

var LutOrgSource97Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Эх сурвалж /БГ/",
	Identity:  "id",
	DataTable: "lut_org_source",
	MainTable: "lut_org_source",
	DataModel: new(LutOrgSource97),
	Data:      new([]LutOrgSource97),
	MainModel: new(LutOrgSourceMainTable97),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "org_source", Label: "Эх сурвалж"},
	},
	ColumnList: []string{"id", "org_source"},
	Filters: map[string]string{
		"id": "",

		"org_source": "",
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
	FillVirtualColumns: fillVirtualColumnsLutOrgSource97,
}

func fillVirtualColumnsLutOrgSource97(rowsPre interface{}) interface{} {
	return rowsPre
}
