package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SendainUa52 struct {
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UaHuree   *string    `gorm:"column:ua_huree" json:"ua_huree"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//  TableName sets the insert table name for this struct type
func (s *SendainUa52) TableName() string {
	return "sendain_ua"
}
