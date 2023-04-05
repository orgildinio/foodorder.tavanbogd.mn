package models

import "time"

type TblArt struct {
	AdultEseh            string     `gorm:"column:adult_eseh" json:"adult_eseh"`
	ArtBatalgaajiltTuluv int        `gorm:"column:art_batalgaajilt_tuluv" json:"art_batalgaajilt_tuluv"`
	ArtNer               string     `gorm:"column:art_ner" json:"art_ner"`
	ArtistID             int        `gorm:"column:artist_id" json:"artist_id"`
	Content              string     `gorm:"column:content" json:"content"`
	Cover                string     `gorm:"column:cover" json:"cover"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedDate          string     `gorm:"column:created_date" json:"created_date"`
	Duu                  string     `gorm:"column:duu" json:"duu"`
	File                 string     `gorm:"column:file" json:"file"`
	HurZurag             string     `gorm:"column:hur_zurag" json:"hur_zurag"`
	ID                   int        `gorm:"column:id;primary_key" json:"id"`
	Image                string     `gorm:"column:image" json:"image"`
	PatentNumber         string     `gorm:"column:patent_number" json:"patent_number"`
	PostStatusID         int        `gorm:"column:post_status_id" json:"post_status_id"`
	Tailbar              string     `gorm:"column:tailbar" json:"tailbar"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID               int        `gorm:"column:user_id" json:"user_id"`
	Video                string     `gorm:"column:video" json:"video"`
	YamBatlahEseh        string     `gorm:"column:yam_batlah_eseh" json:"yam_batlah_eseh"`
}

func (t *TblArt) TableName() string {
	return "tbl_art"
}
