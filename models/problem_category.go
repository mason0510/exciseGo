package models

type ProblemCategory struct {
	ProblemId int64 `gorm:"column:problem_id" db:"problem_id" json:"problem_id" form:"problem_id"`
	CategoryId int64 `gorm:"column:category_id" db:"category_id" json:"category_id" form:"category_id"`
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id;" json:"category_basic"` // 关联分类的基础信息表

}

func (category *ProblemCategory ) TabName() string {
	return "problem_category"
}