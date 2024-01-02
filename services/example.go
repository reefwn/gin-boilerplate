package services

import (
	"gin-boilerplate/models"
	"gin-boilerplate/services/repository"

	"github.com/google/uuid"
)

func FindAllExample() []*models.Example {
	var example []*models.Example
	repository.FindAll(&example)

	return example
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
