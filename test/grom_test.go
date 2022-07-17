package test

import (
	"exciseGo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGromtest(t *testing.T) {
	dsn := "root:zxc6545398@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	data:=make([]*models.ProblemBasic,0)
	err = db.Find(&data).Error
	if err != nil {
		t.Error(err)
	}
	for _,v:=range data{
		t.Log(v)
	}

}