package models

import (
	"gorm.io/gorm"
	"time"
)
//go language via high case and lowercase to control the obtain of table name and variable name
type SubmitBasic struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	Identity string `gorm:"column:identity" db:"identity" json:"identity" form:"identity"`
	ProblemIdentity string `gorm:"column:problem_identity" db:"problem_identity" json:"problem_identity" form:"problem_identity"`
	UserIdentity string `gorm:"column:user_identity" db:"user_identity" json:"user_identity" form:"user_identity"`
	Path string `gorm:"column:path" db:"path" json:"path" form:"path"`
	Status int64 `gorm:"column:status" db:"status" json:"status" form:"status"` //0 waiting 1 right 2 error 3 timeout 4 compile error 5 runtime error
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
	ProblemBasic *ProblemBasic `gorm:"foreignkey:identity;reference:problem_identity"` //connnect problem table
	UserBasic *UserBasic `gorm:"foreignkey:identity;reference:user_identity"` //connect user table
}
func (submit *SubmitBasic) TabName() string {
	return "submit_basic"
}

func GetDefaultSubmitList(problemIdentity string,userIdentity string,status int) *gorm.DB {
	//keyword  choose what to ignore
	tx:=DB.Model(new(SubmitBasic)).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return db.Omit("content")
	}).Preload("UserBasic")

	//题目信息
	if problemIdentity != "" {
		tx.Where("problem_identity = ?", problemIdentity)
	}

	//用户信息
	if userIdentity != "" {
		tx.Where("user_identity = ?", userIdentity)
	}

	if status != 0 {
		tx.Where("status = ?", status)
	}

	return tx
}


