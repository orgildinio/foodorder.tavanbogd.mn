package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type UndesniiZowlol46 struct {
	ButetsBureldhuun *string    `gorm:"column:butets_bureldhuun" json:"butets_bureldhuun"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at"`
	ID               int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SubUndesZovlolHuraldaanUndesniiZowlol46 struct {
	Huraldaan     *string `gorm:"column:huraldaan" json:"huraldaan"`
	HuraldaanEn   *string `gorm:"column:huraldaan_en" json:"huraldaan_en"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UndesZovlolID *int    `gorm:"column:undes_zovlol_id" json:"undes_zovlol_id"`
	Zovlomj       *string `gorm:"column:zovlomj" json:"zovlomj"`
}

func (s *SubUndesZovlolHuraldaanUndesniiZowlol46) TableName() string {
	return "sub_undes_zovlol_huraldaan"
}

type SubUndesZovlolGamshigErsdelBuuruulahUndesniiZowlol46 struct {
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TailbarEn     *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn     *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	TogtoolFile   *string `gorm:"column:togtool_file" json:"togtool_file"`
	UndesZovlolID *int    `gorm:"column:undes_zovlol_id" json:"undes_zovlol_id"`
}

func (s *SubUndesZovlolGamshigErsdelBuuruulahUndesniiZowlol46) TableName() string {
	return "sub_undes_zovlol_gamshig_ersdel_buuruulah"
}

//  TableName sets the insert table name for this struct type
func (u *UndesniiZowlol46) TableName() string {
	return "undesnii_zowlol"
}
