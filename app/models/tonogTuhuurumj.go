package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

type TonogTohooromj struct {
	AimagID             int        `gorm:"column:aimag_id" json:"aimag_id"`
	AjilahTsag          string     `gorm:"column:ajilah_tsag" json:"ajilah_tsag"`
	AjilsanTsag         string     `gorm:"column:ajilsan_Tsag" json:"ajilsan_Tsag"`
	AktFile             string     `gorm:"column:akt_file" json:"akt_file"`
	AlbaCodeNerID       int        `gorm:"column:alba_code_ner_id" json:"alba_code_ner_id"`
	AshiglaltJil        int        `gorm:"column:ashiglalt_jil" json:"ashiglalt_jil"`
	AshiglsanJil        int        `gorm:"column:ashiglsan_jil" json:"ashiglsan_jil"`
	BagID               int        `gorm:"column:bag_id" json:"bag_id"`
	BarilgaID           int        `gorm:"column:barilga_id" json:"barilga_id"`
	Bugd                float32    `gorm:"column:bugd" json:"bugd"`
	CreatedAt           *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt           *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	EhUusverID          int        `gorm:"column:eh_uusver_id" json:"eh_uusver_id"`
	EzemshigchCode      string     `gorm:"column:ezemshigch_code" json:"ezemshigch_code"`
	EzemshigchNer       string     `gorm:"column:ezemshigch_ner" json:"ezemshigch_ner"`
	Geo                 string     `gorm:"column:geo" json:"geo"`
	GisCatID            int        `gorm:"column:gis_cat_id" json:"gis_cat_id"`
	GisLayerID          int        `gorm:"column:gis_layer_id" json:"gis_layer_id"`
	Hayag               string     `gorm:"column:hayag" json:"hayag"`
	HemjihNegjID        int        `gorm:"column:hemjih_negj_id" json:"hemjih_negj_id"`
	HesegID             int        `gorm:"column:heseg_id" json:"heseg_id"`
	HorongoBairshil     string     `gorm:"column:horongo_bairshil" json:"horongo_bairshil"`
	HorongoZeregID      int        `gorm:"column:horongo_zereg_id" json:"horongo_zereg_id"`
	HuchinChadal        string     `gorm:"column:huchin_chadal" json:"huchin_chadal"`
	ID                  int        `gorm:"column:id;primary_key" json:"id"`
	Mark                string     `gorm:"column:mark" json:"mark"`
	NbAngilalID         int        `gorm:"column:nb_angilal_id" json:"nb_angilal_id"`
	NegjUne             float32    `gorm:"column:negj_une" json:"negj_une"`
	QrCode              string     `gorm:"column:qr_code" json:"qr_code"`
	SOgnoo              DB.Date    `gorm:"column:s_ognoo" json:"s_ognoo"`
	SerialNum           string     `gorm:"column:serial_num" json:"serial_num"`
	ShugalsuljeeID      int        `gorm:"column:shugalsuljee_id" json:"shugalsuljee_id"`
	ShugamsuljeeHudagID int        `gorm:"column:shugamsuljee_hudag_id" json:"shugamsuljee_hudag_id"`
	SumID               int        `gorm:"column:sum_id" json:"sum_id"`
	TolovID             int        `gorm:"column:tolov_id" json:"tolov_id"`
	TooHemjee           float32    `gorm:"column:too_hemjee" json:"too_hemjee"`
	TtNer               string     `gorm:"column:tt_ner" json:"tt_ner"`
	TtObjecttorolID     int        `gorm:"column:tt_objecttorol_id" json:"tt_objecttorol_id"`
	UlsID               int        `gorm:"column:uls_id" json:"uls_id"`
	UndsenHorongoCode   string     `gorm:"column:undsen_horongo_code" json:"undsen_horongo_code"`
	UpdatedAt           *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID              int        `gorm:"column:user_id" json:"user_id"`
	ZoriulaltID         int        `gorm:"column:zoriulalt_id" json:"zoriulalt_id"`
}

func (t *TonogTohooromj) TableName() string {
	return "tonog_tohooromj"
}

type LutHemjihNegj struct {
	HemjihNegj string `gorm:"column:hemjih_negj" json:"hemjih_negj"`
	ID         int    `gorm:"column:id;primary_key" json:"id"`
}

func (l *LutHemjihNegj) TableName() string {
	return "lut_hemjih_negj"
}

type LutZoriulalt struct {
	ID            int    `gorm:"column:id;primary_key" json:"id"`
	ObjectTorolID int    `gorm:"column:object_torol_id" json:"object_torol_id"`
	Zoriulalt     string `gorm:"column:zoriulalt" json:"zoriulalt"`
}

func (l *LutZoriulalt) TableName() string {
	return "lut_zoriulalt"
}
