package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutDedEdHorongo179 struct {
	DedEdHorongo     *string `gorm:"column:ded_ed_horongo" json:"ded_ed_horongo"`
	EdHorongoTorolID *int    `gorm:"column:ed_horongo_torol_id" json:"ed_horongo_torol_id"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutDedEdHorongo179) TableName() string {
	return "lut_ded_ed_horongo"
}
