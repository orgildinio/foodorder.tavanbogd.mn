package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewNariichivlsanUnelgeeAdmin269 struct {
	ConfirmStatus   *string    `gorm:"column:confirm_status" json:"confirm_status"`
	ConfirmStatusz  *string    `gorm:"column:confirm_statusz" json:"confirm_statusz"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Dugnelt         *string    `gorm:"column:dugnelt" json:"dugnelt"`
	ID              *int       `gorm:"column:id" json:"id"`
	NariinUnelgeeMn *string    `gorm:"column:nariin_unelgee_mn" json:"nariin_unelgee_mn"`
	TdAvsanOgnoo    *DB.Date   `gorm:"column:td_avsan_ognoo" json:"td_avsan_ognoo"`
	TdMateralFile   *string    `gorm:"column:td_materal_file" json:"td_materal_file"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type NariivchilsanUnelgeeMainTable269 struct {
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

func (n *NariivchilsanUnelgeeMainTable269) TableName() string {
	return "nariivchilsan_unelgee"
}

//  TableName sets the insert table name for this struct type
func (v *ViewNariichivlsanUnelgeeAdmin269) TableName() string {
	return "view_nariichivlsan_unelgee_admin"
}

var ViewNariichivlsanUnelgeeAdmin269Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Нарийвчилсан үнэлгээ /Admin /",
	Identity:  "id",
	DataTable: "view_nariichivlsan_unelgee_admin",
	MainTable: "nariivchilsan_unelgee",
	DataModel: new(ViewNariichivlsanUnelgeeAdmin269),
	Data:      new([]ViewNariichivlsanUnelgeeAdmin269),
	MainModel: new(NariivchilsanUnelgeeMainTable269),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "nariin_unelgee_mn", Label: "Нарийвчилсан үнэлгээ"},
		datagrid.Column{Model: "confirm_status", Label: "Төлөв"},
		datagrid.Column{Model: "td_avsan_ognoo", Label: "Тайлан хүлээн авсан огноо "},
		datagrid.Column{Model: "td_materal_file", Label: "Тайлангийн материал "},
		datagrid.Column{Model: "dugnelt", Label: "Дүгнэлт "},
		datagrid.Column{Model: "confirm_statusz", Label: "Төлөв/Байгууллага/"},
	},
	ColumnList: []string{"id", "nariin_unelgee_mn", "confirm_status", "td_avsan_ognoo", "td_materal_file", "dugnelt", "confirm_statusz"},
	Filters: map[string]string{
		"id": "",

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

		"confirm_status": "",

		"td_avsan_ognoo": "",

		"td_materal_file": "",

		"dugnelt": "",

		"tolovs_id": "",

		"tailbar": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"unelgee_torol_id": "",

		"confirm_statusz": "",
	},
	Relations:   []models.GridRelation{},
	Condition:   "tolov_id = 1",
	Aggergation: "",
	BeforeFetch: nil,

	AfterFetch: nil,

	BeforeDelete: nil,

	AfterDelete: nil,

	BeforePrint:        nil,
	TriggerNameSpace:   "",
	FillVirtualColumns: fillVirtualColumnsViewNariichivlsanUnelgeeAdmin269,
}

func fillVirtualColumnsViewNariichivlsanUnelgeeAdmin269(rowsPre interface{}) interface{} {
	return rowsPre
}
