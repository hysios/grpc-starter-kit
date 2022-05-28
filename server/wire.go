//go:build wireinject
// +build wireinject

package server

import (
	// "github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/hysios/grpc-starter-kit/service"
	"github.com/hysios/mx/platform/config"
	"github.com/hysios/mx/platform/logger"
	"github.com/hysios/mx/platform/model"
	"github.com/hysios/mx/platform/policy"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func SetupServer(addr string) (*grpc.Server, error) {
	wire.Build(config.ProviderDatabase, config.LoadDefault, logger.ProviderLog, policy.EnforecerProvider, service.BuildGPCServer)
	return nil, nil
}

func SetupDB() {
	config.SetupDB(func(db *gorm.DB, logger *logger.Logger) error {
		log := logger.Sugar()
		log.Infof("auto migrates database")
		// 迁移模型
		if err := db.AutoMigrate(
		//TODO: add your models here
		// &model.User{},
		// &model.Certificate{},
		); err != nil {
			log.Errorf("failed to migrate: %v", err)
			return err
		}

		// 把用户 User 表的 ID 字段设置为自增起始为 100000
		model.AutoIncrementStart(db, &model.User{}, "id", 100000)

		// 创建管理员
		_, err := model.CreateAdmin(db, "admin", "admin")
		return err
	})
}
