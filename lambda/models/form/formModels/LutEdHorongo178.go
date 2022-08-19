package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutEdHorongo178 struct {
	EdHorongo *string `gorm:"column:ed_horongo" json:"ed_horongo"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutEdHorongo178) TableName() string {
	return "lut_ed_horongo"
}
