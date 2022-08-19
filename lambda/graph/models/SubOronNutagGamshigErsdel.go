package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubOronNutagGamshigErsdel struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int     `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	Ognoo     *DB.Date `gorm:"column:ognoo" json:"ognoo"`
	Tosov     *string  `gorm:"column:tosov" json:"tosov"`
}

//  TableName sets the insert table name for this struct type
func (s *SubOronNutagGamshigErsdel) TableName() string {
	return "sub_oron_nutag_gamshig_ersdel"
}
