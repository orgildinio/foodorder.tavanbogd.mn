package models

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

type Barilga struct {
	AimagID        int        `gorm:"column:aimag_id" json:"aimag_id"`
	AngilalID      int        `gorm:"column:angilal_id" json:"angilal_id"`
	AshiglaltOgnoo DB.Date    `gorm:"column:ashiglalt_ognoo" json:"ashiglalt_ognoo"`
	BagID          int        `gorm:"column:bag_id" json:"bag_id"`
	Bairshil       string     `gorm:"column:bairshil" json:"bairshil"`
	ButetsID       int        `gorm:"column:butets_id" json:"butets_id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Dun            float32    `gorm:"column:dun" json:"dun"`
	GisCatID       int        `gorm:"column:gis_cat_id" json:"gis_cat_id"`
	GisLayerID     int        `gorm:"column:gis_layer_id" json:"gis_layer_id"`
	HEhUusverDedID int        `gorm:"column:h_eh_uusver_ded_id" json:"h_eh_uusver_ded_id"`
	HEhUusverID    int        `gorm:"column:h_eh_uusver_id" json:"h_eh_uusver_id"`
	HanaID         int        `gorm:"column:hana_id" json:"hana_id"`
	Hayg           string     `gorm:"column:hayg" json:"hayg"`
	HemjihNegjID   int        `gorm:"column:hemjih_negj_id" json:"hemjih_negj_id"`
	HuchiltID      int        `gorm:"column:huchilt_id" json:"huchilt_id"`
	ID             int        `gorm:"column:id;primary_key" json:"id"`
	NegjUne        float32    `gorm:"column:negj_une" json:"negj_une"`
	SumID          int        `gorm:"column:sum_id" json:"sum_id"`
	Talbai         string     `gorm:"column:talbai" json:"talbai"`
	TolovID        int        `gorm:"column:tolov_id" json:"tolov_id"`
	Too            float32    `gorm:"column:too" json:"too"`
	UHCode         string     `gorm:"column:u_h_code" json:"u_h_code"`
	UHNer          string     `gorm:"column:u_h_ner" json:"u_h_ner"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID         int        `gorm:"column:user_id" json:"user_id"`
	ZbNegjCodeID   int        `gorm:"column:zb_negj_code_id" json:"zb_negj_code_id"`
}

func (b *Barilga) TableName() string {
	return "barilga"
}
