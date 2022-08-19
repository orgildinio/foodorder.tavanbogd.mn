package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutTejeeverAmitniiTurul177 struct {
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TejeeverAmitniiTurul *string `gorm:"column:tejeever_amitnii_turul" json:"tejeever_amitnii_turul"`
}

//  TableName sets the insert table name for this struct type
func (l *LutTejeeverAmitniiTurul177) TableName() string {
	return "lut_tejeever_amitnii_turul"
}
