package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type HongiinMedee325 struct {
	Aimagid           *int       `gorm:"column:aimagid" json:"aimagid"`
	AjilahHuch        *string    `gorm:"column:ajilah_huch" json:"ajilah_huch"`
	AvsanArgaHemjee   *string    `gorm:"column:avsan_arga_hemjee" json:"avsan_arga_hemjee"`
	AyulDedDedTorolID *int       `gorm:"column:ayul_ded_ded_torol_id" json:"ayul_ded_ded_torol_id"`
	AyulDedTorolID    *int       `gorm:"column:ayul_ded_torol_id" json:"ayul_ded_torol_id"`
	AyulTorolID       *int       `gorm:"column:ayul_torol_id" json:"ayul_torol_id"`
	Bairshil          *string    `gorm:"column:bairshil" json:"bairshil"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	DOgnoo            *time.Time `gorm:"column:d_ognoo" json:"d_ognoo"`
	EOgnoo            *time.Time `gorm:"column:e_ognoo" json:"e_ognoo"`
	ID                int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MAvsanOgnooObeg   *time.Time `gorm:"column:m_avsan_ognoo_obeg" json:"m_avsan_ognoo_obeg"`
	MAvsanOgnooSalbar *time.Time `gorm:"column:m_avsan_ognoo_salbar" json:"m_avsan_ognoo_salbar"`
	Medeelel          *string    `gorm:"column:medeelel" json:"medeelel"`
	NiitHohirol       *string    `gorm:"column:niit_hohirol" json:"niit_hohirol"`
	Orgorog           *string    `gorm:"column:orgorog" json:"orgorog"`
	SumID             *int       `gorm:"column:sum_id" json:"sum_id"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Urtrag            *string    `gorm:"column:urtrag" json:"urtrag"`
	UserID            *int       `gorm:"column:user_id" json:"user_id"`
	Vaktsin           *string    `gorm:"column:vaktsin" json:"vaktsin"`
}

//  TableName sets the insert table name for this struct type
func (h *HongiinMedee325) TableName() string {
	return "hongiin_medee"
}
