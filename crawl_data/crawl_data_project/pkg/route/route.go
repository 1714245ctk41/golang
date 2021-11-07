package route

import (
	"crawl_data/pkg/handlers"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"gitlab.com/goxp/cloud0/service"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting *extraSetting
}

func NewService() *Service {
	// l := logger.WithField("Func", "New service")
	s := &Service{
		service.NewApp("Ecom Service", "v1.0"),
		&extraSetting{},
	}
	_ = env.Parse(s.setting)
	db := s.GetDB()

	if s.setting.DbDebugEnable {
		db = db.Debug()
	}
	s.Router.Use(gin.Logger())

	MongoRouter(s.Router, db)

	migrateHandler := handlers.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrateHandler.Migrate)

	return s
}
