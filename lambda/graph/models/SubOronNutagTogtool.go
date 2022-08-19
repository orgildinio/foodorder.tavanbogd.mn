package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubOronNutagTogtool struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int    `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	Togtool   *string `gorm:"column:togtool" json:"togtool"`
}

//  TableName sets the insert table name for this struct type
func (s *SubOronNutagTogtool) TableName() string {
	return "sub_oron_nutag_togtool"
}
