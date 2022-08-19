package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutUilAjilgaaniiChiglelAan297 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UilAjilgaaniiChiglel *string `gorm:"column:uil_ajilgaanii_chiglel" json:"uil_ajilgaanii_chiglel"`
}

//  TableName sets the insert table name for this struct type
func (l *LutUilAjilgaaniiChiglelAan297) TableName() string {
	return "lut_uil_ajilgaanii_chiglel_aan"
}
