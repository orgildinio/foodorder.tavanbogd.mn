package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSocialType4 struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SocialType *string `gorm:"column:social_type" json:"social_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSocialType4) TableName() string {
	return "lut_social_type"
}
