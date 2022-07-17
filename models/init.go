package models

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB = Init()
var RDB = InitRedis()

func Init()*gorm.DB{
	dsn := "root:zxc6545398@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

//init redis
func InitRedis()*redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password set
		DB: 0, // use default DB
	})
	return client
}