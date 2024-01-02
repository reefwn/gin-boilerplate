package main

import (
	"gin-boilerplate/config"
	"gin-boilerplate/helpers/logger"
	"gin-boilerplate/routers"
	"gin-boilerplate/services/database"
	"gin-boilerplate/services/migration"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Bangkok")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig error: %s", err)
	}
	dbDsn := config.DbConfiguration()

	if err := database.DbConnection(dbDsn); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	// TODO: separate migration
	migration.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
