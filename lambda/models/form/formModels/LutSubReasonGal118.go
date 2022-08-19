package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutSubReasonGal118 struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanGalID *int    `gorm:"column:shaltgaan_gal_id" json:"shaltgaan_gal_id"`
	SubReasonGal   *string `gorm:"column:sub_reason_gal" json:"sub_reason_gal"`
}

//  TableName sets the insert table name for this struct type
func (l *LutSubReasonGal118) TableName() string {
	return "lut_sub_reason_gal"
}
