package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type MergejliinZowlol49 struct {
	BureldehuunMn *string    `gorm:"column:bureldehuun_mn" json:"bureldehuun_mn"`
	Conference    *string    `gorm:"column:conference" json:"conference"`
	CreatedAt     *time.Time `gorm:"column:created_at" json:"created_at"`
	Delegation    *string    `gorm:"column:delegation" json:"delegation"`
	Huraldaan     *string    `gorm:"column:huraldaan" json:"huraldaan"`
	ID            int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UpdatedAt     *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SubMergejilZovolgooTolvolgooMergejliinZowlol49 struct {
	Description *string `gorm:"column:description" json:"description"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MZovlolID   *int    `gorm:"column:m_zovlol_id" json:"m_zovlol_id"`
	Plan        *string `gorm:"column:plan" json:"plan"`
	Tailbar     *string `gorm:"column:tailbar" json:"tailbar"`
	Tolvolgoo   *string `gorm:"column:tolvolgoo" json:"tolvolgoo"`
}

func (s *SubMergejilZovolgooTolvolgooMergejliinZowlol49) TableName() string {
	return "sub_mergejil_zovolgoo_tolvolgoo"
}

type SubMergejilZovlolHuraldaanMergejliinZowlol49 struct {
	HuraldaanEn *string `gorm:"column:huraldaan_en" json:"huraldaan_en"`
	HuraldaanMn *string `gorm:"column:huraldaan_mn" json:"huraldaan_mn"`
	ID          int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MZovlolID   *int    `gorm:"column:m_zovlol_id" json:"m_zovlol_id"`
	ZovlomjMn   *string `gorm:"column:zovlomj_mn" json:"zovlomj_mn"`
}

func (s *SubMergejilZovlolHuraldaanMergejliinZowlol49) TableName() string {
	return "sub_mergejil_zovlol_huraldaan"
}

type SubMergejilZovlolTushaalMergejliinZowlol49 struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MZovlolID *int    `gorm:"column:m_zovlol_id" json:"m_zovlol_id"`
	TailbarEn *string `gorm:"column:tailbar_en" json:"tailbar_en"`
	TailbarMn *string `gorm:"column:tailbar_mn" json:"tailbar_mn"`
	Tushaal   *string `gorm:"column:tushaal" json:"tushaal"`
}

func (s *SubMergejilZovlolTushaalMergejliinZowlol49) TableName() string {
	return "sub_mergejil_zovlol_tushaal"
}

//  TableName sets the insert table name for this struct type
func (m *MergejliinZowlol49) TableName() string {
	return "mergejliin_zowlol"
}
