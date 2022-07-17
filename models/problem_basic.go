package models

import (
	"gorm.io/gorm"
	"time"
)

type ProblemBasic struct {
	ID 					int64 			`gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         time.Time     	 `json:"deleted_at"`
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`                  // 问题表的唯一标识
	ProblemCategory []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`      // 关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`                       // 文章标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`                           // 文章正文
	MaxRunTime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`                // 最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`                        // 最大运行内存
	TestCase         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"test_cases"` // 关联测试用例表
	PassNum           int64              `gorm:"column:pass_num;type:int(11);" json:"pass_num"`                      // 通过次数
	SubmitNum         int64              `gorm:"column:submit_num;type:int(11);" json:"submit_num"`                  // 提交次数
}

//type ProblemBasic struct {
//	ID int64 `gorm:"primarykey;" db:"id" json:"id" form:"id"`
//	Identity string `gorm:"column:identity" db:"identity" json:"identity" form:"identity"`
//	Title string `gorm:"column:title" db:"title" json:"title" form:"title"`  //  题目
//	Content string `gorm:"column:content" db:"content" json:"content" form:"content"`  //  正文描述
//	MaxRunTime int `gorm:"column:max_runtime" db:"max_runtime" json:"max_runtime" form:"max_runtime"`
//	MaxMem int `gorm:"column:max_mem" db:"max_mem" json:"max_mem" form:"max_mem"`
//	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
//	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
//	DeletedAt time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
//	//others
//	ProblemCategory []*ProblemCategory `gorm:"foreignkey:problem_id;reference:id" json:"problem_categories"`
//	TestCase []*TestCase `gorm:"foreignkey:problem_identity;reference:identity" json:"test_cases"`
//}

func (problem *ProblemBasic) TabName() string {
	return "problem_basic"
}

func GetDefaultProblemList( keyword string,categoryIdentity string) *gorm.DB {
	//keyword 模糊查询
	tx:=DB.Model(new(ProblemBasic)).Preload("ProblemCategory").Preload("ProblemCategory.CategoryBasic").Where("title like ? OR content like ?", "%"+keyword+"%","%"+keyword+"%")

	//分类信息
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category ON problem_basic.id = problem_category.problem_id").Where("problem_category.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	}

	return tx
}