package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type DundHugtsaaStrategy54 struct {
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at"`
	ID             int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MongolStrategy *string    `gorm:"column:mongol_strategy" json:"mongol_strategy"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SubDundHugtsaaStrategyDundHugtsaaStrategy54 struct {
	DundHugtsaaStrategyID *int    `gorm:"column:dund_hugtsaa_strategy_id" json:"dund_hugtsaa_strategy_id"`
	Heregjilt             *int    `gorm:"column:heregjilt" json:"heregjilt"`
	ID                    int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	TerguulehEn           *string `gorm:"column:terguuleh_en" json:"terguuleh_en"`
	TerguulehMn           *string `gorm:"column:terguuleh_mn" json:"terguuleh_mn"`
}

func (s *SubDundHugtsaaStrategyDundHugtsaaStrategy54) TableName() string {
	return "sub_dund_hugtsaa_strategy"
}

//  TableName sets the insert table name for this struct type
func (d *DundHugtsaaStrategy54) TableName() string {
	return "dund_hugtsaa_strategy"
}
