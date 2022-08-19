package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubDugnelt77 struct {
	DangerTypeID *int    `gorm:"column:danger_type_id" json:"danger_type_id"`
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ZHeregjilt   *string `gorm:"column:z_heregjilt" json:"z_heregjilt"`
	Zowlomj      *string `gorm:"column:zowlomj" json:"zowlomj"`
}

//  TableName sets the insert table name for this struct type
func (s *SubDugnelt77) TableName() string {
	return "sub_dugnelt"
}
