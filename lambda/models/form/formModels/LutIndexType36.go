package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutIndexType36 struct {
	ErsdeliinIndex *string `gorm:"column:ersdeliin_index" json:"ersdeliin_index"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutIndexType36) TableName() string {
	return "lut_index_type"
}
