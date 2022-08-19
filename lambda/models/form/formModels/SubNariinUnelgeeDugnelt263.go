package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubNariinUnelgeeDugnelt263 struct {
	AyulTorolID      *int    `gorm:"column:ayul_torol_id" json:"ayul_torol_id"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Zovlomj          *string `gorm:"column:zovlomj" json:"zovlomj"`
	ZovlomjHeregjilt float32 `gorm:"column:zovlomj_heregjilt" json:"zovlomj_heregjilt"`
}

//  TableName sets the insert table name for this struct type
func (s *SubNariinUnelgeeDugnelt263) TableName() string {
	return "sub_nariin_unelgee_dugnelt"
}
