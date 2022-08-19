package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewUndesniiZovlol struct {
	CreatedAt                            *time.Time                              `gorm:"column:created_at" json:"created_at"`
	DeletedAt                            *time.Time                              `gorm:"column:deleted_at" json:"deleted_at"`
	Huraldaan                            *string                                 `gorm:"column:huraldaan" json:"huraldaan"`
	ID                                   *int                                    `gorm:"column:id" json:"id"`
	TailbarMnBureldehuun                 *string                                 `gorm:"column:tailbar_mn_bureldehuun" json:"tailbar_mn_bureldehuun"`
	UpdatedAt                            *time.Time                              `gorm:"column:updated_at" json:"updated_at"`
	SubUndesZovlolGamshigErsdelBuuruulah []*SubUndesZovlolGamshigErsdelBuuruulah `gorm:"-:all"`
	SubUndesZovlolHuraldaan              []*SubUndesZovlolHuraldaan              `gorm:"-:all"`
}

//  TableName sets the insert table name for this struct type
func (v *ViewUndesniiZovlol) TableName() string {
	return "view_undesnii_zovlol"
}
