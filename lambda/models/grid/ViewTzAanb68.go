package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewTzAanb68 struct {
	ActivityTypeID       *int       `gorm:"column:activity_type_id" json:"activity_type_id"`
	AimagID              *int       `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname            *string    `gorm:"column:aimagname" json:"aimagname"`
	BNer                 *string    `gorm:"column:b_ner" json:"b_ner"`
	BagID                *int       `gorm:"column:bag_id" json:"bag_id"`
	Bagname              *string    `gorm:"column:bagname" json:"bagname"`
	Bair                 *string    `gorm:"column:bair" json:"bair"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email                *string    `gorm:"column:email" json:"email"`
	EvalutionSectorID    *int       `gorm:"column:evalution_sector_id" json:"evalution_sector_id"`
	Gudamj               *string    `gorm:"column:gudamj" json:"gudamj"`
	HayagDelegrengui     *string    `gorm:"column:hayag_delegrengui" json:"hayag_delegrengui"`
	ID                   *int       `gorm:"column:id" json:"id"`
	Register             *string    `gorm:"column:register" json:"register"`
	ShiidwerHawsralt     *string    `gorm:"column:shiidwer_hawsralt" json:"shiidwer_hawsralt"`
	SumID                *int       `gorm:"column:sum_id" json:"sum_id"`
	Sumname              *string    `gorm:"column:sumname" json:"sumname"`
	TzDugaar             *string    `gorm:"column:tz_dugaar" json:"tz_dugaar"`
	TzDuusah             *DB.Date   `gorm:"column:tz_duusah" json:"tz_duusah"`
	TzOlgoson            *DB.Date   `gorm:"column:tz_olgoson" json:"tz_olgoson"`
	TzShiidwer           *string    `gorm:"column:tz_shiidwer" json:"tz_shiidwer"`
	UbGerchilgee         *string    `gorm:"column:ub_gerchilgee" json:"ub_gerchilgee"`
	UbGerchilgeeDug      *string    `gorm:"column:ub_gerchilgee_dug" json:"ub_gerchilgee_dug"`
	UilAjilgaaniiChiglel *string    `gorm:"column:uil_ajilgaanii_chiglel" json:"uil_ajilgaanii_chiglel"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas                 *string    `gorm:"column:utas" json:"utas"`
	Zahiral              *string    `gorm:"column:zahiral" json:"zahiral"`
	ZakiUtas             *string    `gorm:"column:zaki_utas" json:"zaki_utas"`
}

type TzBuhiiAanbMainTable68 struct {
	ActivityTypeID    *int       `gorm:"column:activity_type_id" json:"activity_type_id"`
	AimagID           *int       `gorm:"column:aimag_id" json:"aimag_id"`
	BNer              *string    `gorm:"column:b_ner" json:"b_ner"`
	BagID             *int       `gorm:"column:bag_id" json:"bag_id"`
	Bair              *string    `gorm:"column:bair" json:"bair"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email             *string    `gorm:"column:email" json:"email"`
	EvalutionSectorID *int       `gorm:"column:evalution_sector_id" json:"evalution_sector_id"`
	Gudamj            *string    `gorm:"column:gudamj" json:"gudamj"`
	HayagDelegrengui  *string    `gorm:"column:hayag_delegrengui" json:"hayag_delegrengui"`
	ID                int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Register          *string    `gorm:"column:register" json:"register"`
	ShiidwerHawsralt  *string    `gorm:"column:shiidwer_hawsralt" json:"shiidwer_hawsralt"`
	SumID             *int       `gorm:"column:sum_id" json:"sum_id"`
	TzDugaar          *string    `gorm:"column:tz_dugaar" json:"tz_dugaar"`
	TzDuusah          *DB.Date   `gorm:"column:tz_duusah" json:"tz_duusah"`
	TzOlgoson         *DB.Date   `gorm:"column:tz_olgoson" json:"tz_olgoson"`
	TzShiidwer        *string    `gorm:"column:tz_shiidwer" json:"tz_shiidwer"`
	UbGerchilgee      *string    `gorm:"column:ub_gerchilgee" json:"ub_gerchilgee"`
	UbGerchilgeeDug   *string    `gorm:"column:ub_gerchilgee_dug" json:"ub_gerchilgee_dug"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas              *string    `gorm:"column:utas" json:"utas"`
	Zahiral           *string    `gorm:"column:zahiral" json:"zahiral"`
	ZakiUtas          *string    `gorm:"column:zaki_utas" json:"zaki_utas"`
}

func (t *TzBuhiiAanbMainTable68) TableName() string {
	return "tz_buhii_aanb"
}

//  TableName sets the insert table name for this struct type
func (v *ViewTzAanb68) TableName() string {
	return "view_tz_aanb"
}

var ViewTzAanb68Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "ТЗ БҮХИЙ ААНБ",
	Identity:  "id",
	DataTable: "view_tz_aanb",
	MainTable: "tz_buhii_aanb",
	DataModel: new(ViewTzAanb68),
	Data:      new([]ViewTzAanb68),
	MainModel: new(TzBuhiiAanbMainTable68),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "uil_ajilgaanii_chiglel", Label: "Үйл ажиллагааны чиглэл "},
		datagrid.Column{Model: "b_ner", Label: "ААН-ын нэр"},
		datagrid.Column{Model: "ub_gerchilgee_dug", Label: "Улсын бүртгэлийн гэрчилгээний дугаар"},
		datagrid.Column{Model: "register", Label: "Регистрийн дугаар"},
		datagrid.Column{Model: "ub_gerchilgee", Label: "Улсын бүртгэлийн гэрчилгээ"},
		datagrid.Column{Model: "zahiral", Label: "Захирал"},
		datagrid.Column{Model: "zaki_utas", Label: "Захиралын утасны дугаар"},
		datagrid.Column{Model: "utas", Label: "Утасны дугаар"},
		datagrid.Column{Model: "email", Label: "Цахимшуудан"},
		datagrid.Column{Model: "gudamj", Label: "Гудамж / Хороолол"},
		datagrid.Column{Model: "bair", Label: "Байр / Хашаа / Хаалганы дугаар"},
		datagrid.Column{Model: "hayag_delegrengui", Label: "Хаяг дэлгэрэнгүй"},
		datagrid.Column{Model: "tz_dugaar", Label: "Тусгай зөвшөөрлийн гэрчилгээний дугаар "},
		datagrid.Column{Model: "tz_olgoson", Label: "Тусгай зөвшөөрөл олгосон өгнөө"},
		datagrid.Column{Model: "tz_duusah", Label: "Тусгай зөвшөөрөл дуусах огноо"},
		datagrid.Column{Model: "tz_shiidwer", Label: "Зөвшөөрөл олгосон шийдвэр"},
		datagrid.Column{Model: "shiidwer_hawsralt", Label: "Шийдвэрийн хавсралт"},
		datagrid.Column{Model: "aimagname", Label: "Аймаг / Нийслэл"},
		datagrid.Column{Model: "sumname", Label: "Сум  / Дүүрэг"},
		datagrid.Column{Model: "bagname", Label: "Баг / Хороо"},
	},
	ColumnList: []string{"id", "uil_ajilgaanii_chiglel", "b_ner", "ub_gerchilgee_dug", "register", "ub_gerchilgee", "activity_type_id", "evalution_sector_id", "zahiral", "zaki_utas", "utas", "email", "aimag_id", "sum_id", "bag_id", "gudamj", "bair", "hayag_delegrengui", "tz_dugaar", "tz_olgoson", "tz_duusah", "tz_shiidwer", "shiidwer_hawsralt", "aimagname", "sumname", "bagname"},
	Filters: map[string]string{
		"id": "",

		"uil_ajilgaanii_chiglel": "",

		"b_ner": "Text",

		"ub_gerchilgee_dug": "",

		"register": "",

		"ub_gerchilgee": "",

		"activity_type_id": "",

		"evalution_sector_id": "Select",

		"zahiral": "Text",

		"zaki_utas": "",

		"utas": "",

		"email": "",

		"aimag_id": "Select",

		"sum_id": "",

		"bag_id": "",

		"gudamj": "",

		"bair": "",

		"hayag_delegrengui": "",

		"tz_dugaar": "",

		"tz_olgoson": "",

		"tz_duusah": "",

		"tz_shiidwer": "",

		"shiidwer_hawsralt": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"aimagname": "",

		"sumname": "",

		"bagname": "",
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
	FillVirtualColumns: fillVirtualColumnsViewTzAanb68,
}

func fillVirtualColumnsViewTzAanb68(rowsPre interface{}) interface{} {
	return rowsPre
}
