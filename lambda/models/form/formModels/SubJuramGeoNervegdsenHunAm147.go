package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoNervegdsenHunAm147 struct {
	Bugd                       *int `gorm:"column:bugd" json:"bugd"`
	HunAmDedDedTorolID         *int `gorm:"column:hun_am_ded_ded_torol_id" json:"hun_am_ded_ded_torol_id"`
	HunAmDedTorolID            *int `gorm:"column:hun_am_ded_torol_id" json:"hun_am_ded_torol_id"`
	HunAmToo                   *int `gorm:"column:hun_am_too" json:"hun_am_too"`
	HunAmTorolID               *int `gorm:"column:hun_am_torol_id" json:"hun_am_torol_id"`
	HunamdedToo                *int `gorm:"column:hunamded_too" json:"hunamded_too"`
	ID                         int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NervegdesenHunID           *int `gorm:"column:nervegdesen_hun_id" json:"nervegdesen_hun_id"`
	NervegdsenBaiguullagaID    *int `gorm:"column:nervegdsen_baiguullaga_id" json:"nervegdsen_baiguullaga_id"`
	NervegdsenDedBaiguullagaID *int `gorm:"column:nervegdsen_ded_baiguullaga_id" json:"nervegdsen_ded_baiguullaga_id"`
	Too                        *int `gorm:"column:too" json:"too"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoNervegdsenHunAm147) TableName() string {
	return "sub_juram_geo_nervegdsen_hun_am"
}
