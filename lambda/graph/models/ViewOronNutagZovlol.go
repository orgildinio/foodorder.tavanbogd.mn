package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewOronNutagZovlol struct {
	AimagID                   *int                         `gorm:"column:aimag_id" json:"aimag_id"`
	Aimagname                 *string                      `gorm:"column:aimagname" json:"aimagname"`
	CreatedAt                 *time.Time                   `gorm:"column:created_at" json:"created_at"`
	DeletedAt                 *time.Time                   `gorm:"column:deleted_at" json:"deleted_at"`
	ID                        *int                         `gorm:"column:id" json:"id"`
	SumID                     *int                         `gorm:"column:sum_id" json:"sum_id"`
	Sumname                   *string                      `gorm:"column:sumname" json:"sumname"`
	UpdatedAt                 *time.Time                   `gorm:"column:updated_at" json:"updated_at"`
	SubOronNutagGamshigErsdel []*SubOronNutagGamshigErsdel `gorm:"-:all"`
	SubOronNutagTogtool       []*SubOronNutagTogtool       `gorm:"-:all"`
	SubOronNutagUaHeregjilt   []*SubOronNutagUaHeregjilt   `gorm:"-:all"`
	SubOronNutagUaTolvolgoo   []*SubOronNutagUaTolvolgoo   `gorm:"-:all"`
}

//  TableName sets the insert table name for this struct type
func (v *ViewOronNutagZovlol) TableName() string {
	return "view_oron_nutag_zovlol"
}
