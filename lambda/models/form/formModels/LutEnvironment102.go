package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEnvironment102 struct {
	Environment *string `gorm:"column:environment" json:"environment"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutEnvironment102) TableName() string {
	return "lut_environment"
}
