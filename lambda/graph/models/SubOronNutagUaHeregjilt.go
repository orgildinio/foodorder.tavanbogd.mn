package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubOronNutagUaHeregjilt struct {
	HuviMn    *int    `gorm:"column:huvi_mn" json:"huvi_mn"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int    `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	Tailbar   *string `gorm:"column:tailbar" json:"tailbar"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
}

//  TableName sets the insert table name for this struct type
func (s *SubOronNutagUaHeregjilt) TableName() string {
	return "sub_oron_nutag_ua_heregjilt"
}
