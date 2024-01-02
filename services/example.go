package services

import (
	"gin-boilerplate/models"
	"gin-boilerplate/services/repository"
)

func FindAllExample() []*models.Example {
	var example []*models.Example
	repository.FindAll(&example)

	return example
}
