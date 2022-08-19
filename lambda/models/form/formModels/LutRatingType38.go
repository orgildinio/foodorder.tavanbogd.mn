package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutRatingType38 struct {
	ErsdeliinZereglel *string `gorm:"column:ersdeliin_zereglel" json:"ersdeliin_zereglel"`
	ID                int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutRatingType38) TableName() string {
	return "lut_rating_type"
}
