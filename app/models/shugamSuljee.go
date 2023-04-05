package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

type ShugamShuljee struct {
	AimagID           int        `gorm:"column:aimag_id" json:"aimag_id"`
	AlbaCodeNerID     int        `gorm:"column:alba_code_ner_id" json:"alba_code_ner_id"`
	AshiglasanJil     int        `gorm:"column:ashiglasan_jil" json:"ashiglasan_jil"`
	BagID             int        `gorm:"column:bag_id" json:"bag_id"`
	BaiguullagaID     int        `gorm:"column:baiguullaga_id" json:"baiguullaga_id"`
	BugdUne           float32    `gorm:"column:bugd_une" json:"bugd_une"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DiameterID        int        `gorm:"column:diameter_id" json:"diameter_id"`
	EhUusverID        int        `gorm:"column:eh_uusver_id" json:"eh_uusver_id"`
	EzemshigchCode    string     `gorm:"column:ezemshigch_code" json:"ezemshigch_code"`
	EzemshigchNer     string     `gorm:"column:ezemshigch_ner" json:"ezemshigch_ner"`
	Geo               string     `gorm:"column:geo" json:"geo"`
	GisCatID          int        `gorm:"column:gis_cat_id" json:"gis_cat_id"`
	GisLayerID        int        `gorm:"column:gis_layer_id" json:"gis_layer_id"`
	HDugaar           string     `gorm:"column:h_dugaar" json:"h_dugaar"`
	Hayag             string     `gorm:"column:hayag" json:"hayag"`
	HemjihNegjID      int        `gorm:"column:hemjih_negj_id" json:"hemjih_negj_id"`
	HesegID           int        `gorm:"column:heseg_id" json:"heseg_id"`
	ID                int        `gorm:"column:id;primary_key" json:"id"`
	MateralID         int        `gorm:"column:materal_id" json:"materal_id"`
	NbAngilalID       int        `gorm:"column:nb_angilal_id" json:"nb_angilal_id"`
	NegjUne           float32    `gorm:"column:negj_une" json:"negj_une"`
	OgnooAsh          DB.Date    `gorm:"column:ognoo_ash" json:"ognoo_ash"`
	SumID             int        `gorm:"column:sum_id" json:"sum_id"`
	TolovID           int        `gorm:"column:tolov_id" json:"tolov_id"`
	UndsenHorongoCode string     `gorm:"column:undsen_horongo_code" json:"undsen_horongo_code"`
	UndsenHorongoNer  string     `gorm:"column:undsen_horongo_ner" json:"undsen_horongo_ner"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Urt               float32    `gorm:"column:urt" json:"urt"`
	UserID            int        `gorm:"column:user_id" json:"user_id"`
	ZoriulaltID       int        `gorm:"column:zoriulalt_id" json:"zoriulalt_id"`
}

func (s *ShugamShuljee) TableName() string {
	return "shugam_shuljee"
}
