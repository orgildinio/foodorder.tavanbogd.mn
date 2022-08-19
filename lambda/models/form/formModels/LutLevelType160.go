package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutLevelType160 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LevelType *string `gorm:"column:level_type" json:"level_type"`
}

//  TableName sets the insert table name for this struct type
func (l *LutLevelType160) TableName() string {
	return "lut_level_type"
}
