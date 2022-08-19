package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDedNervegdsenHun140 struct {
	DedNervegdsenHun *string `gorm:"column:ded_nervegdsen_hun" json:"ded_nervegdsen_hun"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NervegdsenHunID  *int    `gorm:"column:nervegdsen_hun_id" json:"nervegdsen_hun_id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDedNervegdsenHun140) TableName() string {
	return "lut_ded_nervegdsen_hun"
}
