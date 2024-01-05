package repository

import (
	"gin-boilerplate/helpers/logger"
	"gin-boilerplate/services/database"
)

func logError(fnName string, err error) {
	if err != nil {
		logger.Errorf("[Repository]: %s, %v", fnName, err)
	}
}

func Save(model interface{}) error {
	err := database.DB.Create(model).Error
	logError("Save", err)

	return err
}

func Find(model interface{}) error {
	err := database.DB.Find(model).Error
	logError("Find", err)

	return err
}

func Update(model interface{}, update interface{}) error {
	err := database.DB.Model(model).Updates(update).Error
	logError("Update", err)

	return err
}

func Delete(model interface{}) error {
	err := database.DB.Delete(model).Error
	logError("Delete", err)

	return err
}
