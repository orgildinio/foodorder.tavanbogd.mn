package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubSendainUaZorilt247 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	ZoriltEn  *string `gorm:"column:zorilt_en" json:"zorilt_en"`
	ZoriltMn  *string `gorm:"column:zorilt_mn" json:"zorilt_mn"`
}

//  TableName sets the insert table name for this struct type
func (s *SubSendainUaZorilt247) TableName() string {
	return "sub_sendain_ua_zorilt"
}
