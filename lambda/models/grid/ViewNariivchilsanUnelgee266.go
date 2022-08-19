package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewNariivchilsanUnelgee266 struct {
	Aimagname             *string    `gorm:"column:aimagname" json:"aimagname"`
	BNer                  *string    `gorm:"column:b_ner" json:"b_ner"`
	Bagname               *string    `gorm:"column:bagname" json:"bagname"`
	ConfirmStatusz        *string    `gorm:"column:confirm_statusz" json:"confirm_statusz"`
	CreatedAt             *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt             *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DetailedEvalutionType *string    `gorm:"column:detailed_evalution_type" json:"detailed_evalution_type"`
	DisasterType          *string    `gorm:"column:disaster_type" json:"disaster_type"`
	ID                    *int       `gorm:"column:id" json:"id"`
	NariinUnelgeeEn       *string    `gorm:"column:nariin_unelgee_en" json:"nariin_unelgee_en"`
	NariinUnelgeeMn       *string    `gorm:"column:nariin_unelgee_mn" json:"nariin_unelgee_mn"`
	Sumname               *string    `gorm:"column:sumname" json:"sumname"`
	UpdatedAt             *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas                  *int       `gorm:"column:utas" json:"utas"`
}

type NariivchilsanUnelgeeMainTable266 struct {
	AimagID             *int       `gorm:"column:aimag_id" json:"aimag_id"`
	AyulTorolID         *int       `gorm:"column:ayul_torol_id" json:"ayul_torol_id"`
	Bagid               *int       `gorm:"column:bagid" json:"bagid"`
	BaiguullagaID       *int       `gorm:"column:baiguullaga_id" json:"baiguullaga_id"`
	Bairshil            *string    `gorm:"column:bairshil" json:"bairshil"`
	BbAjilTushilga      *int       `gorm:"column:bb_ajil_tushilga" json:"bb_ajil_tushilga"`
	BbAlbanTushaal      *string    `gorm:"column:bb_alban_tushaal" json:"bb_alban_tushaal"`
	BbBaiguullaga       *string    `gorm:"column:bb_baiguullaga" json:"bb_baiguullaga"`
	BbBolovsrol         *string    `gorm:"column:bb_bolovsrol" json:"bb_bolovsrol"`
	BbMail              *string    `gorm:"column:bb_mail" json:"bb_mail"`
	BbMergejil          *string    `gorm:"column:bb_mergejil" json:"bb_mergejil"`
	BbNer               *string    `gorm:"column:bb_ner" json:"bb_ner"`
	BbOvog              *string    `gorm:"column:bb_ovog" json:"bb_ovog"`
	BbRegister          *string    `gorm:"column:bb_register" json:"bb_register"`
	BbUtas              *int       `gorm:"column:bb_utas" json:"bb_utas"`
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt           *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Dugnelt             *string    `gorm:"column:dugnelt" json:"dugnelt"`
	DugneltDOgnoo       *DB.Date   `gorm:"column:dugnelt_d_ognoo" json:"dugnelt_d_ognoo"`
	DugneltEOgnoo       *DB.Date   `gorm:"column:dugnelt_e_ognoo" json:"dugnelt_e_ognoo"`
	DugneltUrDun        *string    `gorm:"column:dugnelt_ur_dun" json:"dugnelt_ur_dun"`
	FileHavsarsan       *string    `gorm:"column:file_havsarsan" json:"file_havsarsan"`
	GamshigSudalgaaEseh *int       `gorm:"column:gamshig_sudalgaa_eseh" json:"gamshig_sudalgaa_eseh"`
	Hayg                *string    `gorm:"column:hayg" json:"hayg"`
	ID                  int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Mail                *string    `gorm:"column:mail" json:"mail"`
	NariinUnelgeeEn     *string    `gorm:"column:nariin_unelgee_en" json:"nariin_unelgee_en"`
	NariinUnelgeeMn     *string    `gorm:"column:nariin_unelgee_mn" json:"nariin_unelgee_mn"`
	Ognoo               *DB.Date   `gorm:"column:ognoo" json:"ognoo"`
	Solbiltsol          *string    `gorm:"column:solbiltsol" json:"solbiltsol"`
	Sumid               *int       `gorm:"column:sumid" json:"sumid"`
	Tailbar             *string    `gorm:"column:tailbar" json:"tailbar"`
	TdAvsanOgnoo        *DB.Date   `gorm:"column:td_avsan_ognoo" json:"td_avsan_ognoo"`
	TdMateralFile       *string    `gorm:"column:td_materal_file" json:"td_materal_file"`
	TolovID             *int       `gorm:"column:tolov_id" json:"tolov_id"`
	TolovsID            *int       `gorm:"column:tolovs_id" json:"tolovs_id"`
	UHiisenArgachlalID  *int       `gorm:"column:u_hiisen_argachlal_id" json:"u_hiisen_argachlal_id"`
	UHiisenBairshil     *string    `gorm:"column:u_hiisen_bairshil" json:"u_hiisen_bairshil"`
	UHiisenNerID        *int       `gorm:"column:u_hiisen_ner_id" json:"u_hiisen_ner_id"`
	UHiisenSalbarID     *int       `gorm:"column:u_hiisen_salbar_id" json:"u_hiisen_salbar_id"`
	UIChiglelID         *int       `gorm:"column:ui_chiglel_id" json:"ui_chiglel_id"`
	UnelgeeAanbNer      *string    `gorm:"column:unelgee_aanb_ner" json:"unelgee_aanb_ner"`
	UnelgeeTorolID      *int       `gorm:"column:unelgee_torol_id" json:"unelgee_torol_id"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas                *int       `gorm:"column:utas" json:"utas"`
	Zurag               *string    `gorm:"column:zurag" json:"zurag"`
}

func (n *NariivchilsanUnelgeeMainTable266) TableName() string {
	return "nariivchilsan_unelgee"
}

//  TableName sets the insert table name for this struct type
func (v *ViewNariivchilsanUnelgee266) TableName() string {
	return "view_nariivchilsan_unelgee"
}

var ViewNariivchilsanUnelgee266Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Нарийвчилсан үнэлгээ /Байгууллага /",
	Identity:  "id",
	DataTable: "view_nariivchilsan_unelgee",
	MainTable: "nariivchilsan_unelgee",
	DataModel: new(ViewNariivchilsanUnelgee266),
	Data:      new([]ViewNariivchilsanUnelgee266),
	MainModel: new(NariivchilsanUnelgeeMainTable266),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "confirm_statusz", Label: "Төлөв"},
		datagrid.Column{Model: "nariin_unelgee_mn", Label: "Нарийвчилсан үнэлгээ"},
		datagrid.Column{Model: "nariin_unelgee_en", Label: "Нарийвчилсан үнэлгээний төрөл "},
		datagrid.Column{Model: "utas", Label: "Утас"},
		datagrid.Column{Model: "aimagname", Label: "Аймаг / Нийслэл"},
		datagrid.Column{Model: "sumname", Label: "Сум/Дүүрэг"},
		datagrid.Column{Model: "bagname", Label: "Баг/ Хороо"},
		datagrid.Column{Model: "b_ner", Label: "Байгууллага "},
		datagrid.Column{Model: "disaster_type", Label: "Аюулын төрөл "},
		datagrid.Column{Model: "detailed_evalution_type", Label: "Нарийвчилсан үнэлгээний төрөл "},
	},
	ColumnList: []string{"id", "confirm_statusz", "nariin_unelgee_mn", "nariin_unelgee_en", "utas", "aimagname", "sumname", "bagname", "b_ner", "disaster_type", "detailed_evalution_type"},
	Filters: map[string]string{
		"id": "",

		"confirm_statusz": "",

		"nariin_unelgee_mn": "",

		"nariin_unelgee_en": "",

		"baiguullaga_id": "",

		"aimag_id": "",

		"sumid": "",

		"bagid": "",

		"bairshil": "",

		"ayul_torol_id": "",

		"ognoo": "",

		"unelgee_aanb_ner": "",

		"ui_chiglel_id": "",

		"utas": "",

		"mail": "",

		"hayg": "",

		"solbiltsol": "",

		"u_hiisen_ner_id": "",

		"u_hiisen_salbar_id": "",

		"u_hiisen_bairshil": "",

		"u_hiisen_argachlal_id": "",

		"bb_ovog": "",

		"bb_ner": "",

		"bb_baiguullaga": "",

		"bb_alban_tushaal": "",

		"bb_register": "",

		"bb_bolovsrol": "",

		"bb_mergejil": "",

		"bb_ajil_tushilga": "",

		"bb_mail": "",

		"bb_utas": "",

		"gamshig_sudalgaa_eseh": "",

		"dugnelt_ur_dun": "",

		"zurag": "",

		"dugnelt_e_ognoo": "",

		"dugnelt_d_ognoo": "",

		"tolov_id": "",

		"td_avsan_ognoo": "",

		"td_materal_file": "",

		"dugnelt": "",

		"tolovs_id": "",

		"tailbar": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"unelgee_torol_id": "",

		"aimagname": "",

		"sumname": "",

		"bagname": "",

		"b_ner": "",

		"disaster_type": "",

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
	FillVirtualColumns: fillVirtualColumnsViewNariivchilsanUnelgee266,
}

func fillVirtualColumnsViewNariivchilsanUnelgee266(rowsPre interface{}) interface{} {
	return rowsPre
}
