package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type TzHugtsaaSungah72 struct {
	BID                    *int       `gorm:"column:b_id" json:"b_id"`
	ConfirmStatusID        *int       `gorm:"column:confirm_status_id" json:"confirm_status_id"`
	ConfirmStatusIDByAdmin *int       `gorm:"column:confirm_status_id_by_admin" json:"confirm_status_id_by_admin"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt              *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	HuchinteiOgnoo         *DB.Date   `gorm:"column:huchintei_ognoo" json:"huchintei_ognoo"`
	ID                     int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Shiidwer               *string    `gorm:"column:shiidwer" json:"shiidwer"`
	ShiidwerHawsralt       *string    `gorm:"column:shiidwer_hawsralt" json:"shiidwer_hawsralt"`
	SungahEseh             *string    `gorm:"column:sungah_eseh" json:"sungah_eseh"`
	SungahShaltgaan        *string    `gorm:"column:sungah_shaltgaan" json:"sungah_shaltgaan"`
	SungasanOgnoo          *DB.Date   `gorm:"column:sungasan_ognoo" json:"sungasan_ognoo"`
	Tailbar                *string    `gorm:"column:tailbar" json:"tailbar"`
	TzGerchilgee           *string    `gorm:"column:tz_gerchilgee" json:"tz_gerchilgee"`
	TzGerchilgeeDugaar     *string    `gorm:"column:tz_gerchilgee_dugaar" json:"tz_gerchilgee_dugaar"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SubAttachFileTzHugtsaaSungah72 struct {
	FileNer    *string `gorm:"column:file_ner" json:"file_ner"`
	Hawsralt   *string `gorm:"column:hawsralt" json:"hawsralt"`
	ID         int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TzSungahID *int    `gorm:"column:tz_sungah_id" json:"tz_sungah_id"`
}

func (s *SubAttachFileTzHugtsaaSungah72) TableName() string {
	return "sub_attach_file"
}

//  TableName sets the insert table name for this struct type
func (t *TzHugtsaaSungah72) TableName() string {
	return "tz_hugtsaa_sungah"
}
