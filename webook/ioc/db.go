package ioc

import (
	"basic-go/webook/internal/repository/dao"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//dsn := viper.GetString("db.mysql.dsn")
	//fmt.Println(dsn)

	type Config struct {
		DSN string `yaml:"dsn"`
	}
	//var cfg Config

	var cfg Config = Config{
		DSN: "root:root@tcp(47.123.5.217:13316)/mysql2",
	}

	err := viper.UnmarshalKey("db.mysql", &cfg)
	if err != nil {
		panic(err)
	}
	//db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	db, err := gorm.Open(mysql.Open(cfg.DSN))
	//fmt.Println(err)
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
