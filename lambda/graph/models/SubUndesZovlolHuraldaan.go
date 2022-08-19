package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubUndesZovlolHuraldaan struct {
	Huraldaan     *string `gorm:"column:huraldaan" json:"huraldaan"`
	HuraldaanEn   *string `gorm:"column:huraldaan_en" json:"huraldaan_en"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndesZovlolID *int    `gorm:"column:undes_zovlol_id" json:"undes_zovlol_id"`
	Zovlomj       *string `gorm:"column:zovlomj" json:"zovlomj"`
}

//  TableName sets the insert table name for this struct type
func (s *SubUndesZovlolHuraldaan) TableName() string {
	return "sub_undes_zovlol_huraldaan"
}
