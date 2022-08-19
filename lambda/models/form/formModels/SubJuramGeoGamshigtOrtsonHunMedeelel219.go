package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoGamshigtOrtsonHunMedeelel219 struct {
	GamshigtOrtsonBaidalHunID     *int    `gorm:"column:gamshigt_ortson_baidal_hun_id" json:"gamshigt_ortson_baidal_hun_id"`
	GamshigtOrtsonBaidalMateralID *int    `gorm:"column:gamshigt_ortson_baidal_materal_id" json:"gamshigt_ortson_baidal_materal_id"`
	HbEseh                        *int    `gorm:"column:hb_eseh" json:"hb_eseh"`
	HogjliinBerhsheelID           *int    `gorm:"column:hogjliin_berhsheel_id" json:"hogjliin_berhsheel_id"`
	Huis                          *string `gorm:"column:huis" json:"huis"`
	ID                            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	JiremsenEseh                  *int    `gorm:"column:jiremsen_eseh" json:"jiremsen_eseh"`
	Nas                           *int    `gorm:"column:nas" json:"nas"`
	Ner                           *string `gorm:"column:ner" json:"ner"`
	NiigmiinBaidalID              *int    `gorm:"column:niigmiin_baidal_id" json:"niigmiin_baidal_id"`
	Ovog                          *string `gorm:"column:ovog" json:"ovog"`
	Register                      *string `gorm:"column:register" json:"register"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoGamshigtOrtsonHunMedeelel219) TableName() string {
	return "sub_juram_geo_gamshigt_ortson_hun_medeelel"
}
