package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

type TeevriinHeregsel struct {
	AimagID           int        `gorm:"column:aimag_id" json:"aimag_id"`
	AlbaCodeNerID     int        `gorm:"column:alba_code_ner_id" json:"alba_code_ner_id"`
	ArliinDugaar      string     `gorm:"column:arliin_dugaar" json:"arliin_dugaar"`
	AshiglaltJil      int        `gorm:"column:ashiglalt_jil" json:"ashiglalt_jil"`
	AshiglaltOgnoo    DB.Date    `gorm:"column:ashiglalt_ognoo" json:"ashiglalt_ognoo"`
	BagID             int        `gorm:"column:bag_id" json:"bag_id"`
	BugdUne           float32    `gorm:"column:bugd_une" json:"bugd_une"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt         *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	EhUusverID        int        `gorm:"column:eh_uusver_id" json:"eh_uusver_id"`
	EzemshigchCode    string     `gorm:"column:ezemshigch_code" json:"ezemshigch_code"`
	EzemshigchNer     string     `gorm:"column:ezemshigch_ner" json:"ezemshigch_ner"`
	Geo               string     `gorm:"column:geo" json:"geo"`
	Hayag             string     `gorm:"column:hayag" json:"hayag"`
	HesegID           int        `gorm:"column:heseg_id" json:"heseg_id"`
	HorongoBaidal     string     `gorm:"column:horongo_baidal" json:"horongo_baidal"`
	HuchinChadal      string     `gorm:"column:huchin_chadal" json:"huchin_chadal"`
	ID                int        `gorm:"column:id;primary_key" json:"id"`
	Mark              string     `gorm:"column:mark" json:"mark"`
	NbAngilalID       int        `gorm:"column:nb_angilal_id" json:"nb_angilal_id"`
	NegjUne           float32    `gorm:"column:negj_une" json:"negj_une"`
	NiitYvahGuilt     int        `gorm:"column:niit_yvah_guilt" json:"niit_yvah_guilt"`
	Qrcode            string     `gorm:"column:qrcode" json:"qrcode"`
	SerialDugaar      string     `gorm:"column:serial_dugaar" json:"serial_dugaar"`
	SumID             int        `gorm:"column:sum_id" json:"sum_id"`
	THGerchilgee      string     `gorm:"column:t_h_gerchilgee" json:"t_h_gerchilgee"`
	TolovID           int        `gorm:"column:tolov_id" json:"tolov_id"`
	UldID             int        `gorm:"column:uld_id" json:"uld_id"`
	UlsiinDugaar      string     `gorm:"column:ulsiin_dugaar" json:"ulsiin_dugaar"`
	UndsenHorongoCode string     `gorm:"column:undsen_horongo_code" json:"undsen_horongo_code"`
	UndsenHorongoNer  string     `gorm:"column:undsen_horongo_ner" json:"undsen_horongo_ner"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID            int        `gorm:"column:user_id" json:"user_id"`
	YvsanGuilt        string     `gorm:"column:yvsan_guilt" json:"yvsan_guilt"`
	ZoriulaltID       int        `gorm:"column:zoriulalt_id" json:"zoriulalt_id"`
}

func (t *TeevriinHeregsel) TableName() string {
	return "teevriin_heregsel"
}
