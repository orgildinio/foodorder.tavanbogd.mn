package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewTzSungalt75 struct {
	AdminConfirmStatus *string    `gorm:"column:admin_confirm_status" json:"admin_confirm_status"`
	BNer               *string    `gorm:"column:b_ner" json:"b_ner"`
	ConfirmStatus      *string    `gorm:"column:confirm_status" json:"confirm_status"`
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	HuchinteiOgnoo     *DB.Date   `gorm:"column:huchintei_ognoo" json:"huchintei_ognoo"`
	ID                 *int       `gorm:"column:id" json:"id"`
	SungahEseh         *string    `gorm:"column:sungah_eseh" json:"sungah_eseh"`
	SungasanOgnoo      *DB.Date   `gorm:"column:sungasan_ognoo" json:"sungasan_ognoo"`
	TzGerchilgee       *string    `gorm:"column:tz_gerchilgee" json:"tz_gerchilgee"`
	TzGerchilgeeDugaar *string    `gorm:"column:tz_gerchilgee_dugaar" json:"tz_gerchilgee_dugaar"`
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type TzHugtsaaSungahMainTable75 struct {
	BID                    *int       `gorm:"column:b_id" json:"b_id"`
	ConfirmStatusID        *int       `gorm:"column:confirm_status_id" json:"confirm_status_id"`
	ConfirmStatusIDByAdmin *int       `gorm:"column:confirm_status_id_by_admin" json:"confirm_status_id_by_admin"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt              *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	HuchinteiOgnoo         *DB.Date   `gorm:"column:huchintei_ognoo" json:"huchintei_ognoo"`
	ID                     int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Shiidwer               *string    `gorm:"column:shiidwer" json:"shiidwer"`
	ShiidwerHawsralt       *string    `gorm:"column:shiidwer_hawsralt" json:"shiidwer_hawsralt"`
	SungahEseh             *string    `gorm:"column:sungah_eseh" json:"sungah_eseh"`
	SungahShaltgaan        *string    `gorm:"column:sungah_shaltgaan" json:"sungah_shaltgaan"`
	SungasanOgnoo          *DB.Date   `gorm:"column:sungasan_ognoo" json:"sungasan_ognoo"`
	Tailbar                *string    `gorm:"column:tailbar" json:"tailbar"`
	TzGerchilgee           *string    `gorm:"column:tz_gerchilgee" json:"tz_gerchilgee"`
	TzGerchilgeeDugaar     *string    `gorm:"column:tz_gerchilgee_dugaar" json:"tz_gerchilgee_dugaar"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (t *TzHugtsaaSungahMainTable75) TableName() string {
	return "tz_hugtsaa_sungah"
}

//  TableName sets the insert table name for this struct type
func (v *ViewTzSungalt75) TableName() string {
	return "view_tz_sungalt"
}

var ViewTzSungalt75Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Тусгай зөвшөөрлийн хугацаа сунгах",
	Identity:  "id",
	DataTable: "view_tz_sungalt",
	MainTable: "tz_hugtsaa_sungah",
	DataModel: new(ViewTzSungalt75),
	Data:      new([]ViewTzSungalt75),
	MainModel: new(TzHugtsaaSungahMainTable75),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "b_ner", Label: "Байгууллага "},
		datagrid.Column{Model: "sungah_eseh", Label: "Тусгай зөвшөөрлийн хугацаа сунгах эсэх "},
		datagrid.Column{Model: "tz_gerchilgee_dugaar", Label: "гэрчилгээний дугаар "},
		datagrid.Column{Model: "tz_gerchilgee", Label: "гэрчилгээ"},
		datagrid.Column{Model: "confirm_status", Label: "Төлөв"},
		datagrid.Column{Model: "admin_confirm_status", Label: "Төлөв админ"},
		datagrid.Column{Model: "sungasan_ognoo", Label: "Сунгасан огноо "},
		datagrid.Column{Model: "huchintei_ognoo", Label: "Хүчинтэй огноо"},
	},
	ColumnList: []string{"id", "b_ner", "sungah_eseh", "tz_gerchilgee_dugaar", "tz_gerchilgee", "confirm_status", "admin_confirm_status", "sungasan_ognoo", "huchintei_ognoo"},
	Filters: map[string]string{
		"id": "",

		"b_ner": "",

		"b_id": "",

		"sungah_eseh": "",

		"sungah_shaltgaan": "",

		"tz_gerchilgee_dugaar": "",

		"tz_gerchilgee": "",

		"confirm_status": "",

		"confirm_status_id": "",

		"admin_confirm_status": "",

		"confirm_status_id_by_admin": "",

		"shiidwer": "",

		"shiidwer_hawsralt": "",

		"sungasan_ognoo": "",

		"huchintei_ognoo": "",

		"tailbar": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",
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
	FillVirtualColumns: fillVirtualColumnsViewTzSungalt75,
}

func fillVirtualColumnsViewTzSungalt75(rowsPre interface{}) interface{} {
	return rowsPre
}
