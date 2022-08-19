package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoTejeeverAmitan154 struct {
	HorogdsonTejeeverAmitanID *int `gorm:"column:horogdson_tejeever_amitan_id" json:"horogdson_tejeever_amitan_id"`
	ID                        int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NiitToo                   *int `gorm:"column:niit_too" json:"niit_too"`
	TejeeverAmitanTorolID     *int `gorm:"column:tejeever_amitan_torol_id" json:"tejeever_amitan_torol_id"`
	Too                       *int `gorm:"column:too" json:"too"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoTejeeverAmitan154) TableName() string {
	return "sub_juram_geo_tejeever_amitan"
}
