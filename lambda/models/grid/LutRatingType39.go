package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutRatingType39 struct {
	ErsdeliinZereglel *string `gorm:"column:ersdeliin_zereglel" json:"ersdeliin_zereglel"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

type LutRatingTypeMainTable39 struct {
	ErsdeliinZereglel *string `gorm:"column:ersdeliin_zereglel" json:"ersdeliin_zereglel"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (l *LutRatingTypeMainTable39) TableName() string {
	return "lut_rating_type"
}

//  TableName sets the insert table name for this struct type
func (l *LutRatingType39) TableName() string {
	return "lut_rating_type"
}

var LutRatingType39Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Эрсдлийн зэрэглэл",
	Identity:  "id",
	DataTable: "lut_rating_type",
	MainTable: "lut_rating_type",
	DataModel: new(LutRatingType39),
	Data:      new([]LutRatingType39),
	MainModel: new(LutRatingTypeMainTable39),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ersdeliin_zereglel", Label: "Эрсдлийн зэрэглэл"},
	},
	ColumnList: []string{"id", "ersdeliin_zereglel"},
	Filters: map[string]string{
		"id": "",

		"ersdeliin_zereglel": "",
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
	FillVirtualColumns: fillVirtualColumnsLutRatingType39,
}

func fillVirtualColumnsLutRatingType39(rowsPre interface{}) interface{} {
	return rowsPre
}
