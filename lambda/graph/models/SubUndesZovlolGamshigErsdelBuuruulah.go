package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubUndesZovlolGamshigErsdelBuuruulah struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TailbarEn     *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn     *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	TogtoolFile   *string `gorm:"column:togtool_file" json:"togtool_file"`
	UndesZovlolID *int    `gorm:"column:undes_zovlol_id" json:"undes_zovlol_id"`
}

//  TableName sets the insert table name for this struct type
func (s *SubUndesZovlolGamshigErsdelBuuruulah) TableName() string {
	return "sub_undes_zovlol_gamshig_ersdel_buuruulah"
}
