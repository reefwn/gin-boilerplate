package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnection(dbDsn string) error {
	var db = DB

	logMode := viper.GetBool("DB_LOG_MODE")

	logLevel := logger.Silent
	if logMode {
		logLevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(dbDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Fatalf("[DbConnection]: gorm.Open, %v", err)
		return err
	}

	DB = db
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
