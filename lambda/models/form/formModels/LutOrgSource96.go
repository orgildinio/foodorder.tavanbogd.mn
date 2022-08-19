package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type LutOrgSource96 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OrgSource *string `gorm:"column:org_source" json:"org_source"`
}

//  TableName sets the insert table name for this struct type
func (l *LutOrgSource96) TableName() string {
	return "lut_org_source"
}
