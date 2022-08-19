package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoAvrahSergeeh149 struct {
	BaiguullagaTorolID *int    `gorm:"column:baiguullaga_torol_id" json:"baiguullaga_torol_id"`
	ID                 int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Mongo              float32 `gorm:"column:mongo" json:"mongo"`
	NiitToo            float32 `gorm:"column:niit_too" json:"niit_too"`
	Too                float32 `gorm:"column:too" json:"too"`
	UndsenDedTorolID   *int    `gorm:"column:undsen_ded_torol_id" json:"undsen_ded_torol_id"`
	UndsenTorolID      *int    `gorm:"column:undsen_torol_id" json:"undsen_torol_id"`
}

type SubBaiguullagaTorolSubJuramGeoAvrahSergeeh149 struct {
	Bugd                   *int `gorm:"column:bugd" json:"bugd"`
	DedUndsenTorolID       *int `gorm:"column:ded_undsen_torol_id" json:"ded_undsen_torol_id"`
	ID                     int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	JuramGeoAvrahSergeehID *int `gorm:"column:juram_geo_avrah_sergeeh_id" json:"juram_geo_avrah_sergeeh_id"`
	Too                    *int `gorm:"column:too" json:"too"`
	UndsenTorolID          *int `gorm:"column:undsen_torol_id" json:"undsen_torol_id"`
}

func (s *SubBaiguullagaTorolSubJuramGeoAvrahSergeeh149) TableName() string {
	return "sub_baiguullaga_torol"
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoAvrahSergeeh149) TableName() string {
	return "sub_juram_geo_avrah_sergeeh"
}
