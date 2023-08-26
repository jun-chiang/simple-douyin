package config

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBase struct {
	Conn *gorm.DB
}

var (
	instance *DataBase
	doOnce   sync.Once
)

// 把数据库连接封装成单例模式
func GetInstance() *DataBase {
	doOnce.Do(func() {
		dsn := "douyin:123456789@tcp(192.168.209.159:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = &DataBase{Conn: conn}
	})
	return instance
}
