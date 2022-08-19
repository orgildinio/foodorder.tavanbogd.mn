package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSocialType5 struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SocialType *string `gorm:"column:social_type" json:"social_type"`
}

type LutSocialTypeMainTable5 struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SocialType *string `gorm:"column:social_type" json:"social_type"`
}

func (l *LutSocialTypeMainTable5) TableName() string {
	return "lut_social_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutSocialType5) TableName() string {
	return "lut_social_type"
}

var LutSocialType5Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Social type",
	Identity:  "id",
	DataTable: "lut_social_type",
	MainTable: "lut_social_type",
	DataModel: new(LutSocialType5),
	Data:      new([]LutSocialType5),
	MainModel: new(LutSocialTypeMainTable5),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "social_type", Label: "Social type"},
	},
	ColumnList: []string{"id", "social_type"},
	Filters: map[string]string{
		"id": "",

		"social_type": "",
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
	FillVirtualColumns: fillVirtualColumnsLutSocialType5,
}

func fillVirtualColumnsLutSocialType5(rowsPre interface{}) interface{} {
	return rowsPre
}
