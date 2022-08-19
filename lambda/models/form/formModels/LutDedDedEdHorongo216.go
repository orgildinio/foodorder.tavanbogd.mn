package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDedDedEdHorongo216 struct {
	DedDedHorongo  *string `gorm:"column:ded_ded_horongo" json:"ded_ded_horongo"`
	DedEdHorongoID *int    `gorm:"column:ded_ed_horongo_id" json:"ded_ed_horongo_id"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDedDedEdHorongo216) TableName() string {
	return "lut_ded_ded_ed_horongo"
}
