package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewAanbTzErhOlgoh71 struct {
	BNer        *string `gorm:"column:b_ner" json:"b_ner"`
	DisplayName *string `gorm:"column:display_name" json:"display_name"`
	ID          *int    `gorm:"column:id" json:"id"`
	Login       *string `gorm:"column:login" json:"login"`
}

type UsersMainTable71 struct {
	AimagID        *int       `gorm:"column:aimag_id" json:"aimag_id"`
	AlbanTushaalID *int       `gorm:"column:alban_tushaal_id" json:"alban_tushaal_id"`
	AngiID         *int       `gorm:"column:angi_id" json:"angi_id"`
	Avatar         *string    `gorm:"column:avatar" json:"avatar"`
	BID            *int       `gorm:"column:b_id" json:"b_id"`
	BagID          *int       `gorm:"column:bag_id" json:"bag_id"`
	Bio            *string    `gorm:"column:bio" json:"bio"`
	Birthday       *DB.Date   `gorm:"column:birthday" json:"birthday"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	Email          string     `gorm:"column:email" json:"email"`
	FcmToken       *string    `gorm:"column:fcm_token" json:"fcm_token"`
	FirstName      *string    `gorm:"column:first_name" json:"first_name"`
	Gender         *string    `gorm:"column:gender" json:"gender"`
	ID             int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	LastName       *string    `gorm:"column:last_name" json:"last_name"`
	Login          string     `gorm:"column:login" json:"login"`
	Password       string     `gorm:"column:password" json:"password"`
	Phone          *string    `gorm:"column:phone" json:"phone"`
	RegisterNumber string     `gorm:"column:register_number" json:"register_number"`
	Role           *int       `gorm:"column:role" json:"role"`
	SalbariID      *int       `gorm:"column:salbari_id" json:"salbari_id"`
	Status         *string    `gorm:"column:status" json:"status"`
	SumID          *int       `gorm:"column:sum_id" json:"sum_id"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserTypeID     *int       `gorm:"column:user_type_id" json:"user_type_id"`
}

func (u *UsersMainTable71) TableName() string {
	return "users"
}

//  TableName sets the insert table name for this struct type
func (v *ViewAanbTzErhOlgoh71) TableName() string {
	return "view_aanb_tz_erh_olgoh"
}

var ViewAanbTzErhOlgoh71Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "ААНБ-д эрх олгох",
	Identity:  "id",
	DataTable: "view_aanb_tz_erh_olgoh",
	MainTable: "users",
	DataModel: new(ViewAanbTzErhOlgoh71),
	Data:      new([]ViewAanbTzErhOlgoh71),
	MainModel: new(UsersMainTable71),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "b_ner", Label: "Байгууллага"},
		datagrid.Column{Model: "display_name", Label: "Хандах эрх"},
		datagrid.Column{Model: "login", Label: "Нэвтрэх нэр"},
	},
	ColumnList: []string{"id", "b_ner", "display_name", "login"},
	Filters: map[string]string{
		"id": "",

		"b_ner": "",

		"display_name": "",

		"login": "",

		"role": "",

		"b_id": "",
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
	FillVirtualColumns: fillVirtualColumnsViewAanbTzErhOlgoh71,
}

func fillVirtualColumnsViewAanbTzErhOlgoh71(rowsPre interface{}) interface{} {
	return rowsPre
}
