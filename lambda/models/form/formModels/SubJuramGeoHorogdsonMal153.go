package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubJuramGeoHorogdsonMal153 struct {
	AvsanArgaHemjeeID       *int    `gorm:"column:avsan_arga_hemjee_id" json:"avsan_arga_hemjee_id"`
	Huis                    *string `gorm:"column:huis" json:"huis"`
	HuisToo                 *int    `gorm:"column:huis_too" json:"huis_too"`
	ID                      int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MalDDedTorolID          *int    `gorm:"column:mal_d_ded_torol_id" json:"mal_d_ded_torol_id"`
	MalDedTorolID           *int    `gorm:"column:mal_ded_torol_id" json:"mal_ded_torol_id"`
	MalToo                  *int    `gorm:"column:mal_too" json:"mal_too"`
	MalTorolID              *int    `gorm:"column:mal_torol_id" json:"mal_torol_id"`
	NiitHorogdsonMal        float32 `gorm:"column:niit_horogdson_mal" json:"niit_horogdson_mal"`
	NiitHorogdsonMalTorloor *int    `gorm:"column:niit_horogdson_mal_torloor" json:"niit_horogdson_mal_torloor"`
	NiitOvchilsonMal        float32 `gorm:"column:niit_ovchilson_mal" json:"niit_ovchilson_mal"`
	NiitOvchilsonMalTorloor *int    `gorm:"column:niit_ovchilson_mal_torloor" json:"niit_ovchilson_mal_torloor"`
}

//  TableName sets the insert table name for this struct type
func (s *SubJuramGeoHorogdsonMal153) TableName() string {
	return "sub_juram_geo_horogdson_mal"
}
