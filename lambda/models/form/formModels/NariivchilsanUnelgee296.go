package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type NariivchilsanUnelgee296 struct {
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

type SubNariinUnelgeeFileNariivchilsanUnelgee296 struct {
	FileNer         *string `gorm:"column:file_ner" json:"file_ner"`
	ID              int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NariinUnelgeeID *int    `gorm:"column:nariin_unelgee_id" json:"nariin_unelgee_id"`
}

func (s *SubNariinUnelgeeFileNariivchilsanUnelgee296) TableName() string {
	return "sub_nariin_unelgee_file"
}

//  TableName sets the insert table name for this struct type
func (n *NariivchilsanUnelgee296) TableName() string {
	return "nariivchilsan_unelgee"
}
