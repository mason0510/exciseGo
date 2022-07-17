package models

import (
	"time"
)

type CategoryBasic struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	ParentId int64 `gorm:"column:parent_id" db:"parent_id" json:"parent_id" form:"parent_id"`  //  父级id
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (category *CategoryBasic ) TabName() string {
	return "category_basic"
}
