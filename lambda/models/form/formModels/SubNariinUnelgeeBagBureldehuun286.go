package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type SubNariinUnelgeeBagBureldehuun286 struct {
	AlbanTushaal        *string `gorm:"column:alban_tushaal" json:"alban_tushaal"`
	Baiguullaga         *string `gorm:"column:baiguullaga" json:"baiguullaga"`
	Bolovsrol           *string `gorm:"column:bolovsrol" json:"bolovsrol"`
	GamshigSudalgaaEseh *int    `gorm:"column:gamshig_sudalgaa_eseh" json:"gamshig_sudalgaa_eseh"`
	ID                  int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Mail                *string `gorm:"column:mail" json:"mail"`
	Mergejil            *string `gorm:"column:mergejil" json:"mergejil"`
	Ner                 *string `gorm:"column:ner" json:"ner"`
	Ovog                *string `gorm:"column:ovog" json:"ovog"`
	Register            *string `gorm:"column:register" json:"register"`
	Turshilga           *string `gorm:"column:turshilga" json:"turshilga"`
	Utas                *int    `gorm:"column:utas" json:"utas"`
}

//  TableName sets the insert table name for this struct type
func (s *SubNariinUnelgeeBagBureldehuun286) TableName() string {
	return "sub_nariin_unelgee_bag_bureldehuun"
}
