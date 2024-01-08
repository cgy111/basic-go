package ioc

import (
	"basic-go/webook/config"
	"basic-go/webook/internal/repository/dao"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	fmt.Println(err)
	if err != nil {
		//只在初始化的时候panic
		//panic相当于整个goroutine结束
		//一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
