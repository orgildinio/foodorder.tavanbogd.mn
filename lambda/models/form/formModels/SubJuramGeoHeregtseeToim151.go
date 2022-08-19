package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoHeregtseeToim151 struct {
	HemjihNegjID         *int    `gorm:"column:hemjih_negj_id" json:"hemjih_negj_id"`
	HeregtseeDDedTorolID *int    `gorm:"column:heregtsee_d_ded_torol_id" json:"heregtsee_d_ded_torol_id"`
	HeregtseeDedTorolID  *int    `gorm:"column:heregtsee_ded_torol_id" json:"heregtsee_ded_torol_id"`
	HeregtseeTorolID     *int    `gorm:"column:heregtsee_torol_id" json:"heregtsee_torol_id"`
	ID                   int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Too                  float32 `gorm:"column:too" json:"too"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoHeregtseeToim151) TableName() string {
	return "sub_juram_geo_heregtsee_toim"
}
