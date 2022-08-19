package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoUchirsanBoditHorhirol150 struct {
	HemjihNegjID       *int `gorm:"column:hemjih_negj_id" json:"hemjih_negj_id"`
	HohirolBaidalID    *int `gorm:"column:hohirol_baidal_id" json:"hohirol_baidal_id"`
	HohirolDDedTorolID *int `gorm:"column:hohirol_d_ded_torol_id" json:"hohirol_d_ded_torol_id"`
	HohirolDedTorolID  *int `gorm:"column:hohirol_ded_torol_id" json:"hohirol_ded_torol_id"`
	HohirolToo         *int `gorm:"column:hohirol_too" json:"hohirol_too"`
	HohirolTorolID     *int `gorm:"column:hohirol_torol_id" json:"hohirol_torol_id"`
	HorogdsonBaidalID  *int `gorm:"column:horogdson_baidal_id" json:"horogdson_baidal_id"`
	ID                 int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Too                *int `gorm:"column:too" json:"too"`
	UchirsanHohirol    *int `gorm:"column:uchirsan_hohirol" json:"uchirsan_hohirol"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoUchirsanBoditHorhirol150) TableName() string {
	return "sub_juram_geo_uchirsan_bodit_horhirol"
}
