package models

import "time"

type ZbNegjKod struct {
	Code      string     `gorm:"column:code" json:"code"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	Icon      string     `gorm:"column:icon" json:"icon"`
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	NegjNer   string     `gorm:"column:negj_ner" json:"negj_ner"`
	TovchNer  string     `gorm:"column:tovch_ner" json:"tovch_ner"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (z *ZbNegjKod) TableName() string {
	return "zb_negj_kod"
}

type SubZbNegjKod struct {
	Code        string `gorm:"column:code" json:"code"`
	ID          int    `gorm:"column:id;primary_key" json:"id"`
	Ner         string `gorm:"column:ner" json:"ner"`
	ZbNegjKodID int    `gorm:"column:zb_negj_kod_id" json:"zb_negj_kod_id"`
}

func (s *SubZbNegjKod) TableName() string {
	return "sub_zb_negj_kod"
}
