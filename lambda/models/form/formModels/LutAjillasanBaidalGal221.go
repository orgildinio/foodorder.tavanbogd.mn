package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutAjillasanBaidalGal221 struct {
	Ajillasan_baidalGal *string `gorm:"column:ajillasan baidal_gal" json:"ajillasan baidal_gal"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

//  TableName sets the insert table name for this struct type
func (l *LutAjillasanBaidalGal221) TableName() string {
	return "lut_ajillasan_baidal_gal"
}
