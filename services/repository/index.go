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

func Save(model interface{}) interface{} {
	err := database.DB.Create(model).Error
	logError("Save", err)

	return err
}

func FindOne(model interface{}) interface{} {
	err := database.DB.Find(model).Error
	logError("FindOne", err)

	return err
}

func FindAll(model interface{}) interface{} {
	err := database.DB.Last(model).Error
	logError("FindAll", err)

	return err
}

func Update(id string, model interface{}) interface{} {
	err := database.DB.Update(id, model).Error
	logError("Update", err)

	return err
}

func Delete(id string, model interface{}) interface{} {
	err := database.DB.Delete(id, model).Error
	logError("Delete", err)

	return err
}
