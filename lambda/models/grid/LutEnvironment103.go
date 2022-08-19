package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEnvironment103 struct {
	Environment *string `gorm:"column:environment" json:"environment"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutEnvironmentMainTable103 struct {
	Environment *string `gorm:"column:environment" json:"environment"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutEnvironmentMainTable103) TableName() string {
	return "lut_environment"
}

//  TableName sets the insert table name for this struct type
func (l *LutEnvironment103) TableName() string {
	return "lut_environment"
}

var LutEnvironment103Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Алдагдсан орчин",
	Identity:  "id",
	DataTable: "lut_environment",
	MainTable: "lut_environment",
	DataModel: new(LutEnvironment103),
	Data:      new([]LutEnvironment103),
	MainModel: new(LutEnvironmentMainTable103),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "environment", Label: "Алдагдсан орчин"},
	},
	ColumnList: []string{"id", "environment"},
	Filters: map[string]string{
		"id": "",

		"environment": "",
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
	FillVirtualColumns: fillVirtualColumnsLutEnvironment103,
}

func fillVirtualColumnsLutEnvironment103(rowsPre interface{}) interface{} {
	return rowsPre
}
