package services

import (
	"gin-boilerplate/models"
	"gin-boilerplate/services/repository"

	"github.com/google/uuid"
)

func FindAllExample() []*models.Example {
	var example []*models.Example
	repository.Find(&example)

	return example
}

func FindOneExample(id string) *models.Example {
	model := models.Example{
		Id: uuid.MustParse(id),
	}
	repository.Find(&model)

	return &model
}

func CreateExample(body *models.Example) error {
	return repository.Save(&body)
}

func UpdateExample(body *models.Example) error {
	model := models.Example{
		Id: body.Id,
	}
	update := map[string]interface{}{
		"data": body.Data,
	}

	return repository.Update(&model, &update)
}

func DeleteExample(id string) error {
	model := models.Example{
		Id: uuid.MustParse(id),
	}

	return repository.Delete(&model)
}
