package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewEronhiiUnelgee79 struct {
	ATushaal             *string    `gorm:"column:a_tushaal" json:"a_tushaal"`
	Aimagname            *string    `gorm:"column:aimagname" json:"aimagname"`
	Bagname              *string    `gorm:"column:bagname" json:"bagname"`
	Baiguullaga          *string    `gorm:"column:baiguullaga" json:"baiguullaga"`
	ConfirmStatus        *string    `gorm:"column:confirm_status" json:"confirm_status"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	EvalutionSector      *string    `gorm:"column:evalution_sector" json:"evalution_sector"`
	FirstName            *string    `gorm:"column:first_name" json:"first_name"`
	GeneralEvalutionType *string    `gorm:"column:general_evalution_type" json:"general_evalution_type"`
	ID                   *int       `gorm:"column:id" json:"id"`
	Indicator            *string    `gorm:"column:indicator" json:"indicator"`
	IndicatorType        *string    `gorm:"column:indicator_type" json:"indicator_type"`
	LastName             *string    `gorm:"column:last_name" json:"last_name"`
	Ner                  *string    `gorm:"column:ner" json:"ner"`
	Ognoo                *DB.Date   `gorm:"column:ognoo" json:"ognoo"`
	Owog                 *string    `gorm:"column:owog" json:"owog"`
	SubIndicator         *string    `gorm:"column:sub_indicator" json:"sub_indicator"`
	Sumname              *string    `gorm:"column:sumname" json:"sumname"`
	UHiijDuussanOgnoo    *DB.Date   `gorm:"column:u_hiij_duussan_ognoo" json:"u_hiij_duussan_ognoo"`
	UHiijEhelsenOgnoo    *DB.Date   `gorm:"column:u_hiij_ehelsen_ognoo" json:"u_hiij_ehelsen_ognoo"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UrDun                *string    `gorm:"column:ur_dun" json:"ur_dun"`
	ZeregleliinHemjee    *int       `gorm:"column:zeregleliin_hemjee" json:"zeregleliin_hemjee"`
	Zurag                *string    `gorm:"column:zurag" json:"zurag"`
}

type EronhiiUnelgeeMainTable79 struct {
	ATushaal               *string    `gorm:"column:a_tushaal" json:"a_tushaal"`
	AimagID                *int       `gorm:"column:aimag_id" json:"aimag_id"`
	BagID                  *int       `gorm:"column:bag_id" json:"bag_id"`
	Baiguullaga            *string    `gorm:"column:baiguullaga" json:"baiguullaga"`
	Bairshil               *string    `gorm:"column:bairshil" json:"bairshil"`
	ConfirmStatusID        *int       `gorm:"column:confirm_status_id" json:"confirm_status_id"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at"`
	DangerTypeID           *int       `gorm:"column:danger_type_id" json:"danger_type_id"`
	DeletedAt              *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
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
	ZeregleliinHemjee      *int       `gorm:"column:zeregleliin_hemjee" json:"zeregleliin_hemjee"`
	Zurag                  *string    `gorm:"column:zurag" json:"zurag"`
}

func (e *EronhiiUnelgeeMainTable79) TableName() string {
	return "eronhii_unelgee"
}

//  TableName sets the insert table name for this struct type
func (v *ViewEronhiiUnelgee79) TableName() string {
	return "view_eronhii_unelgee"
}

var ViewEronhiiUnelgee79Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Ерөнхий үнэлгээний тайлан",
	Identity:  "id",
	DataTable: "view_eronhii_unelgee",
	MainTable: "eronhii_unelgee",
	DataModel: new(ViewEronhiiUnelgee79),
	Data:      new([]ViewEronhiiUnelgee79),
	MainModel: new(EronhiiUnelgeeMainTable79),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "last_name", Label: "Овог"},
		datagrid.Column{Model: "first_name", Label: "Нэр"},
		datagrid.Column{Model: "ognoo", Label: "Огноо"},
		datagrid.Column{Model: "zeregleliin_hemjee", Label: "Зэрэглэлийн хэмжээ"},
		datagrid.Column{Model: "owog", Label: "Овог"},
		datagrid.Column{Model: "ner", Label: "Нэр"},
		datagrid.Column{Model: "baiguullaga", Label: "Байгууллага"},
		datagrid.Column{Model: "a_tushaal", Label: "Албан тушаал"},
		datagrid.Column{Model: "ur_dun", Label: "Үр дүн"},
		datagrid.Column{Model: "zurag", Label: "Зураг"},
		datagrid.Column{Model: "u_hiij_ehelsen_ognoo", Label: "Үнэлгээ хийж эхэлсэн огноо "},
		datagrid.Column{Model: "u_hiij_duussan_ognoo", Label: "Үнэлгээ хийж дууссан огноо "},
		datagrid.Column{Model: "general_evalution_type", Label: "Ерөнхий үнэлгээний төрөл "},
		datagrid.Column{Model: "evalution_sector", Label: "Үнэлгээ хийлгэсэн салбар "},
	},
	ColumnList: []string{"id", "last_name", "first_name", "ognoo", "zeregleliin_hemjee", "owog", "ner", "baiguullaga", "a_tushaal", "ur_dun", "zurag", "u_hiij_ehelsen_ognoo", "u_hiij_duussan_ognoo", "general_evalution_type", "evalution_sector", "aimagname", "sumname", "bagname", "indicator", "sub_indicator", "indicator_type", "confirm_status"},
	Filters: map[string]string{
		"id": "",

		"last_name": "",

		"first_name": "",

		"user_id": "",

		"general_evalution_type_id": "",

		"evalution_sector_id": "",

		"aimag_id": "",

		"sum_id": "",

		"bag_id": "",

		"bairshil": "",

		"danger_type_id": "",

		"ognoo": "",

		"indicator_id": "",

		"sub_indicator_id": "",

		"zeregleliin_hemjee": "",

		"owog": "",

		"ner": "",

		"baiguullaga": "",

		"a_tushaal": "",

		"ur_dun": "",

		"zurag": "",

		"u_hiij_ehelsen_ognoo": "",

		"u_hiij_duussan_ognoo": "",

		"confirm_status_id": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"general_evalution_type": "",

		"evalution_sector": "",

		"aimagname": "",

		"sumname": "",

		"bagname": "",

		"danger_type": "",

		"indicator": "",

		"sub_indicator": "",

		"indicator_type": "",

		"confirm_status": "",
	},
	Relations:   []models.GridRelation{},
	Condition:   "",
	Aggergation: "",
	BeforeFetch: nil,

	AfterFetch: nil,

	BeforeDelete: nil,

	AfterDelete: nil,

	BeforePrint:        nil,
	TriggerNameSpace:   "",
	FillVirtualColumns: fillVirtualColumnsViewEronhiiUnelgee79,
}

func fillVirtualColumnsViewEronhiiUnelgee79(rowsPre interface{}) interface{} {
	return rowsPre
}
