package models

import (
	"time"
)

type UserBasic struct {
	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"` //default main key
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	PassWord string `gorm:"column:password" db:"password" json:"password" form:"password"`
	Phone string `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`
	Mail      string         `gorm:"column:mail;type:varchar(100);" json:"mail"`
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
	Identity string `gorm:"column:identity" db:"identity" json:"identity" form:"identity"`
	PassNum int64 `gorm:"column:pass_num" db:"pass_num" json:"pass_num" form:"pass_num"`
	SubmitNum int64 `gorm:"column:submit_num" db:"submit_num" json:"submit_num" form:"submit_num"`
	IsAdmin int64 `gorm:"column:is_admin" db:"is_admin" json:"is_admin" form:"is_admin"`  //  0 不是 1 是
}

//type UserBasic struct {
//	gorm.Model
//	ID        uint           `gorm:"primarykey;" json:"id"`
//	Name      string         `gorm:"column:name;type:varchar(100);" json:"name"`        // 用户名
//	Password  string         `gorm:"column:password;type:varchar(32);" json:"password"` // 密码
//	Phone     string         `gorm:"column:phone;type:varchar(20);" json:"phone"`       // 手机号
//	Mail      string         `gorm:"column:mail;type:varchar(100);" json:"mail"`        // 邮箱
//	CreatedAt time.Time      `json:"created_at"`
//	UpdatedAt time.Time      `json:"updated_at"`
//	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
//	Identity  string         `gorm:"column:identity;type:varchar(36);" json:"identity"` // 用户的唯一标识
//	PassNum   int64          `gorm:"column:pass_num;type:int(11);" json:"pass_num"`     // 通过的次数
//	SubmitNum int64          `gorm:"column:submit_num;type:int(11);" json:"submit_num"` // 提交次数
//	IsAdmin   int            `gorm:"column:is_admin;type:tinyint(1);" json:"is_admin"`  // 是否是管理员【0-否，1-是】
//}



//type UserBasic struct {
//	gorm.Model
//	ID int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
//	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
//	Identity  string         `gorm:"column:identity;type:varchar(36);" json:"identity"` // 用户的唯一标识
//	Password string `gorm:"column:password" db:"password" json:"password" form:"password"`
//	Phone string `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`
//	Mail string `gorm:"column:mail" db:"mail" json:"mail" form:"mail"`
//	PassNum   int64          `gorm:"column:pass_num;type:int(11);" json:"pass_num"`     // 通过的次数
//	SubmitNum int64          `gorm:"column:submit_num;type:int(11);" json:"submit_num"` // 提交次数
//	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
//	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
//	DeletedAt time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
//	IsAdmin   int            `gorm:"column:is_admin;type:tinyint(1);" json:"is_admin"`  // 是否是管理员【0-否，1-是】
//}

func (user *UserBasic) TabName() string {
	return "user_basic"
}