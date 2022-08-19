package grid

import "github.com/lambda-platform/lambda/datagrid"

import "github.com/lambda-platform/lambda/models"

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type ViewHongiinMedee330 struct {
	Aimagname       *string    `gorm:"column:aimagname" json:"aimagname"`
	AyulDedDedTorol *string    `gorm:"column:ayul_ded_ded_torol" json:"ayul_ded_ded_torol"`
	AyulDedTorol    *string    `gorm:"column:ayul_ded_torol" json:"ayul_ded_torol"`
	AyulTorol       *string    `gorm:"column:ayul_torol" json:"ayul_torol"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`
	ID              *int       `gorm:"column:id" json:"id"`
	Sumname         *string    `gorm:"column:sumname" json:"sumname"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type HongiinMedeeMainTable330 struct {
	Aimagid           *int       `gorm:"column:aimagid" json:"aimagid"`
	AjilahHuch        *string    `gorm:"column:ajilah_huch" json:"ajilah_huch"`
	AvsanArgaHemjee   *string    `gorm:"column:avsan_arga_hemjee" json:"avsan_arga_hemjee"`
	AyulDedDedTorolID *int       `gorm:"column:ayul_ded_ded_torol_id" json:"ayul_ded_ded_torol_id"`
	AyulDedTorolID    *int       `gorm:"column:ayul_ded_torol_id" json:"ayul_ded_torol_id"`
	AyulTorolID       *int       `gorm:"column:ayul_torol_id" json:"ayul_torol_id"`
	Bairshil          *string    `gorm:"column:bairshil" json:"bairshil"`
	CreatedAt         *time.Time `gorm:"column:created_at" json:"created_at"`
	DOgnoo            *time.Time `gorm:"column:d_ognoo" json:"d_ognoo"`
	EOgnoo            *time.Time `gorm:"column:e_ognoo" json:"e_ognoo"`
	ID                int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MAvsanOgnooObeg   *time.Time `gorm:"column:m_avsan_ognoo_obeg" json:"m_avsan_ognoo_obeg"`
	MAvsanOgnooSalbar *time.Time `gorm:"column:m_avsan_ognoo_salbar" json:"m_avsan_ognoo_salbar"`
	Medeelel          *string    `gorm:"column:medeelel" json:"medeelel"`
	NiitHohirol       *string    `gorm:"column:niit_hohirol" json:"niit_hohirol"`
	Orgorog           *string    `gorm:"column:orgorog" json:"orgorog"`
	SumID             *int       `gorm:"column:sum_id" json:"sum_id"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Urtrag            *string    `gorm:"column:urtrag" json:"urtrag"`
	UserID            *int       `gorm:"column:user_id" json:"user_id"`
	Vaktsin           *string    `gorm:"column:vaktsin" json:"vaktsin"`
}

func (h *HongiinMedeeMainTable330) TableName() string {
	return "hongiin_medee"
}

//  TableName sets the insert table name for this struct type
func (v *ViewHongiinMedee330) TableName() string {
	return "view_hongiin_medee"
}

var ViewHongiinMedee330Datagrid datagrid.Datagrid = datagrid.Datagrid{
	Name:      "Хоногийн мэдээ",
	Identity:  "id",
	DataTable: "view_hongiin_medee",
	MainTable: "hongiin_medee",
	DataModel: new(ViewHongiinMedee330),
	Data:      new([]ViewHongiinMedee330),
	MainModel: new(HongiinMedeeMainTable330),
	Columns: []datagrid.Column{
		datagrid.Column{Model: "ayul_torol", Label: "Аюулын төрөл"},
		datagrid.Column{Model: "ayul_ded_torol", Label: "Аюулын дэд төрөл"},
		datagrid.Column{Model: "ayul_ded_ded_torol", Label: "Аюулын дэд дэд төрөл"},
		datagrid.Column{Model: "aimagname", Label: "Аймаг / Хот "},
		datagrid.Column{Model: "sumname", Label: "Сум / Дүүрэг"},
	},
	ColumnList: []string{"id", "ayul_torol", "ayul_ded_torol", "ayul_ded_ded_torol", "aimagname", "sumname"},
	Filters: map[string]string{
		"id": "",

		"ayul_torol_id": "",

		"ayul_ded_torol_id": "",

		"ayul_ded_ded_torol_id": "",

		"aimagid": "",

		"sum_id": "",

		"bairshil": "",

		"urtrag": "",

		"orgorog": "",

		"m_avsan_ognoo_obeg": "",

		"m_avsan_ognoo_salbar": "",

		"e_ognoo": "",

		"d_ognoo": "",

		"ajilah_huch": "",

		"vaktsin": "",

		"niit_hohirol": "",

		"avsan_arga_hemjee": "",

		"created_at": "",

		"updated_at": "",

		"user_id": "",

		"ayul_torol": "",

		"ayul_ded_torol": "",

		"ayul_ded_ded_torol": "",

		"aimagname": "",

		"sumname": "",

		"medeelel": "",
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
	FillVirtualColumns: fillVirtualColumnsViewHongiinMedee330,
}

func fillVirtualColumnsViewHongiinMedee330(rowsPre interface{}) interface{} {
	return rowsPre
}
