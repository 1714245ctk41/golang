package main

import (
	"context"
	"crawl_data/conf"
	route "crawl_data/pkg/route"

	"os"

	"github.com/sirupsen/logrus"
	"gitlab.com/goxp/cloud0/logger"
)

const (
	APPNAME = "Adapter"
)

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	// Dev
	logger.DefaultLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		PadLevelText:     true,
		ForceQuote:       true,
		QuoteEmptyFields: true,
	})

	_ = os.Setenv("PORT", conf.LoadEnv().Port)
	_ = os.Setenv("DB_HOST", conf.LoadEnv().DBHost)
	_ = os.Setenv("DB_PORT", conf.LoadEnv().DBPort)
	_ = os.Setenv("DB_USER", conf.LoadEnv().DBUser)
	_ = os.Setenv("DB_PASS", conf.LoadEnv().DBPass)
	_ = os.Setenv("DB_NAME", conf.LoadEnv().DBName)
	_ = os.Setenv("ENABLE_DB", conf.LoadEnv().EnableDB)

	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
