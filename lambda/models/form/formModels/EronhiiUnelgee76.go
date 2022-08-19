package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type EronhiiUnelgee76 struct {
	ATushaal               *string    `gorm:"column:a_tushaal" json:"a_tushaal"`
	AimagID                *int       `gorm:"column:aimag_id" json:"aimag_id"`
	BagID                  *int       `gorm:"column:bag_id" json:"bag_id"`
	Baiguullaga            *string    `gorm:"column:baiguullaga" json:"baiguullaga"`
	Bairshil               *string    `gorm:"column:bairshil" json:"bairshil"`
	ConfirmStatusID        *int       `gorm:"column:confirm_status_id" json:"confirm_status_id"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	DangerTypeID           *int       `gorm:"column:danger_type_id" json:"danger_type_id"`
	EvalutionSectorID      *int       `gorm:"column:evalution_sector_id" json:"evalution_sector_id"`
	GeneralEvalutionTypeID *int       `gorm:"column:general_evalution_type_id" json:"general_evalution_type_id"`
	ID                     int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IndicatorID            *int       `gorm:"column:indicator_id" json:"indicator_id"`
	IndicatorTypeID        *int       `gorm:"column:indicator_type_id" json:"indicator_type_id"`
	Ner                    *string    `gorm:"column:ner" json:"ner"`
	Ognoo                  *DB.Date   `gorm:"column:ognoo" json:"ognoo"`
	Owog                   *string    `gorm:"column:owog" json:"owog"`
	SubIndicatorID         *int       `gorm:"column:sub_indicator_id" json:"sub_indicator_id"`
	SumID                  *int       `gorm:"column:sum_id" json:"sum_id"`
	UHiijDuussanOgnoo      *DB.Date   `gorm:"column:u_hiij_duussan_ognoo" json:"u_hiij_duussan_ognoo"`
	UHiijEhelsenOgnoo      *DB.Date   `gorm:"column:u_hiij_ehelsen_ognoo" json:"u_hiij_ehelsen_ognoo"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UrDun                  *string    `gorm:"column:ur_dun" json:"ur_dun"`
	UserID                 *int       `gorm:"column:user_id" json:"user_id"`
	Zurag                  *string    `gorm:"column:zurag" json:"zurag"`
}

type SubErsdelUnelgeeEronhiiUnelgee76 struct {
	DangerTypeID  *int    `gorm:"column:danger_type_id" json:"danger_type_id"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Jinlelt       *string `gorm:"column:jinlelt" json:"jinlelt"`
	UnelgeeFormID *int    `gorm:"column:unelgee_form_id" json:"unelgee_form_id"`
	Zereglel      *string `gorm:"column:zereglel" json:"zereglel"`
}

func (s *SubErsdelUnelgeeEronhiiUnelgee76) TableName() string {
	return "sub_ersdel_unelgee"
}

type SubErsdelIndexEronhiiUnelgee76 struct {
	DangerTypeID   *int    `gorm:"column:danger_type_id" json:"danger_type_id"`
	DisIndicatorID *int    `gorm:"column:dis_indicator_id" json:"dis_indicator_id"`
	Ersdel         *string `gorm:"column:ersdel" json:"ersdel"`
	ID             int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IndicatorID    *int    `gorm:"column:indicator_id" json:"indicator_id"`
	LevelTypeID    *int    `gorm:"column:level_type_id" json:"level_type_id"`
	RatingTypeID   *int    `gorm:"column:rating_type_id" json:"rating_type_id"`
	SubIndicatorID *int    `gorm:"column:sub_indicator_id" json:"sub_indicator_id"`
	UnelgeeFormID  *int    `gorm:"column:unelgee_form_id" json:"unelgee_form_id"`
}

func (s *SubErsdelIndexEronhiiUnelgee76) TableName() string {
	return "sub_ersdel_index"
}

type SubHawsralBbEronhiiUnelgee76 struct {
	BbNer         *string `gorm:"column:bb_ner" json:"bb_ner"`
	Hawsralt      *string `gorm:"column:hawsralt" json:"hawsralt"`
	ID            int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UnelgeeFormID *int    `gorm:"column:unelgee_form_id" json:"unelgee_form_id"`
}

func (s *SubHawsralBbEronhiiUnelgee76) TableName() string {
	return "sub_hawsral_bb"
}

type SubYronhiiunelgeeBagEronhiiUnelgee76 struct {
	BagAlbanTushaal  *string `gorm:"column:bag_alban_tushaal" json:"bag_alban_tushaal"`
	BagBaiguullaga   *string `gorm:"column:bag_baiguullaga" json:"bag_baiguullaga"`
	BagNer           *string `gorm:"column:bag_ner" json:"bag_ner"`
	BagOvog          *string `gorm:"column:bag_ovog" json:"bag_ovog"`
	EronhiiUnelgeeid *int    `gorm:"column:eronhii_unelgeeid" json:"eronhii_unelgeeid"`
	ID               int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
}

func (s *SubYronhiiunelgeeBagEronhiiUnelgee76) TableName() string {
	return "sub_yronhiiunelgee_bag"
}

//  TableName sets the insert table name for this struct type
func (e *EronhiiUnelgee76) TableName() string {
	return "eronhii_unelgee"
}
