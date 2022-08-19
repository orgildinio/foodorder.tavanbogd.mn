package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewUsers281 struct {
	Aimagname      *string    `gorm:"column:aimagname" json:"aimagname"`
	AlbanTushaal   *string    `gorm:"column:alban_tushaal" json:"alban_tushaal"`
	Angi           *string    `gorm:"column:angi" json:"angi"`
	Bagname        *string    `gorm:"column:bagname" json:"bagname"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	FirstName      *string    `gorm:"column:first_name" json:"first_name"`
	ID             *int       `gorm:"column:id" json:"id"`
	LastName       *string    `gorm:"column:last_name" json:"last_name"`
	Login          *string    `gorm:"column:login" json:"login"`
	RegisterNumber *string    `gorm:"column:register_number" json:"register_number"`
	Salbar         *string    `gorm:"column:salbar" json:"salbar"`
	Sumname        *string    `gorm:"column:sumname" json:"sumname"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserType       *string    `gorm:"column:user_type" json:"user_type"`
}

type UsersMainTable281 struct {
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

func (u *UsersMainTable281) TableName() string {
	return "users"
}

//  TableName sets the insert table name for this struct type
func (v *ViewUsers281) TableName() string {
	return "view_users"
}

var ViewUsers281Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Хэрэглэгч",
	Identity:  "id",
	DataTable: "view_users",
	MainTable: "users",
	DataModel: new(ViewUsers281),
	Data:      new([]ViewUsers281),
	MainModel: new(UsersMainTable281),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "login", Label: "Нэвтрэх нэр"},
		datagrid.Column{Model: "register_number", Label: "Регистрийн дугаар"},
		datagrid.Column{Model: "last_name", Label: "Овог"},
		datagrid.Column{Model: "first_name", Label: "Нэр"},
		datagrid.Column{Model: "user_type", Label: "Хэрэглэгчийн төрөл"},
		datagrid.Column{Model: "aimagname", Label: "Аймаг/Нийслэл"},
		datagrid.Column{Model: "sumname", Label: " Сум / Дүүрэг"},
		datagrid.Column{Model: "bagname", Label: " Баг / Хороо"},
		datagrid.Column{Model: "angi", Label: "Анги"},
		datagrid.Column{Model: "salbar", Label: "Салбар"},
		datagrid.Column{Model: "alban_tushaal", Label: "Албан тушаал"},
	},
	ColumnList: []string{"id", "login", "register_number", "last_name", "first_name", "user_type", "aimagname", "sumname", "bagname", "angi", "salbar", "alban_tushaal"},
	Filters: map[string]string{
		"id": "",

		"created_at": "",

		"updated_at": "",

		"deleted_at": "",

		"status": "",

		"role": "",

		"login": "",

		"email": "",

		"register_number": "",

		"avatar": "",

		"bio": "",

		"last_name": "",

		"first_name": "",

		"birthday": "",

		"phone": "",

		"gender": "",

		"password": "",

		"fcm_token": "",

		"b_id": "",

		"user_type_id": "",

		"alban_tushaal_id": "",

		"aimag_id": "",

		"sum_id": "",

		"bag_id": "",

		"angi_id": "",

		"salbari_id": "",

		"user_type": "",

		"aimagname": "",

		"sumname": "",

		"bagname": "",

		"angi": "",

		"salbar": "",

		"alban_tushaal": "",
	},
	Relations:   []models.GridRelation{},
	Condition:   "id >= 2",
	Aggergation: "",
	BeforeFetch: nil,

	AfterFetch: nil,

	BeforeDelete: nil,

	AfterDelete: nil,

	BeforePrint:        nil,
	TriggerNameSpace:   "",
	FillVirtualColumns: fillVirtualColumnsViewUsers281,
}

func fillVirtualColumnsViewUsers281(rowsPre interface{}) interface{} {
	return rowsPre
}
