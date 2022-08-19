package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type News63 struct {
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	Delgerengui *string    `gorm:"column:delgerengui" json:"delgerengui"`
	Garchig     *string    `gorm:"column:garchig" json:"garchig"`
	ID          int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsTypeID  *int       `gorm:"column:news_type_id" json:"news_type_id"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UrsahEseh   *int       `gorm:"column:ursah_eseh" json:"ursah_eseh"`
	Zurag       *string    `gorm:"column:zurag" json:"zurag"`
}

type SubNewsSocialTypiesNews63 struct {
	ID           int  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NewsID       *int `gorm:"column:news_id" json:"news_id"`
	SocialTypeID *int `gorm:"column:social_type_id" json:"social_type_id"`
}

func (s *SubNewsSocialTypiesNews63) TableName() string {
	return "sub_news_social_typies"
}

//  TableName sets the insert table name for this struct type
func (n *News63) TableName() string {
	return "news"
}
