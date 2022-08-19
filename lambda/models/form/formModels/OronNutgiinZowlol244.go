package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type OronNutgiinZowlol244 struct {
	AimagID   *int       `gorm:"column:aimag_id" json:"aimag_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	ID        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SumID     *int       `gorm:"column:sum_id" json:"sum_id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SubOronNutagTogtoolOronNutgiinZowlol244 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int    `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	Togtool   *string `gorm:"column:togtool" json:"togtool"`
}

func (s *SubOronNutagTogtoolOronNutgiinZowlol244) TableName() string {
	return "sub_oron_nutag_togtool"
}

type SubOronNutagUaTolvolgooOronNutgiinZowlol244 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int    `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	Tolvolgoo *string `gorm:"column:tolvolgoo" json:"tolvolgoo"`
}

func (s *SubOronNutagUaTolvolgooOronNutgiinZowlol244) TableName() string {
	return "sub_oron_nutag_ua_tolvolgoo"
}

type SubOronNutagUaHeregjiltOronNutgiinZowlol244 struct {
	HuviMn    *int    `gorm:"column:huvi_mn" json:"huvi_mn"`
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int    `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	Tailbar   *string `gorm:"column:tailbar" json:"tailbar"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
}

func (s *SubOronNutagUaHeregjiltOronNutgiinZowlol244) TableName() string {
	return "sub_oron_nutag_ua_heregjilt"
}

type SubOronNutagGamshigErsdelOronNutgiinZowlol244 struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OZovlolID *int     `gorm:"column:o_zovlol_id" json:"o_zovlol_id"`
	Ognoo     *DB.Date `gorm:"column:ognoo" json:"ognoo"`
	Tosov     *string  `gorm:"column:tosov" json:"tosov"`
}

func (s *SubOronNutagGamshigErsdelOronNutgiinZowlol244) TableName() string {
	return "sub_oron_nutag_gamshig_ersdel"
}

//  TableName sets the insert table name for this struct type
func (o *OronNutgiinZowlol244) TableName() string {
	return "oron_nutgiin_zowlol"
}
