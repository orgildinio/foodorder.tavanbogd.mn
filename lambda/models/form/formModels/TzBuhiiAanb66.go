package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type TzBuhiiAanb66 struct {
	ActivityTypeID    *int       `gorm:"column:activity_type_id" json:"activity_type_id"`
	AimagID           *int       `gorm:"column:aimag_id" json:"aimag_id"`
	BNer              *string    `gorm:"column:b_ner" json:"b_ner"`
	BagID             *int       `gorm:"column:bag_id" json:"bag_id"`
	Bair              *string    `gorm:"column:bair" json:"bair"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	Email             *string    `gorm:"column:email" json:"email"`
	EvalutionSectorID *int       `gorm:"column:evalution_sector_id" json:"evalution_sector_id"`
	Gudamj            *string    `gorm:"column:gudamj" json:"gudamj"`
	HayagDelegrengui  *string    `gorm:"column:hayag_delegrengui" json:"hayag_delegrengui"`
	ID                int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Register          *string    `gorm:"column:register" json:"register"`
	ShiidwerHawsralt  *string    `gorm:"column:shiidwer_hawsralt" json:"shiidwer_hawsralt"`
	SumID             *int       `gorm:"column:sum_id" json:"sum_id"`
	TzDugaar          *string    `gorm:"column:tz_dugaar" json:"tz_dugaar"`
	TzDuusah          *DB.Date   `gorm:"column:tz_duusah" json:"tz_duusah"`
	TzOlgoson         *DB.Date   `gorm:"column:tz_olgoson" json:"tz_olgoson"`
	TzShiidwer        *string    `gorm:"column:tz_shiidwer" json:"tz_shiidwer"`
	UbGerchilgee      *string    `gorm:"column:ub_gerchilgee" json:"ub_gerchilgee"`
	UbGerchilgeeDug   *string    `gorm:"column:ub_gerchilgee_dug" json:"ub_gerchilgee_dug"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Utas              *string    `gorm:"column:utas" json:"utas"`
	Zahiral           *string    `gorm:"column:zahiral" json:"zahiral"`
	ZakiUtas          *string    `gorm:"column:zaki_utas" json:"zaki_utas"`
}

type SubTzAanUnelgeeSalbarTzBuhiiAanb66 struct {
	AanID           *int `gorm:"column:aan_id" json:"aan_id"`
	ID              int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UnelgeeSalbarID *int `gorm:"column:unelgee_salbar_id" json:"unelgee_salbar_id"`
}

func (s *SubTzAanUnelgeeSalbarTzBuhiiAanb66) TableName() string {
	return "sub_tz_aan_unelgee_salbar"
}

//  TableName sets the insert table name for this struct type
func (t *TzBuhiiAanb66) TableName() string {
	return "tz_buhii_aanb"
}
