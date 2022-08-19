package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutShaltgaanGal116 struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ShaltgaanGal *string `gorm:"column:shaltgaan_gal" json:"shaltgaan_gal"`
}

//  TableName sets the insert table name for this struct type
func (l *LutShaltgaanGal116) TableName() string {
	return "lut_shaltgaan_gal"
}
