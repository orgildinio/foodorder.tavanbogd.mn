package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoNervegdsenAmitan207 struct {
	Bugd                      float32 `gorm:"column:bugd" json:"bugd"`
	DedmalToo                 *int    `gorm:"column:dedmal_too" json:"dedmal_too"`
	ID                        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalToo                    *int    `gorm:"column:mal_too" json:"mal_too"`
	OvchilsonAmitanDedTorolID *int    `gorm:"column:ovchilson_amitan_ded_torol_id" json:"ovchilson_amitan_ded_torol_id"`
	OvchilsonAmitanTorolID    *int    `gorm:"column:ovchilson_amitan_torol_id" json:"ovchilson_amitan_torol_id"`
	OvchilsonDedMalID         *int    `gorm:"column:ovchilson_ded_mal_id" json:"ovchilson_ded_mal_id"`
	OvchilsonMalID            *int    `gorm:"column:ovchilson_mal_id" json:"ovchilson_mal_id"`
	UsID                      *int    `gorm:"column:us_id" json:"us_id"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoNervegdsenAmitan207) TableName() string {
	return "sub_juram_geo_nervegdsen_amitan"
}
