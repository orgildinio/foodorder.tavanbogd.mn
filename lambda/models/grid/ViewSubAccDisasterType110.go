package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewSubAccDisasterType110 struct {
	DisasterType       *string `gorm:"column:disaster_type" json:"disaster_type"`
	ID                 *int    `gorm:"column:id" json:"id"`
	SubAccDisasterType *string `gorm:"column:sub_acc_disaster_type" json:"sub_acc_disaster_type"`
}

type LutSubAccDisasterTypeMainTable110 struct {
	AccDisasterTypeID  *int    `gorm:"column:acc_disaster_type_id" json:"acc_disaster_type_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SubAccDisasterType *string `gorm:"column:sub_acc_disaster_type" json:"sub_acc_disaster_type"`
}

func (l *LutSubAccDisasterTypeMainTable110) TableName() string {
	return "lut_sub_acc_disaster_type"
}

//  TableName sets the insert table name for this struct type
func (v *ViewSubAccDisasterType110) TableName() string {
	return "view_sub_acc_disaster_type"
}

var ViewSubAccDisasterType110Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "L Ослын дэд төрөл",
	Identity:  "id",
	DataTable: "view_sub_acc_disaster_type",
	MainTable: "lut_sub_acc_disaster_type",
	DataModel: new(ViewSubAccDisasterType110),
	Data:      new([]ViewSubAccDisasterType110),
	MainModel: new(LutSubAccDisasterTypeMainTable110),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "disaster_type", Label: "Ослын төрөл"},
		datagrid.Column{Model: "sub_acc_disaster_type", Label: "Ослын дэд төрөл"},
	},
	ColumnList: []string{"id", "disaster_type", "sub_acc_disaster_type"},
	Filters: map[string]string{
		"id": "",

		"acc_disaster_type_id": "",

		"disaster_type": "",

		"sub_acc_disaster_type": "",
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
	FillVirtualColumns: fillVirtualColumnsViewSubAccDisasterType110,
}

func fillVirtualColumnsViewSubAccDisasterType110(rowsPre interface{}) interface{} {
	return rowsPre
}
