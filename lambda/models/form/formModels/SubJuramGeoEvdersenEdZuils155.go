package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoEvdersenEdZuils155 struct {
	Bugd                   *int    `gorm:"column:bugd" json:"bugd"`
	DaatgalEseh            *int    `gorm:"column:daatgal_eseh" json:"daatgal_eseh"`
	EdHorongoDedDedTorolID *int    `gorm:"column:ed_horongo_ded_ded_torol_id" json:"ed_horongo_ded_ded_torol_id"`
	EdHorongoDedTorolID    *int    `gorm:"column:ed_horongo_ded_torol_id" json:"ed_horongo_ded_torol_id"`
	EdHorongoTorolID       *int    `gorm:"column:ed_horongo_torol_id" json:"ed_horongo_torol_id"`
	HohirolID              *int    `gorm:"column:hohirol_id" json:"hohirol_id"`
	ID                     int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Todorkhoilolt          *string `gorm:"column:todorkhoilolt" json:"todorkhoilolt"`
	Too                    float32 `gorm:"column:too" json:"too"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoEvdersenEdZuils155) TableName() string {
	return "sub_juram_geo_evdersen_ed_zuils"
}
