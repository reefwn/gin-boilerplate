package migration

import (
	"gin-boilerplate/models"
	"gin-boilerplate/services/database"
)

func Migrate() {
	// TODO: separate migration each models
	var migrationModels = []interface{}{&models.Example{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
