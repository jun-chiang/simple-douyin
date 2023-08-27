package migrations

import (
	"github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/config"
	dbModels "github.com/jun-chiang/simple-douyin/userservice/infrastructure/persistence/models"
)

// 数据迁移操作入口
func Migration() {
	migrateUser()
}

func migrateUser() {
	db := config.GetInstance().Conn
	db.AutoMigrate(&dbModels.User{})

}
