package models

import "time"

type Comment struct {
	Comments  string     `gorm:"column:comments" json:"comments"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	FormID    int        `gorm:"column:form_id" json:"form_id"`
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int        `gorm:"column:user_id" json:"user_id"`
}

func (s *Comment) TableName() string {
	return "sub_comment"
}
